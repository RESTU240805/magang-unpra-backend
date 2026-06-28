package handlers

import (
	"log"
	"magang-unpra-backend/config"
	"magang-unpra-backend/middleware"
	"magang-unpra-backend/models"
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type attemptInfo struct {
	count    int
	firstHit time.Time
}

var (
	loginAttempts   = make(map[string]*attemptInfo)
	loginAttemptsMu sync.Mutex
)

const maxLoginAttempts = 5
const lockoutDuration = 15 * time.Minute

func checkLoginLockout(email string) (int, bool) {
	loginAttemptsMu.Lock()
	defer loginAttemptsMu.Unlock()
	info, exists := loginAttempts[email]
	if !exists {
		return 0, false
	}
	if time.Since(info.firstHit) > lockoutDuration {
		delete(loginAttempts, email)
		return 0, false
	}
	if info.count >= maxLoginAttempts {
		return info.count, true
	}
	return info.count, false
}

func recordLoginAttempt(email string, success bool) {
	loginAttemptsMu.Lock()
	defer loginAttemptsMu.Unlock()
	if success {
		delete(loginAttempts, email)
		return
	}
	info, exists := loginAttempts[email]
	if !exists {
		loginAttempts[email] = &attemptInfo{count: 1, firstHit: time.Now()}
		return
	}
	if time.Since(info.firstHit) > lockoutDuration {
		info.count = 1
		info.firstHit = time.Now()
		return
	}
	info.count++
}

func Login(c *gin.Context) {
	var input struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	attempts, locked := checkLoginLockout(input.Email)
	if locked {
		log.Printf("Login locked for %s (%d attempts)", input.Email, attempts)
		c.JSON(http.StatusTooManyRequests, gin.H{"error": "Terlalu banyak percobaan login. Coba lagi dalam 15 menit"})
		return
	}

	var user models.User
	if err := config.DB.Where("email = ?", input.Email).First(&user).Error; err != nil {
		recordLoginAttempt(input.Email, false)
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Email atau password salah"})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
		recordLoginAttempt(input.Email, false)
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Email atau password salah"})
		return
	}

	recordLoginAttempt(input.Email, true)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"email":   user.Email,
		"role":    user.Role,
		"exp":     time.Now().Add(2 * time.Hour).Unix(),
	})

	tokenString, err := token.SignedString(middleware.GetJWTSecret())
	if err != nil {
		log.Printf("Failed to generate token: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	c.SetCookie("token", tokenString, 7200, "/", "", false, true)

	c.JSON(http.StatusOK, gin.H{
		"token": tokenString,
		"user": gin.H{
			"name":  user.Email,
			"email": user.Email,
			"role":  user.Role,
		},
	})
}
