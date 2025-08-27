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
// @Summary Rejestracja użytkownika
// @Description Tworzy nowego użytkownika
// @Tags auth
// @Accept json
// @Produce json
// @Param user body handlers.Credentials true "User credentials"
// @Success 201 {object} map[string]string
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
// @Summary Logowanie użytkownika
// @Description Loguje użytkownika i zwraca token JWT
// @Tags auth
// @Accept json
// @Produce json
// @Param user body handlers.Credentials true "User credentials"
// @Success 200 {object} map[string]string
// @Failure 401 {object} map[string]string
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
// @Summary Wylogowanie użytkownika
// @Description Wylogowuje użytkownika (stateless, JWT token client-side)
// @Tags auth
// @Produce json
// @Success 200 {object} map[string]string
// @Router /auth/logout [post]
func Logout(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "logged out"})
}
