package controller

import (
	"api_rpm/model"
	"api_rpm/utils"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)
var SECRET = []byte("secret-auth-key")

const api_key = "rpm_express_shipment"
const jwtSecret = "indonesia"
const content_type = "application/json"

func HashAccount(c *gin.Context) {
	var body model.Users
	if err := c.ShouldBindJSON(&body); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// Generate hash of the password
	bytes, err := bcrypt.GenerateFromPassword([]byte(body.Password), 14)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	// Update the body variable with the hashed password
	body.Password = string(bytes)

	// Insert the user into the database
	result := model.DB.Table("users").Create(&body)
	if result.Error != nil {
		log.Printf("Error creating user: %v", result.Error)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": result.Error})
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}

func HashPassword(password string) (string, error) {

	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(hashedPwd string, plainPwd []byte) bool {
	bytheHash := []byte(hashedPwd)
	err := bcrypt.CompareHashAndPassword(bytheHash, plainPwd)
	return err == nil
}

func verifyUser(username, password string) (*model.Users, error) {

	var user model.Users
	// hash, _ := HashPassword(password)

	result := model.DB.Raw("SELECT id, username, password FROM users WHERE username = ?", username).Scan(&user)

	if user.Username == "" {
		return nil, fmt.Errorf("Username is empty")
	}

	if user.Password == "" {
		return nil, fmt.Errorf("Password is empty")
	}

	if err := result.Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("user not found")
		}

		return nil, fmt.Errorf("error querying database")
	}

	hashedPwd := user.Password
	match := CheckPasswordHash(hashedPwd, []byte(password))
	if !match {
		fmt.Println(user.Password)

		fmt.Println(match)
		return nil, fmt.Errorf("invalid password")
	}

	return &user, nil
}

func ValidateJWT(next func(c *gin.Context)) gin.HandlerFunc {
	return func(c *gin.Context) {

		tokenEmpty := utils.TokenEmpty()
		statusUnauthorized := utils.StatusUnauthorized()

		tokenHeader := c.GetHeader("Token")
		if tokenHeader == "" {
			c.JSON(http.StatusUnauthorized, tokenEmpty)
			c.Abort()
			return
		}

		token, err := jwt.Parse(tokenHeader, func(t *jwt.Token) (interface{}, error) {
			_, ok := t.Method.(*jwt.SigningMethodHMAC)
			if !ok {
				return nil, fmt.Errorf("unexpected signing method : %v", t.Header["alg"])
			}

			return []byte(jwtSecret), nil
		})

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, statusUnauthorized)
			c.Abort()
			return
		}

		next(c)
	}
}

func LoginHandler(c *gin.Context) {

	tokenEmpty := utils.TokenEmpty()
	// statusUnauthorized := utils.StatusUnauthorized()

	var req model.LoginRequest
	if c.GetHeader("Access") == api_key && c.GetHeader("Content-Type") == content_type {
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
	} else {

		c.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid Header"})
	}

	user, err := verifyUser(req.Username, req.Password)

	if err != nil {
		apiError := &APIError{}
		switch err.Error() {
		case "username is empty":
			apiError.StatusCode = http.StatusBadRequest
			apiError.Message = "Username is empty"
		case "user not found":
			apiError.StatusCode = http.StatusNotFound
			apiError.Message = "User not found"
		case "password is empty":
			apiError.StatusCode = http.StatusBadRequest
			apiError.Message = "Password is empty"
		case "invalid password":
			apiError.StatusCode = http.StatusBadRequest
			apiError.Message = "Invalid password"
		default:
			apiError.StatusCode = http.StatusNotFound
			apiError.Message = "Akun tidak ditemukan"
		}
		c.JSON(apiError.StatusCode, gin.H{"error": apiError.Message})
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.Id,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})
	tokenString, err := token.SignedString([]byte(jwtSecret))
	if err != nil {
		c.JSON(http.StatusInternalServerError, tokenEmpty)
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": tokenString})
}

func HandleUpload(c *gin.Context) {
	// Get the image file from the form data
	file, err := c.FormFile("image")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// Save the image file to disk
	err = c.SaveUploadedFile(file, "images/"+file.Filename)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	// Return a success response
	c.JSON(http.StatusOK, gin.H{
		"message": "Image uploaded successfully",
	})
}

func HandleImage(c *gin.Context) {
	// Get the image ID from the URL
	id := c.Param("id")

	// Read the image file from disk
	image, err := ioutil.ReadFile("images/" + id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Image not found",
		})
		return
	}

	// Set the response headers and return the image file
	c.Header("Content-Type", "image/jpeg")
	c.Header("Content-Disposition", "attachment; filename="+id)
	c.Data(http.StatusOK, "image/jpeg", image)
}

type APIError struct {
	StatusCode int    `json:"statusCode"`
	Message    string `json:"message"`
}

func (e *APIError) Error() string {
	return e.Message
}
