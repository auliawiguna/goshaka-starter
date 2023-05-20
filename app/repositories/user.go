package repositories_v1

import (
	"fmt"
	"goshaka/app/models"
	"goshaka/app/models/scopes"
	"goshaka/app/structs"
	"goshaka/database"
	"goshaka/helpers"
	"strconv"
	"time"

	"github.com/goccy/go-json"

	appConfig "goshaka/configs"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
)

// Show user by given ID
//
//	param id string
//	return models.User
func UserShow(id string) models.User {
	db := database.DB
	var user models.User

	db.Preload("RoleUser.Role").Find(&user, "id = ?", id)

	return user
}

// Handle user login using google one tap
func LoginUsingGooleOneTap(c *fiber.Ctx) (models.User, string, error) {
	var user models.User
	var oneTapStructure structs.GoogleOneTap

	body := c.Body()

	err := json.Unmarshal(body, &oneTapStructure)

	if err != nil {
		return user, "", fmt.Errorf("payload error")
	}
	idToken := helpers.SanitiseText(oneTapStructure.IdToken)

	tokenInfo, err := helpers.VerifyIdToken(c.Context(), idToken)

	if err != nil {
		return user, "", fmt.Errorf("%s", err.Error())
	}

	var email string = tokenInfo.Email
	db := database.DB
	db.Preload("RoleUser.Role").Find(&user, "email = ?", email)

	if user.ID == 0 {
		return user, "", fmt.Errorf("credential cannot be found")
	}

	// Remove all reset password token
	removeResetPasswordTokens(fmt.Sprint(user.ID))

	return GenerateJwt(&user)
}

// Handle user login
//
//	receiver c *fiber.Ctx
//	return models.User, string, error
func Login(c *fiber.Ctx) (models.User, string, error) {

	var user models.User
	var loginStructure structs.Login

	body := c.Body()

	err := json.Unmarshal(body, &loginStructure)

	if err != nil {
		return user, "", fmt.Errorf("payload error")
	}
	email := loginStructure.Email
	password := loginStructure.Password

	db := database.DB
	db.Preload("RoleUser.Role").Find(&user, "email = ?", email)

	errHash := helpers.CompareHash(user.Password, password)

	if !errHash {
		return user, "", fmt.Errorf("credential cannot be found")
	}

	// Remove all reset password token
	go func() {
		removeResetPasswordTokens(fmt.Sprint(user.ID))
	}()

	return GenerateJwt(&user)
}

// To handle request reset password, back end will generate token, save plain token to database, and then send it to user using goroutine
//
//	param	c *fiber.Ctx
//	return	string, error
func RequestResetPassword(c *fiber.Ctx) (string, error) {

	var user models.User
	var isExists bool
	var requestResetPasswordStructure structs.RequestResetPassword

	body := c.Body()

	err := json.Unmarshal(body, &requestResetPasswordStructure)

	if err != nil {
		return "", fmt.Errorf("payload error")
	}
	email := helpers.SanitiseText(requestResetPasswordStructure.Email)

	db := database.DB

	// Check the existence first
	user, isExists = FindByEmail(email)
	if isExists {
		token := helpers.RandomNumber(6)
		hashedToken, err := helpers.EncryptText(token)
		if err != nil {
			return "failed", err
		}

		emailData := struct {
			FirstName   string
			Token       string
			HashedToken string
			FrontendUrl string
			AppUrl      string
		}{
			FirstName:   user.FirstName,
			Token:       string(token),
			HashedToken: hashedToken,
			FrontendUrl: appConfig.GetEnv("FRONTEND_URL"),
			AppUrl:      appConfig.GetEnv("APP_URL"),
		}

		db.Create(&models.UserToken{
			UserId:    user.ID,
			Type:      "reset_password",
			Token:     token,
			ExpiredAt: time.Now().Add(time.Minute * 3), // 3minutes
		})

		errS := helpers.SendMail(email, "Request Reset Password", "request_reset_password", emailData)
		if errS != nil {
			fmt.Println("error send email request_reset_password")
		}
	}

	return "success", nil
}

