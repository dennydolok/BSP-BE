package Controller

import (
	"dennydolok/BSP-BE/Helper"
	"dennydolok/BSP-BE/config"
	"dennydolok/BSP-BE/model/entitas"
	"encoding/base64"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

func Login(c echo.Context) error {
	var users entitas.User
	var payload entitas.Login
	err := c.Bind(&payload)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err,
		})
	}
	config.PrintLog(payload.Email)
	err = config.DB.First(&users, "email = ?", payload.Email).Error
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"message": "username not found",
		})
	}
	password := base64.StdEncoding.EncodeToString([]byte(payload.Password))
	config.PrintLog(password)
	config.PrintLog(users.Password)
	if users.Password != password {
		return c.JSON(http.StatusUnauthorized, map[string]interface{}{
			"message": "password incorrect",
		})
	}
	token, err := helper.CreateToken(users.Role, users.ID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error":   err,
			"message": "failed to generate token",
		})
	}
	users.Password = ""
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"token":   token,
	})
}

func Register(c echo.Context) error {
	var payload entitas.Register
	err := c.Bind(&payload)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err,
		})
	}
	if payload.Password != payload.VerifyPassword || payload.Email != payload.VerifyEmail {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "password or email doesn't match",
		})
	}
	var existUser []entitas.User
	check := config.DB.Find(&entitas.User{}).Where("email = ?", payload.Email).Scan(&existUser)
	if check.Error != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": check.Error,
		})
	}
	if len(existUser) > 0 {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "email already exist",
		})
	}

	password := base64.StdEncoding.EncodeToString([]byte(payload.Password))
	user := entitas.User{
		Email:    payload.Email,
		Password: password,
		Nama:     payload.Nama,
		Role:     2,
	}
	err = config.DB.Save(&user).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err,
		})
	}
	return c.JSON(http.StatusCreated, map[string]interface{}{
		"message": "success",
	})
}

func Update(c echo.Context) error {
	authCheck := helper.GetClaimsRole(c.Request().Header.Get("Authorization"))
	if authCheck == 0 {
		return c.JSON(http.StatusUnauthorized, map[string]interface{}{
			"message": "unauthorized",
		})
	}
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "bad request",
		})
	}
	var payload entitas.Register
	err = c.Bind(&payload)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err,
		})
	}
	if payload.Password != payload.VerifyPassword || payload.Email != payload.VerifyEmail {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "password or email doesn't match",
		})
	}
	password := base64.StdEncoding.EncodeToString([]byte(payload.Password))
	var existUser entitas.User
	check := config.DB.First(&existUser, "id = ?", id)
	if check.Error != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": check.Error,
		})
	}
	err = config.DB.Model(&entitas.User{}).Where("id = ?", id).Updates(entitas.User{
		Email:    payload.Email,
		Password: password,
		Nama:     payload.Nama,
	}).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err,
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
	})
}
