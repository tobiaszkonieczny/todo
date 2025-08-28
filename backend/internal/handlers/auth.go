package handlers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/tobiaszkonieczny/todo.git/internal/models"
	"golang.org/x/crypto/bcrypt"
)

var jwtKey = []byte("your_secret_key") // could be from env variable

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// Register godoc
// @Summary Register new user
// @Description Create a new user account with username and password. Password will be hashed before storing.
// @Tags auth
// @Accept json
// @Produce json
// @Param user body handlers.Credentials true "User registration credentials"
// @Success 201 {object} map[string]string "User created successfully" example({"message": "user created"})
// @Failure 400 {object} map[string]string "Bad request - invalid input" example({"error": "invalid JSON format"})
// @Router /auth/register [post]
func Register(c *gin.Context) {
	var creds Credentials
	if err := c.ShouldBindJSON(&creds); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(creds.Password), bcrypt.DefaultCost)

	user := models.User{
		Username: creds.Username,
		Password: string(hashedPassword),
	}

	models.DB.Create(&user)
	c.JSON(http.StatusCreated, gin.H{"message": "user created"})
}

// Login godoc
// @Summary User login
// @Description Authenticate user with username and password. Returns JWT token for authenticated requests.
// @Tags auth
// @Accept json
// @Produce json
// @Param user body handlers.Credentials true "User login credentials"
// @Success 200 {object} map[string]string "Login successful with JWT token" example({"token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."})
// @Failure 400 {object} map[string]string "Bad request - invalid input" example({"error": "invalid JSON format"})
// @Failure 401 {object} map[string]string "Unauthorized - invalid credentials" example({"error": "invalid credentials"})
// @Router /auth/login [post]
func Login(c *gin.Context) {
	var creds Credentials
	if err := c.ShouldBindJSON(&creds); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user models.User
	if err := models.DB.Where("username = ?", creds.Username).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(creds.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"exp":     time.Now().Add(time.Hour * 24).Unix(), //24h expiration
	})

	tokenString, _ := token.SignedString(jwtKey)
	c.JSON(http.StatusOK, gin.H{"token": tokenString})
}

// Logout godoc
// @Summary User logout
// @Description Logout user (stateless operation as JWT tokens are managed client-side). No server-side session invalidation needed.
// @Tags auth
// @Produce json
// @Security BearerAuth
// @Success 200 {object} map[string]string "Logout successful" example({"message": "logged out"})
// @Router /auth/logout [post]
func Logout(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "logged out"})
}