// To handle reset password,
// user will sends encrypted token, password, and password confirmation,
// once password has been updated, system will remove any reset password
// tokens related to current user
//
//	param	c *fiber.Ctx
//	return	string, error
func ResetPassword(c *fiber.Ctx) (string, error) {

	var user models.User

	var resetPasswordStructure structs.ResetPassword

	body := c.Body()

	err := json.Unmarshal(body, &resetPasswordStructure)

	if err != nil {
		return "", fmt.Errorf("payload error")
	}
	token, err := helpers.DecryptText(helpers.SanitiseText(resetPasswordStructure.Token))
	if err != nil {
		return "", fmt.Errorf("invalid request token")
	}
	password := helpers.SanitiseText(resetPasswordStructure.Password)

	db := database.DB

	var existingToken models.UserToken
	checkToken := db.Where("token = ?", token).Where("type = ?", "reset_password").Where("expired_at > ?", time.Now()).Find(&existingToken)
	if checkToken.RowsAffected == 0 {
		return "failed", fmt.Errorf("token not found")
	}

	// Check the existence first
	user = FindById(existingToken.UserId)
	if user.ID != 0 {
		// Update user
		db.Model(&user).Where("id = ?", user.ID).UpdateColumns(&models.User{
			Password: password,
		})
		// Remove all reset password token
		removeResetPasswordTokens(fmt.Sprint(user.ID))
	}

	return "success", nil
}

// To validate user's registration
// user will sends token, password, and email,
// once password has been updated, system will remove any registration
// tokens related to current user
//
//	param	c *fiber.Ctx
//	return	string, error
func ValidateRegistration(c *fiber.Ctx) (models.User, string, error) {

	var user models.User
	var loginStructure structs.RegistrationToken

	body := c.Body()

	err := json.Unmarshal(body, &loginStructure)

	if err != nil {
		return user, "", fmt.Errorf("payload error")
	}
	email := helpers.SanitiseText(loginStructure.Email)
	password := helpers.SanitiseText(loginStructure.Password)
	tokenPayload := helpers.SanitiseText(loginStructure.Token)

	db := database.DB
	db.Preload("RoleUser.Role").Find(&user, "email = ?", email)

	errHash := helpers.CompareHash(user.Password, password)

	if !errHash {
		return user, "", fmt.Errorf("credential cannot be found")
	}

	// Find the token
	var token models.UserToken
	checkToken := db.Where("user_id = ?", user.ID).Where("type = ?", "registration").Where("expired_at > ?", time.Now()).First(&token)
	if checkToken.RowsAffected == 0 {
		return user, "", fmt.Errorf("token is not found")
	}

	// Verify token
	errToken := helpers.CompareHash(token.Token, tokenPayload)
	if !errToken {
		return user, "", fmt.Errorf("credential cannot be found")
	}

	// Delete token
	db.Unscoped().Delete(&token)

	// Validate user
	db.Model(&user).Update("validated_at", time.Now())

	// Generate token
	return GenerateJwt(&user)
}

// To resend a new user's registration token
// user will sends email,
//
//	param	c *fiber.Ctx
//	return	string, error
func ResendNewRegistrationToken(c *fiber.Ctx) error {

	var user models.User
	var loginStructure structs.RegistrationToken

	body := c.Body()

	err := json.Unmarshal(body, &loginStructure)

	if err != nil {
		return fmt.Errorf("payload error")
	}
	email := helpers.SanitiseText(loginStructure.Email)

	db := database.DB
	// find user
	db.Where("email = ?", email).Where("validated_at is null").Find(&user)

	if user.ID == 0 {
		return fmt.Errorf("user not found")
	}

	// Delete the token first
	db.Unscoped().Where("user_id = ?", user.ID).Where("type = ?", "registration").Delete(&models.UserToken{})

	token := helpers.RandomNumber(6)
	hashedToken, _ := helpers.CreateHash(token)

	emailData := struct {
		FirstName string
		Token     string
		AppUrl    string
	}{
		FirstName: user.FirstName,
		Token:     string(token),
		AppUrl:    appConfig.GetEnv("APP_URL"),
	}

	db.Create(&models.UserToken{
		UserId:    user.ID,
		Type:      "registration",
		Token:     hashedToken,
		ExpiredAt: time.Now().Add(time.Hour * 72), // 3days
	})

	errS := helpers.SendMail(email, "Complete your registration", "registration", emailData)
	if errS != nil {
		fmt.Println("error send email registration")
	}

	return nil
}

