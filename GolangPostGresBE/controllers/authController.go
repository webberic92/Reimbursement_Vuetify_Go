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
			"message": "User already exists with that Email.",
		})
	}

}

func Login(c *fiber.Ctx) error {
	var data map[string]string
	//Verifies body is not blank.
	if err := c.BodyParser(&data); err != nil {
		return err
	}
	//Get users with parsed email from database
	var user models.User
	database.DB.Where("email = ?", data["email"]).First(&user)
	//Backend validation
	if data["email"] == "" {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "Email can not be blank",
		})
	}
	//Backend validation
	if data["password"] == "" {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "password can not be blank",
		})
	}
	//Backend validation
	if user.Id == 0 {
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			"message": "Email does not exist",
		})
	}
	//Password validation
	if err := bcrypt.CompareHashAndPassword(user.Password, []byte(data["password"])); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "Wrong password",
		})
	}
	//Creates JWT
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
	//Creates Cookie.
	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 24),
		HTTPOnly: true,
	}
	//Returns Cookie with JWT and message string saying success.
	c.Cookie(&cookie)
	return c.JSON(fiber.Map{
		"message": "success",
	})
}

func GetUserById(c *fiber.Ctx) error {
	checkCookie(c)
	var user models.User
	id := string(c.Params("id"))
	fmt.Print((id))

	database.DB.Where(id).First(&user)
	if user.Id > 0 {
		return c.JSON(user)

	} else {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "User Does Not Exist",
		})
	}

}
func User(c *fiber.Ctx) error {

	checkCookie(c)
	token := getTokenFromCookie(c)
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

	checkCookie(c)
	token := getTokenFromCookie(c)
	claims := token.Claims.(*jwt.StandardClaims)

	var data map[string]string
	c.BodyParser(&data)
	if err := c.BodyParser(&data); err != nil {
		return err
	}
	//  issuerId := string(claims.Issuer)
	issuerId, err := strconv.ParseUint(claims.Issuer, 10, 64)
	if err != nil {
		return c.JSON(err)

	}
	reimbursment := models.Reimbursment{

		UserID:       issuerId,
		Title:        data["title"],
		Description:  data["description"],
		Amount:       data["amount"],
		DateApproved: "",
	}

	database.DB.Create(&reimbursment)

	return c.JSON(reimbursment)

}

func GetReimbursments(c *fiber.Ctx) error {
	checkCookie(c)
	token := getTokenFromCookie(c)
	claims := token.Claims.(*jwt.StandardClaims)
	var reimbursment []models.Reimbursment

	database.DB.Where(map[string]interface{}{"user_id": claims.Issuer, "approved_status": ""}).Find(&reimbursment)
	return c.JSON(reimbursment)

}

func GetHistory(c *fiber.Ctx) error {
	checkCookie(c)
	token := getTokenFromCookie(c)

	cols := []map[string]interface{}{}
	claims := token.Claims.(*jwt.StandardClaims)
	if claims != nil {
		fmt.Printf(claims.Issuer)
		err := database.DB.Raw("SELECT request_id, user_id, title, description, amount, approved_status, date_approved, approved_by,concat(submitter.last_name ,', ', submitter.first_name ) as submitter , concat(approver.last_name,', ',approver.first_name) as approver FROM public.reimbursments as reimbursment INNER JOIN users as approver ON approver.id = reimbursment.approved_by INNER JOIN users as submitter ON submitter.id = reimbursment.user_id where reimbursment.user_id = " + claims.Issuer).Find(&cols).Error
		if err != nil {
			fmt.Print(err)
		}
	}

	return c.JSON(&cols)

}

func GetAllOpenReimbursments(c *fiber.Ctx) error {
	checkCookie(c)
	token := getTokenFromCookie(c)
	claims := token.Claims.(*jwt.StandardClaims)
	cols := []map[string]interface{}{}

	if claims != nil {
		err := database.DB.Raw(" SELECT request_id, user_id, title, description, amount,concat(submitter.last_name ,', ', submitter.first_name ) as submitter FROM public.reimbursments as reimbursment INNER JOIN users as submitter ON submitter.id = reimbursment.user_id where approved_status = ''").Find(&cols).Error
		if err != nil {
			return c.JSON(err)
		}
	}
	return c.JSON(&cols)
}

func ApproveOrDeny(c *fiber.Ctx) error {
	checkCookie(c)
	token := getTokenFromCookie(c)
	claims := token.Claims.(*jwt.StandardClaims)
	user := models.User{
		UserType: "",
	}
	database.DB.Model(&models.User{}).Where("id = ?", claims.Issuer).Find(&user)

	if user.UserType == "R" {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "User is NOT an ADMIN.",
		})
	}
	var data map[string]string
	c.BodyParser(&data)
	if err := c.BodyParser(&data); err != nil {
		return err
	}

	if data["approveOrDeny"] != "D" && data["approveOrDeny"] != "A" {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "Approve or Deny can only be D or A.",
		})
	}
	i, err := strconv.ParseUint(claims.Issuer, 10, 64)
	if err != nil {
		fmt.Printf("Cant parse to an Uint64: %T \n", i)
	}

	timeStamp := time.Now().Format("01-02-2006")
	reimbursment := models.Reimbursment{

		RequestId:      0,
		ApprovedStatus: data["approveOrDeny"],
		DateApproved:   timeStamp,
		ApprovedBy:     i,
		UserID:         0,
		Title:          "",
		Description:    "",
		Amount:         "",
	}

	database.DB.Model(&models.Reimbursment{}).Where("request_id = ?", data["requestId"]).Find(&reimbursment)
	if reimbursment.RequestId == 0 {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "Request ID does not exsist.",
		})
	}
	if reimbursment.ApprovedStatus == "A" || reimbursment.ApprovedStatus == "D" {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "Request has allready been approved or denied.",
		})
	}
	database.DB.Model(&models.Reimbursment{}).Where("request_id = ?", data["requestId"]).Update("approved_status", data["approveOrDeny"]).Update("approved_by", claims.Issuer).Update("date_approved", timeStamp).Find(&reimbursment)

	return c.JSON(reimbursment)

}
func GetAllHistory(c *fiber.Ctx) error {
	checkCookie(c)
	token := getTokenFromCookie(c)
	claims := token.Claims.(*jwt.StandardClaims)
	cols := []map[string]interface{}{}

	if claims != nil {
		err := database.DB.Raw("SELECT request_id, user_id, title, description, amount, approved_status, date_approved,approved_by,concat(submitter.last_name ,', ', submitter.first_name ) as submitter, concat(approver.last_name,', ',approver.first_name) as approver FROM public.reimbursments as reimbursment INNER JOIN users as approver ON approver.id = reimbursment.approved_by INNER JOIN users as submitter ON submitter.id = reimbursment.user_id").Find(&cols).Error
		if err != nil {
			return c.JSON(err)
		}
	}
	return c.JSON(&cols)
}

func checkCookie(c *fiber.Ctx) error {
	cookie := c.Cookies("jwt")

	_, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})

	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "Unauthenticated User",
		})
	}

	return nil
}

func getTokenFromCookie(c *fiber.Ctx) *jwt.Token {
	cookie := c.Cookies("jwt")

	token, _ := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})

	return token
}
