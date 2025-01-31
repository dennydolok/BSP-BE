package Controller

import (
	"dennydolok/BSP-BE/config"
	"dennydolok/BSP-BE/model/entitas"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

func AddTipeBangunan(c echo.Context) error {
	data := entitas.TipeBangunan{}
	err := c.Bind(&data)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}
	err = config.DB.Save(&data).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}
	return c.JSON(http.StatusCreated, map[string]interface{}{
		"message": "success",
	})
}

func GetTipeBangunan(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "bad request",
		})
	}
	data := entitas.TipeBangunan{}
	err = config.DB.Find(&data, id).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"data":    data,
	})
}

func UpdateTipeBangunan(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "bad request",
		})
	}
	data := entitas.TipeBangunan{}
	err = c.Bind(&data)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}
	err = config.DB.Model(&data).Where("id = ?", id).Updates(data).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
	})
}

func GetAllTipeBangunan(c echo.Context) error {
	var result []entitas.TipeBangunan
	err := config.DB.Raw(`SELECT * FROM tipe_bangunans`).Scan(&result).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"data":    result,
	})
}

func DeleteTipeBangunan(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "bad request",
		})
	}
	data := entitas.TipeBangunan{}
	err = config.DB.Delete(&data, id).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
	})
}