// To send an OTP to a user
//
//	param	c *fiber.Ctx
//	return	string, error
func SendOtpToken(c *fiber.Ctx) error {

	var user models.User
	var emailStruct structs.EmailOnly

	body := c.Body()

	err := json.Unmarshal(body, &emailStruct)

	if err != nil {
		return fmt.Errorf("payload error")
	}
	email := helpers.SanitiseText(emailStruct.Email)

	db := database.DB
	// find user
	db.Where("email = ?", email).Find(&user)

	if user.ID == 0 {
		return fmt.Errorf("user not found")
	}

	// Delete the token first
	db.Unscoped().Where("user_id = ?", user.ID).Where("type = ?", "otp").Delete(&models.UserToken{})

	token := helpers.RandomNumber(6)
	hashedToken, _ := helpers.CreateHash(token)

	emailData := struct {
		FirstName string
		Token     string
		AppUrl    string
	}{
		FirstName: user.FirstName,
		Token:     string(token),
		AppUrl:    appConfig.GetEnv("APP_URL"),
	}

	db.Create(&models.UserToken{
		UserId:    user.ID,
		Type:      "otp",
		Token:     hashedToken,
		ExpiredAt: time.Now().Add(time.Minute * 3), // 3 minutes
	})

	errS := helpers.SendMail(email, "OTP", "otp", emailData)
	if errS != nil {
		fmt.Println("error send email otp")
	}

	return nil
}

// To validate OTP request
// user will sends token, password, and email,
// once password has been updated, system will remove any registration
// tokens related to current user
//
//	param	c *fiber.Ctx
//	return	string, error
func ValidateOtpRequest(c *fiber.Ctx) (models.User, string, error) {

	var user models.User
	var loginStructure structs.EmailAndToken

	body := c.Body()

	err := json.Unmarshal(body, &loginStructure)

	if err != nil {
		return user, "", fmt.Errorf("payload error")
	}
	email := helpers.SanitiseText(loginStructure.Email)
	tokenPayload := helpers.SanitiseText(loginStructure.Token)

	db := database.DB
	db.Preload("RoleUser.Role").Find(&user, "email = ?", email)

	// Find the token
	var token models.UserToken
	checkToken := db.Where("user_id = ?", user.ID).Where("type = ?", "otp").Where("expired_at > ?", time.Now()).First(&token)
	if checkToken.RowsAffected == 0 {
		return user, "", fmt.Errorf("token is not found")
	}

	// Verify token
	errToken := helpers.CompareHash(token.Token, tokenPayload)
	if !errToken {
		return user, "", fmt.Errorf("credential cannot be found")
	}

	// Delete token
	db.Unscoped().Delete(&token)

	// Validate user
	db.Model(&user).Update("validated_at", time.Now())

	// Generate token
	return GenerateJwt(&user)
}

// Generate stateless JWT auth
//
//	param user models.User
//	return models.User, string, error
func GenerateJwt(user *models.User) (models.User, string, error) {
	tokenData := jwt.New(jwt.SigningMethodHS256)

	claims := tokenData.Claims.(jwt.MapClaims)
	claims["username"] = user.Username
	claims["id"] = user.ID
	claims["email"] = user.Email
	claims["exp"] = time.Now().Add(time.Hour * 24 * 30).Unix()

	secret := []byte(appConfig.GetEnv("JWT_KEY"))

	tokenString, err := tokenData.SignedString(secret)

	if err != nil {
		return *user, "", fmt.Errorf("error when generate JWT")
	}

	return *user, tokenString, nil
}

// Find a user by email address or username
//
//	param	email string
//	param	username string
//	return	models.User, bool
func FindByEmailOrUsername(email, username string) (models.User, bool) {
	db := database.DB

	var existingUser models.User
	checkUser := db.Where("email = ?", email).Or("username = ?", username).Find(&existingUser)
	if checkUser.RowsAffected > 0 {
		return existingUser, true
	}
	return existingUser, false
}

