package repositories_v1

import (
	"encoding/json"
	"fmt"
	"goshaka/app/models"
	"goshaka/app/structs"
	"goshaka/database"
	"goshaka/helpers"
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
	var existingUser models.User
	checkUser := db.Where("email = ?", email).Or("username = ?", username).Find(&existingUser)
	if checkUser.RowsAffected > 0 {
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

	db.Preload("RoleUser.Role").Find(&user, "email = ?", email)
	//Set Role
	db.Create(&models.RoleUser{
		UserId: user.ID,
		RoleId: 1,
	})

	return user, nil
}
