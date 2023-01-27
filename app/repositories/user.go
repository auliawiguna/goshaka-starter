package repositories_v1

import (
	"encoding/json"
	"fmt"
	"goshaka/app/models"
	"goshaka/app/models/scopes"
	"goshaka/app/structs"
	"goshaka/database"
	"goshaka/helpers"
	"strconv"
	"time"

	appConfig "goshaka/configs"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"github.com/microcosm-cc/bluemonday"
	"golang.org/x/crypto/bcrypt"
)

func UserShow(id string) models.User {
	db := database.DB
	var user models.User

	db.Preload("RoleUser.Role").Find(&user, "id = ?", id)

	return user
}

func Login(c *fiber.Ctx) (models.User, string, error) {
	sanitise := bluemonday.UGCPolicy()

	var user models.User
	var loginStructure structs.Login

	body := c.Body()

	err := json.Unmarshal(body, &loginStructure)

	if err != nil {
		return user, "", fmt.Errorf("payload error")
	}
	email := sanitise.Sanitize(loginStructure.Email)
	password := sanitise.Sanitize(loginStructure.Password)

	db := database.DB
	db.Preload("RoleUser.Role").Find(&user, "email = ?", email)

	errHash := helpers.CompareHash(user.Password, password)

	if !errHash {
		return user, "", fmt.Errorf("credential cannot be found")
	}

	return GenerateJwt(user)
}

func ValidateRegistration(c *fiber.Ctx) (models.User, string, error) {
	sanitise := bluemonday.UGCPolicy()

	var user models.User
	var loginStructure structs.RegistrationToken

	body := c.Body()

	err := json.Unmarshal(body, &loginStructure)

	if err != nil {
		return user, "", fmt.Errorf("payload error")
	}
	email := sanitise.Sanitize(loginStructure.Email)
	password := sanitise.Sanitize(loginStructure.Password)
	tokenPayload := sanitise.Sanitize(loginStructure.Token)

	db := database.DB
	db.Preload("RoleUser.Role").Find(&user, "email = ?", email)

	errHash := helpers.CompareHash(user.Password, password)

	if !errHash {
		return user, "", fmt.Errorf("credential cannot be found")
	}

	//Find the token
	var token models.UserToken
	checkToken := db.Where("user_id = ?", user.ID).Where("type = ?", "registration").Where("expired_at > ?", time.Now()).First(&token)
	if checkToken.RowsAffected == 0 {
		return user, "", fmt.Errorf("token is not found")
	}

	//Verify token
	errToken := helpers.CompareHash(token.Token, tokenPayload)
	if !errToken {
		return user, "", fmt.Errorf("credential cannot be found")
	}

	//Delete token
	db.Unscoped().Delete(&token)

	//Validate user
	db.Model(&user).Update("validated_at", time.Now())

	//Generate token
	return GenerateJwt(user)
}

func GenerateJwt(user models.User) (models.User, string, error) {
	tokenData := jwt.New(jwt.SigningMethodHS256)

	claims := tokenData.Claims.(jwt.MapClaims)
	claims["username"] = user.Username
	claims["id"] = user.ID
	claims["email"] = user.Email
	claims["exp"] = time.Now().Add(time.Hour * 24 * 30).Unix()

	secret := []byte(appConfig.GetEnv("JWT_KEY"))

	tokenString, err := tokenData.SignedString(secret)

	if err != nil {
		return user, "", fmt.Errorf("error when generate JWT")
	}

	return user, tokenString, nil
}

func _CheckUserByEmailOrUsername(email string, username string) (models.User, bool) {
	db := database.DB

	var existingUser models.User
	checkUser := db.Where("email = ?", email).Or("username = ?", username).Find(&existingUser)
	if checkUser.RowsAffected > 0 {
		return existingUser, true
	}
	return existingUser, false
}

func _DeleteRolesByUser(userId uint) {
	db := database.DB
	db.Unscoped().Delete(&models.RoleUser{}, "user_id = ?", userId)
}

func _ResetRole(userId uint, roleId uint) {
	db := database.DB
	_DeleteRolesByUser(userId)
	db.Create(&models.RoleUser{
		UserId: userId,
		RoleId: roleId,
	})
}

func _SetRole(userId uint, roleId uint) {
	db := database.DB
	db.Create(&models.RoleUser{
		UserId: userId,
		RoleId: roleId,
	})
}