// Find a user by email address
//
//	param	email string
//	param	username string
//	return	models.User, bool
func FindByEmail(email string) (models.User, bool) {
	db := database.DB

	var existingUser models.User
	checkUser := db.Where("email = ?", email).Find(&existingUser)
	if checkUser.RowsAffected > 0 {
		return existingUser, true
	}
	return existingUser, false
}

// Find a user by ID
//
//	param	email string
//	param	username string
//	return	models.User, bool
func FindById(id uint) models.User {
	db := database.DB
	var user models.User

	db.Find(&user, "id = ?", id)

	return user

}

// To delete roles by user ID
//
//	params	uid	uint
//
// return void
func _DeleteRolesByUser(uid uint) {
	db := database.DB
	db.Unscoped().Delete(&models.RoleUser{}, "user_id = ?", uid)
}

// To delete all roles by user ID and repopulate a role by roleId
//
//	params	uid	uint
//	params	roleId	uint
//
// return void
func _ResetRole(uid, roleId uint) {
	db := database.DB
	_DeleteRolesByUser(uid)
	db.Create(&models.RoleUser{
		UserId: uid,
		RoleId: roleId,
	})
}

// To set a role by user ID
//
//	params	uid	uint
//	params	roleId	uint
//
// return void
func _SetRole(uid, roleId uint) {
	db := database.DB
	db.Create(&models.RoleUser{
		UserId: uid,
		RoleId: roleId,
	})
}

// Handle user's registration, user will sends email, username, password, first name, and last name
// System will sends verification token for users to verify their account
// Email will be handled by goroutine
//
//	params c *fiber.Ctx
//	return models.User, error
func Register(c *fiber.Ctx) (models.User, error) {

	var user models.User
	var userCreateStructure structs.UserCreate

	body := c.Body()

	err := json.Unmarshal(body, &userCreateStructure)

	if err != nil {
		return user, fmt.Errorf("payload error")
	}
	email := helpers.SanitiseText(userCreateStructure.Email)
	username := helpers.SanitiseText(userCreateStructure.Username)
	password := helpers.SanitiseText(userCreateStructure.Password)
	first_name := helpers.SanitiseText(userCreateStructure.FirstName)
	last_name := helpers.SanitiseText(userCreateStructure.LastName)

	db := database.DB

	// Check the existence first
	_, isExists := FindByEmailOrUsername(email, username)
	if isExists {
		return user, fmt.Errorf("registration is failed")
	}

	// Set User
	newUser := models.User{
		FirstName: first_name,
		LastName:  last_name,
		Password:  password,
		Email:     email,
		Username:  username,
	}
	db.Create(&newUser)

	token := helpers.RandomNumber(6)
	hashedToken, _ := helpers.CreateHash(token)

	emailData := struct {
		FirstName string
		Token     string
		AppUrl    string
	}{
		FirstName: first_name,
		Token:     string(token),
		AppUrl:    appConfig.GetEnv("APP_URL"),
	}

	db.Create(&models.UserToken{
		UserId:    newUser.ID,
		Type:      "registration",
		Token:     hashedToken,
		ExpiredAt: time.Now().Add(time.Hour * 72), // 3days
	})

	errS := helpers.SendMail(email, "Complete your registration", "registration", emailData)
	if errS != nil {
		fmt.Println("error send email registration")
	}

	// Set Role
	_SetRole(newUser.ID, 1)

	db.Preload("RoleUser.Role").Find(&user, "id = ?", newUser.ID)

	return user, nil
}

// Show all users
//
//	receiver pagination helpers.Pagination
//	return *helpers.Pagination, bool
func UserShowAll(pagination helpers.Pagination) *helpers.Pagination {
	db := database.DB
	var users []*models.User

	db.Scopes(scopes.Paginate(users, &pagination, db)).Preload("RoleUser.Role").Find(&users)
	pagination.Rows = users

	return &pagination
}

