package controllers

import (
	"fmt"
	"strconv"
	"time"

	"github.com/devesh/golang-react-jwt/database"
	"github.com/devesh/golang-react-jwt/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

const SecretKey = "mySecret"

func Register(c *fiber.Ctx) error {
	var data map[string]string

	// extarcting the body data from request and storing into data map
	if err := c.BodyParser(&data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid input",
		})
	}

	// validating empty input fields
	requiredFields := []string{"name", "email", "password"}
	for _, field := range requiredFields {
		if data[field] == "" {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": fmt.Sprintf("%s is required", field),
			})
		}
	}

	// encrypting the password by hashing
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(data["password"]), 14)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to hash password",
		})
	}

	// Checking if the user already exists
	var existingUser models.User

	database.DB.Where("email = ?", data["email"]).First(&existingUser)
	if existingUser.Id != 0 {
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{
			"message": "User already exists please go for Login !",
		})
	} else {
		user := models.User{
			Name:     data["name"],
			Email:    data["email"],
			Password: hashedPassword,
		}
		if err := database.DB.Create(&user).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Failed to create user",
			})
		}

		return c.JSON(user)
	}
}

func Login(c *fiber.Ctx) error {
	// extarcting json body from request
	var data map[string]string
	if err := c.BodyParser(&data); err != nil {
		return err
	}

	// validating empty input fields
	requiredFields := []string{"email", "password"}
	for _, field := range requiredFields {
		if data[field] == "" {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": fmt.Sprintf("%s is required", field),
			})
		}
	}

	// checking if user registred or not
	var user models.User
	database.DB.Where("email = ?", data["email"]).First(&user)
	if user.Id == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Sorry, you are not registered with this email please go to Register !",
		})
	}

	// if user is registered then proceed for login and comparing  the user provided password and password saved in database.
	if err := bcrypt.CompareHashAndPassword(user.Password, []byte(data["password"])); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "Incorrect password Please type again !",
		})
	}

	// creating the standard claims
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer:    strconv.Itoa(int(user.Id)),
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(), //1 day
	})

	// create jwt token
	token, err := claims.SignedString([]byte(SecretKey))
	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": "could not login",
		})
	}

	// creating  cookies token struct
	cookie := fiber.Cookie{
		Name:     "jwtToken",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 6),
		HTTPOnly: true,
	}

	c.Cookie(&cookie) // sending token with response to cookies

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": " Login Success !",
	})
}

func GetUser(c *fiber.Ctx) error {
	// Check for missing cookie
	cookie := c.Cookies("jwtToken")
	if cookie == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": " User is currently not logged-In !",
		})
	}

	// Parse JWT token with enhanced error handling
	token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(SecretKey), nil
	})

	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "unauthenticated",
		})
	}

	/*
		if err != nil {
			switch err := err.(type) {
			case *jwt.ValidationError:
				var msg string
				if err.Errors&jwt.ValidationErrorMalformed == jwt.ValidationErrorMalformed {
					msg = "Invalid token format"
				} else if err.Errors&jwt.ValidationErrorExpired == jwt.ValidationErrorExpired {
					msg = "Expired token"
				} else {
					msg = "JWT validation error"
				}
				return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
					"message": msg,
				})

			default:
				return err // Handle other errors
			}
		}
	*/

	// Extract user information
	claims := token.Claims.(*jwt.StandardClaims)
	var user models.User
	database.DB.Where("id = ?", claims.Issuer).First(&user)

	return c.JSON(user)

}

func Logout(c *fiber.Ctx) error {
	// Get the cookie from the request
	cookie := c.Cookies("jwtToken")

	// Check if the cookie exists
	if cookie == "" {
		return c.JSON(fiber.Map{
			"message": "Already loggedout User !",
		})
	}

	expiredCookies := fiber.Cookie{
		Name:     "jwtToken",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HTTPOnly: true,
	}

	c.Cookie(&expiredCookies)

	return c.JSON(fiber.Map{
		"message": "Logged Out successfully !",
	})
}
