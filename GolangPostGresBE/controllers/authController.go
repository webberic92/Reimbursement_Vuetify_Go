package controllers

import (
	"fmt"
	"strconv"
	"time"

	"example.com/m/database"
	"example.com/m/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

const secretKey = "adfasdfasdfasdf!@#$!#@$12342134"

func Register(c *fiber.Ctx) error {

	var data map[string]string
	c.BodyParser(&data)
	if err := c.BodyParser(&data); err != nil {
		return err
	}

	sou, err := strconv.ParseBool(data["sou"])

	if err != nil {
		c.Status(fiber.StatusBadRequest)

		return c.JSON(fiber.Map{
			"message": "Cant parse boolean for sou",
		})
	}
	password, _ := bcrypt.GenerateFromPassword([]byte(data["password"]), 14)
	user := models.User{
		FirstName:   data["firstName"],
		LastName:    data["lastName"],
		Email:       data["email"],
		PhoneNumber: data["phoneNumber"],
		UserType:    data["userType"],
		Sou:         sou,
		Password:    password,
	}

	if data["firstName"] == "" {
		c.Status(fiber.StatusBadRequest)

		return c.JSON(fiber.Map{
			"message": "No first name provided.",
		})
	}

	if data["lastName"] == "" {
		c.Status(fiber.StatusBadRequest)

		return c.JSON(fiber.Map{
			"message": "No last name provided.",
		})
	}

	if data["email"] == "" {
		c.Status(fiber.StatusBadRequest)

		return c.JSON(fiber.Map{
			"message": "No email provided",
		})
	}
	if data["phoneNumber"] == "" {
		c.Status(fiber.StatusBadRequest)

		return c.JSON(fiber.Map{
			"message": "No phone number provided",
		})
	}
	if data["password"] == "" {
		c.Status(fiber.StatusBadRequest)

		return c.JSON(fiber.Map{
			"message": "No password provided.",
		})
	}

	if data["userType"] == "" {
		c.Status(fiber.StatusBadRequest)

		return c.JSON(fiber.Map{
			"message": "No user type provided.",
		})
	}

	if data["sou"] == "" {
		c.Status(fiber.StatusBadRequest)

		return c.JSON(fiber.Map{
			"message": "No statement of understanding provided.",
		})
	}

	if err := database.DB.Where("email = ?", data["email"]).First(&user).Error; err != nil {

		database.DB.Create(&user)
		return c.JSON(fiber.Map{
			"message": "You Successfully created a new user.",
		})

	} else {

		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "User already exsists with that Email.",
		})
	}

}

func Login(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}
	var user models.User
	database.DB.Where("email = ?", data["email"]).First(&user)

	if data["email"] == "" {
		c.Status(fiber.StatusBadRequest)

		return c.JSON(fiber.Map{
			"message": "Email can not be blank",
		})
	}

	if data["password"] == "" {
		c.Status(fiber.StatusBadRequest)

		return c.JSON(fiber.Map{
			"message": "password can not be blank",
		})
	}

	if user.Id == 0 {

		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			"message": "Email does not exist",
		})
	}

	if err := bcrypt.CompareHashAndPassword(user.Password, []byte(data["password"])); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{

			"message": "Wrong password",
		})

	}

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer:    strconv.Itoa(int(user.Id)),
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
	})

	token, err := claims.SignedString([]byte(secretKey))

	if err != nil {

		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": "Could not login",
		})
	}

	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 24),
		HTTPOnly: true,
	}
	c.Cookie(&cookie)
	return c.JSON(fiber.Map{
		"message": "success",
	})
}

func User(c *fiber.Ctx) error {
	cookie := c.Cookies("jwt")

	token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})

	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "Unauthenticated User",
		})
	}
	claims := token.Claims.(*jwt.StandardClaims)
	var user models.User

	database.DB.Where("id = ?", claims.Issuer).First(&user)

	return c.JSON(user)

}

func Logout(c *fiber.Ctx) error {

	User(c)

	if c.Response().StatusCode() == 401 {
		return c.JSON(fiber.Map{
			"message": "Cant logout, User not even logged in.",
		})
	} else {
		cookie := fiber.Cookie{
			Name:     "jwt",
			Value:    "",
			Expires:  time.Now().Add(-time.Hour),
			HTTPOnly: true,
		}

		c.Cookie(&cookie)
		return c.JSON(fiber.Map{
			"message": "Logout successful",
		})
	}

}

func CreateReimbursment(c *fiber.Ctx) error {

	cookie := c.Cookies("jwt")

	token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})

	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "Unauthenticated User",
		})
	}
	claims := token.Claims.(*jwt.StandardClaims)

	var data map[string]string
	c.BodyParser(&data)
	if err := c.BodyParser(&data); err != nil {
		return err
	}

	reimbursment := models.Reimbursment{

		UserID:         claims.Issuer,
		Title:          data["title"],
		Description:    data["description"],
		Amount:         data["amount"],
		ApprovedStatus: "",
		DateApproved:   "",
		ApprovedBy:     "",
	}

	database.DB.Create(&reimbursment)

	return c.JSON(reimbursment)

}

func GetReimbursments(c *fiber.Ctx) error {
	cookie := c.Cookies("jwt")

	token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})

	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "Unauthenticated User",
		})
	}
	claims := token.Claims.(*jwt.StandardClaims)
	var reimbursment []models.Reimbursment

	database.DB.Where(map[string]interface{}{"user_id": claims.Issuer, "approved_status": ""}).Find(&reimbursment)
	return c.JSON(reimbursment)

}

func GetHistory(c *fiber.Ctx) error {
	cookie := c.Cookies("jwt")

	token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})

	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "Unauthenticated User",
		})
	}
	claims := token.Claims.(*jwt.StandardClaims)
	var reimbursment []models.Reimbursment

	database.DB.Where(map[string]interface{}{"user_id": claims.Issuer, "approved_status": "A"}).Or(map[string]interface{}{"user_id": claims.Issuer, "approved_status": "D"}).Find(&reimbursment)
	return c.JSON(reimbursment)

}

func GetAllOpenReimbursments(c *fiber.Ctx) error {
	cookie := c.Cookies("jwt")

	token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})

	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "Unauthenticated User",
		})
	}
	claims := token.Claims.(*jwt.StandardClaims)
	var reimbursment []models.Reimbursment

	if claims != nil {
		database.DB.Where(map[string]interface{}{"approved_status": ""}).Find(&reimbursment)
	}
	return c.JSON(reimbursment)
}

func ApproveOrDeny(c *fiber.Ctx) error {

	cookie := c.Cookies("jwt")

	token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})

	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "Unauthenticated User",
		})
	}
	claims := token.Claims.(*jwt.StandardClaims)

	var data map[string]string
	c.BodyParser(&data)
	if err := c.BodyParser(&data); err != nil {
		return err
	}

	// requestId,err := strconv.Atoi(data["requestId"])
	requestId, err := strconv.ParseUint(data["requestId"], 10, 32)
	if err != nil {
		fmt.Println(err)
	}
	rID := uint(requestId)

	reimbursment := models.Reimbursment{

		RequestId:      rID,
		ApprovedStatus: "", // need date from fe hre should be data["approveOrDeny"]
		DateApproved:   "", //Need todays date here.
		ApprovedBy:     claims.Issuer,
	}

	database.DB.Create(&reimbursment)

	return c.JSON(reimbursment)

}
