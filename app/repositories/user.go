package repositories_v1

import (
	"encoding/json"
	"fmt"
	"goshaka/app/models"
	"goshaka/app/structs"
	"goshaka/database"
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

	errHash := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))

	if errHash != nil {
		fmt.Print(user.Password)
		return user, "", fmt.Errorf("credential cannot be found")
	}

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = user.Username
	claims["id"] = user.ID
	claims["email"] = user.Email
	claims["exp"] = time.Now().Add(time.Hour * 24 * 30).Unix()

	secret := []byte(appConfig.GetEnv("JWT_KEY"))

	tokenString, err := token.SignedString(secret)

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
	password := sanitise.Sanitize(userCreateStructure.Password)
	first_name := sanitise.Sanitize(userCreateStructure.FirstName)
	last_name := sanitise.Sanitize(userCreateStructure.LastName)

	db := database.DB

	//Set User
	db.Create(&models.User{
		FirstName: first_name,
		LastName:  last_name,
		Password:  password,
		Email:     email,
	})

	db.Preload("RoleUser.Role").Find(&user, "email = ?", email)
	//Set Role
	db.Create(&models.RoleUser{
		UserId: user.ID,
		RoleId: 1,
	})

	return user, nil
}