// Create a users
//
//	param c *fiber.Ctx
//	return models.User, error
func UserCreate(c *fiber.Ctx) (models.User, error) {
	db := database.DB

	var user models.User
	var isExists bool
	var userCreateStructure structs.UserCreate

	body := c.Body()

	err := json.Unmarshal(body, &userCreateStructure)

	if err != nil {
		return user, fmt.Errorf("payload error")
	}
	email := helpers.SanitiseText(userCreateStructure.Email)
	username := helpers.SanitiseText(userCreateStructure.Username)
	password := helpers.SanitiseText(userCreateStructure.Password)
	first_name := helpers.SanitiseText(userCreateStructure.FirstName)
	last_name := helpers.SanitiseText(userCreateStructure.LastName)
	role_id, _ := strconv.Atoi(helpers.SanitiseText(fmt.Sprint(userCreateStructure.RoleId)))

	user, isExists = FindByEmailOrUsername(email, username)
	if isExists {
		return user, fmt.Errorf("user already exists")
	}

	newUser := models.User{
		FirstName: first_name,
		LastName:  last_name,
		Password:  password,
		Email:     email,
		Username:  username,
	}
	db.Create(&newUser)

	// Set Role
	_SetRole(newUser.ID, uint(role_id))

	db.Preload("RoleUser.Role").Find(&user, "id = ?", newUser.ID)

	return user, nil
}

// Update a users
//
//	param c *fiber.Ctx
//	param id string
//	return models.User, error
func UserUpdate(c *fiber.Ctx, id string) (models.User, error) {
	db := database.DB

	var user models.User
	var isExists bool
	var userStructure structs.UserUpdate

	body := c.Body()

	err := json.Unmarshal(body, &userStructure)

	if err != nil {
		return user, fmt.Errorf("payload error")
	}
	email := helpers.SanitiseText(userStructure.Email)
	username := helpers.SanitiseText(userStructure.Username)
	password := helpers.SanitiseText(userStructure.Password)
	first_name := helpers.SanitiseText(userStructure.FirstName)
	last_name := helpers.SanitiseText(userStructure.LastName)
	role_id, _ := strconv.Atoi(helpers.SanitiseText(fmt.Sprint(userStructure.RoleId)))

	user, isExists = FindByEmailOrUsername(email, username)
	if !isExists {
		return user, fmt.Errorf("user is not exists")
	}

	var dataToUpdate = &models.User{
		FirstName: first_name,
		LastName:  last_name,
		Email:     email,
		Username:  username,
	}

	if password != "" {
		dataToUpdate.Password = password
	}

	db.Model(&user).Where("id = ?", id).UpdateColumns(dataToUpdate)

	// Set Role
	_ResetRole(user.ID, uint(role_id))

	db.Preload("RoleUser.Role").Find(&user, "id = ?", user.ID)

	return user, nil
}

// Delete a users
//
//	param c *fiber.Ctx
//	param id string
//	return models.User, error
func UserDestroy(c *fiber.Ctx, id string) (models.User, error) {
	db := database.DB
	var user models.User

	db.Find(&user, "id = ?", id)

	if user.ID == 0 {
		return user, fmt.Errorf("not found")
	}

	currentUserId := fmt.Sprintf("%v", c.Locals("user_id"))
	if fmt.Sprint(user.ID) == currentUserId {
		return user, fmt.Errorf("you are not allowed to delete your own account")
	}

	// To soft delete, just remove .Unscoped()
	_DeleteRolesByUser(user.ID)
	err := db.Unscoped().Delete(&user).Error

	return user, err
}