func Register(c *fiber.Ctx) (models.User, error) {
	sanitise := bluemonday.UGCPolicy()

	var user models.User
	var userCreateStructure structs.UserCreate

	body := c.Body()

	err := json.Unmarshal(body, &userCreateStructure)

	if err != nil {
		return user, fmt.Errorf("payload error")
	}
	email := sanitise.Sanitize(userCreateStructure.Email)
	username := sanitise.Sanitize(userCreateStructure.Username)
	password := sanitise.Sanitize(userCreateStructure.Password)
	first_name := sanitise.Sanitize(userCreateStructure.FirstName)
	last_name := sanitise.Sanitize(userCreateStructure.LastName)

	db := database.DB

	//Check the existence first
	_, isExists := _CheckUserByEmailOrUsername(email, username)
	if isExists {
		return user, fmt.Errorf("registration is failed")
	}

	//Set User
	newUser := models.User{
		FirstName: first_name,
		LastName:  last_name,
		Password:  password,
		Email:     email,
		Username:  username,
	}
	db.Create(&newUser)

	token := helpers.RandomNumber(6)
	hashedToken, _ := bcrypt.GenerateFromPassword([]byte(token), bcrypt.DefaultCost)

	emailData := struct {
		FirstName string
		Token     string
	}{
		FirstName: first_name,
		Token:     string(token),
	}

	db.Create(&models.UserToken{
		UserId:    newUser.ID,
		Type:      "registration",
		Token:     string(hashedToken),
		ExpiredAt: time.Now().Add(time.Hour * 72), // 3days
	})

	helpers.SendMail(email, "Complete your registration", "registration", emailData)

	//Set Role
	_SetRole(newUser.ID, 1)

	db.Preload("RoleUser.Role").Find(&user, "id = ?", newUser.ID)

	return user, nil
}

func UserShowAll(pagination helpers.Pagination) (*helpers.Pagination, bool) {
	db := database.DB
	var users []*models.User
	var error bool = false

	db.Scopes(scopes.Paginate(users, &pagination, db)).Find(&users)
	pagination.Rows = users

	return &pagination, error
}

func UserCreate(c *fiber.Ctx) (models.User, error) {
	db := database.DB
	sanitise := bluemonday.UGCPolicy()

	var user models.User
	var isExists bool
	var userCreateStructure structs.UserCreate

	body := c.Body()

	err := json.Unmarshal(body, &userCreateStructure)

	if err != nil {
		return user, fmt.Errorf("payload error")
	}
	email := sanitise.Sanitize(userCreateStructure.Email)
	username := sanitise.Sanitize(userCreateStructure.Username)
	password := sanitise.Sanitize(userCreateStructure.Password)
	first_name := sanitise.Sanitize(userCreateStructure.FirstName)
	last_name := sanitise.Sanitize(userCreateStructure.LastName)
	role_id, _ := strconv.Atoi(sanitise.Sanitize(fmt.Sprint(userCreateStructure.RoleId)))

	user, isExists = _CheckUserByEmailOrUsername(email, username)
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

	//Set Role
	_SetRole(newUser.ID, uint(role_id))

	db.Preload("RoleUser.Role").Find(&user, "id = ?", newUser.ID)

	return user, nil
}

func UserUpdate(c *fiber.Ctx, id string) (models.User, error) {
	db := database.DB
	sanitise := bluemonday.UGCPolicy()

	var user models.User
	var isExists bool
	var userStructure structs.UserUpdate

	body := c.Body()

	err := json.Unmarshal(body, &userStructure)

	if err != nil {
		return user, fmt.Errorf("payload error")
	}
	email := sanitise.Sanitize(userStructure.Email)
	username := sanitise.Sanitize(userStructure.Username)
	password := sanitise.Sanitize(userStructure.Password)
	first_name := sanitise.Sanitize(userStructure.FirstName)
	last_name := sanitise.Sanitize(userStructure.LastName)
	role_id, _ := strconv.Atoi(sanitise.Sanitize(fmt.Sprint(userStructure.RoleId)))

	user, isExists = _CheckUserByEmailOrUsername(email, username)
	if !isExists {
		return user, fmt.Errorf("user is not exists")
	}

	db.Model(&user).Where("id = ?", id).UpdateColumns(&models.User{
		FirstName: first_name,
		LastName:  last_name,
		Password:  password,
		Email:     email,
		Username:  username,
	})

	//Set Role
	_ResetRole(user.ID, uint(role_id))

	db.Preload("RoleUser.Role").Find(&user, "id = ?", user.ID)

	return user, nil
}

func UserDestroy(c *fiber.Ctx, id string) (models.User, error) {
	db := database.DB
	var user models.User

	db.Find(&user, "id = ?", id)

	if user.ID == 0 {
		return user, fmt.Errorf("not found")
	}

	if user.ID == c.Locals("user_id") {
		return user, fmt.Errorf("you are not allowed to delete your own account")
	}

	//To soft delete, just remove .Unscoped()
	_DeleteRolesByUser(user.ID)
	err := db.Unscoped().Delete(&user).Error

	return user, err
}
