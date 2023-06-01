package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sautsihotang/DatingApp/db"
	"github.com/sautsihotang/DatingApp/models"
	"golang.org/x/crypto/bcrypt"
)

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