// Update user profile
//
//	param c *fiber.Ctx
//	param id string
//	return models.User, error
func UpdateProfile(c *fiber.Ctx, id string) (models.User, error) {
	db := database.DB

	var user models.User
	var userStructure structs.UserUpdate

	body := c.Body()

	err := json.Unmarshal(body, &userStructure)

	if err != nil {
		return user, fmt.Errorf("payload error")
	}
	email := helpers.SanitiseText(userStructure.Email)
	password := helpers.SanitiseText(userStructure.Password)
	first_name := helpers.SanitiseText(userStructure.FirstName)
	last_name := helpers.SanitiseText(userStructure.LastName)

	db.Find(&user, "id = ?", id)

	if user.ID == 0 {
		return user, fmt.Errorf("user is not exists")
	}

	var sendEmail bool = false
	if first_name != user.FirstName || last_name != user.LastName {
		sendEmail = true
	}

	if sendEmail {
		emailData := struct {
			NewFirstName string
			NewLastName  string
			OldFirstName string
			OldLastName  string
			AppUrl       string
		}{
			NewFirstName: first_name,
			NewLastName:  last_name,
			OldFirstName: user.FirstName,
			OldLastName:  user.LastName,
			AppUrl:       appConfig.GetEnv("APP_URL"),
		}

		errS := helpers.SendMail(email, "Your account has been updated", "updated_account", emailData)
		if errS != nil {
			fmt.Println("error send email updated_account")
		}
	}

	var mustVerifyChangeEmail bool = false
	if email != user.Email {
		mustVerifyChangeEmail = true
	}

	if mustVerifyChangeEmail {
		token := helpers.RandomNumber(6)

		verifyEmailData := struct {
			FirstName   string
			OldEmail    string
			NewEmail    string
			Token       string
			AppUrl      string
			FrontendUrl string
		}{
			FirstName:   user.FirstName,
			OldEmail:    user.Email,
			NewEmail:    email,
			Token:       string(token),
			AppUrl:      appConfig.GetEnv("APP_URL"),
			FrontendUrl: appConfig.GetEnv("FRONTEND_URL"),
		}

		db.Create(&models.ChangeEmail{
			UserId:    user.ID,
			Token:     token,
			ExpiredAt: time.Now().Add(time.Hour * 1), // 1 hour
			OldEmail:  user.Email,
			NewEmail:  email,
		})
		errS := helpers.SendMail(user.Email, "Verify your email change", "new_email_verification", verifyEmailData)
		if errS != nil {
			fmt.Println("error send email new_email_verification")
		}
	}

	var dataToUpdate = &models.User{
		FirstName: first_name,
		LastName:  last_name,
		// Email:     email,
	}

	if password != "" {
		dataToUpdate.Password = password
	}

	db.Model(&user).Where("id = ?", id).UpdateColumns(dataToUpdate)

	db.Preload("RoleUser.Role").Find(&user, "id = ?", user.ID)

	return user, nil
}

// Update email address
//
//	param c *fiber.Ctx
//	param id string
//	return models.User, error
func UpdateEmailAddress(c *fiber.Ctx, id string) (models.User, error) {
	db := database.DB

	var user models.User
	var emailStructure structs.EmailUpdate
	var token models.ChangeEmail

	body := c.Body()

	err := json.Unmarshal(body, &emailStructure)

	if err != nil {
		return user, fmt.Errorf("payload error")
	}

	tokenString := helpers.SanitiseText(emailStructure.Token)

	checkToken := db.Where("token = ?", tokenString).Where("expired_at > ?", time.Now()).Find(&token)
	if checkToken.RowsAffected == 0 {
		return user, fmt.Errorf("token is not exists")
	}

	db.Find(&user, "id = ?", token.UserId)

	if user.ID == 0 {
		return user, fmt.Errorf("user is not exists")
	}

	var dataToUpdate = &models.User{
		Email: token.NewEmail,
	}

	db.Model(&user).Where("id = ?", id).UpdateColumns(dataToUpdate)
	db.Unscoped().Delete(&models.ChangeEmail{}, "user_id = ?", user.ID)

	emailData := struct {
		FirstName   string
		NewEmail    string
		OldEmail    string
		FrontendUrl string
		AppUrl      string
	}{
		FirstName:   user.FirstName,
		NewEmail:    token.NewEmail,
		OldEmail:    token.OldEmail,
		FrontendUrl: appConfig.GetEnv("FRONTEND_URL"),
		AppUrl:      appConfig.GetEnv("APP_URL"),
	}
	errS1 := helpers.SendMail(token.NewEmail, "Your email address has been updated", "updated_email_address", emailData)
	if errS1 != nil {
		fmt.Println("error send email updated_email_address")
	}
	errS2 := helpers.SendMail(token.OldEmail, "Your email address has been updated", "updated_email_address", emailData)
	if errS2 != nil {
		fmt.Println("error send email updated_email_address")
	}

	db.Preload("RoleUser.Role").Find(&user, "id = ?", user.ID)

	return user, nil
}

func removeResetPasswordTokens(id string) {
	db := database.DB
	db.Unscoped().Where("user_id = ?", id).Where("type = ?", "reset_password").Delete(&models.UserToken{})
}
