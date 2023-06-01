package controllers

import (
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"github.com/sautsihotang/DatingApp/db"
	"github.com/sautsihotang/DatingApp/models"
	"golang.org/x/crypto/bcrypt"
)

var signingKey = []byte("SangatRahasia")

func Register(c echo.Context) error {
	user := new(models.User)
	if err := c.Bind(user); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	//generate has
	hasPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	user.Password = string(hasPassword)

	err = db.SaveUser(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, user)
}

func Login(c echo.Context) error {
	userLogin := new(models.Login)
	if err := c.Bind(userLogin); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	username := userLogin.Username
	password := userLogin.Password

	user, err := db.GetUserByUsername(username)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, err)
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return c.JSON(http.StatusUnauthorized, err)
	}

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["userId"] = user.UserID
	claims["username"] = user.Username
	claims["email"] = user.Email
	claims["premiumStatus"] = user.PremiumStatus
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix() // 1 hari

	tokenString, err := token.SignedString(signingKey)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, map[string]string{
		"status": "Berhasil Login",
		"token":  tokenString,
	})
}
