package Controller

import (
	helper "dennydolok/BSP-BE/Helper"
	"dennydolok/BSP-BE/config"
	"dennydolok/BSP-BE/model/entitas"
	"errors"
	"fmt"
	"net/http"
	"regexp"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func AddAsuransi(c echo.Context) error {
	asuransi := entitas.Asuransi{}
	err := c.Bind(&asuransi)
	if err != nil {
		config.PrintLog(err.Error())
		return c.JSON(400, map[string]interface{}{
			"message": "Invalid request",
		})
	}
	userID := helper.GetClaimsID(c.Request().Header.Get("Authorization"))
	user := entitas.User{}

	err = config.DB.Where("id = ?", userID).First(&user).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "User not found",
		})
	}
	Age := time.Now().Year() - user.DateOfBirth.Year()
	asuransi.Usia = Age
	asuransi.UserID = userID
	asuransi.Approve = 0
	asuransi.NomorPolis = "Belum Terbit"
	asuransi.Status = "Belum dibayar"
	asuransi.Total = asuransi.Premi + 10.000
	err = config.DB.Create(&asuransi).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}
	return c.JSON(http.StatusCreated, map[string]interface{}{
		"message": "Asuransi created",
	})
}

func KalkulasiPremi(c echo.Context) error {
	asuransi := entitas.Asuransi{}
	err := c.Bind(&asuransi)
	if err != nil {
		config.PrintLog(err.Error())
		return c.JSON(400, map[string]interface{}{
			"message": "Invalid request",
		})
	}
	bangunan := entitas.TipeBangunan{}
	err = config.DB.Where("id = ?", asuransi.TipaBangunanID).First(&bangunan).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Tipe bangunan not found",
		})
	}
	config.PrintLog(bangunan.Tarif, asuransi.JangaWaktu, asuransi.HargaBangunan)
	premi := asuransi.HargaBangunan * bangunan.Tarif / 1000 * float64(asuransi.JangaWaktu)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"premi": premi,
	})
}

func Approve(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "bad request",
		})
	}
	nomorPolis, err := getNextNomorPolis(config.DB)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}
	err = config.DB.Model(&entitas.Asuransi{}).Where("id = ?", id).Updates(&entitas.Asuransi{Approve: 1, NomorPolis: nomorPolis, Status: "Sudah dibayar"}).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Asuransi approved",
	})
}
func Reject(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "bad request",
		})
	}
	err = config.DB.Model(&entitas.Asuransi{}).Where("id = ?", id).Updates(&entitas.Asuransi{Approve: 0, Status: "Belum dibayar"}).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Asuransi approved",
	})
}

func GetAsuransi(c echo.Context) error {
	var asuransi []entitas.Asuransi
	authCheck := helper.GetClaimsRole(c.Request().Header.Get("Authorization"))
	if authCheck == 0 {
		return c.JSON(http.StatusUnauthorized, map[string]interface{}{
			"message": "unauthorized",
		})
	}
	if authCheck == 1 {
		config.DB.Preload(clause.Associations).Find(&asuransi)
	} else {
		userID := helper.GetClaimsID(c.Request().Header.Get("Authorization"))
		config.DB.Preload(clause.Associations).Where("user_id = ?", userID).Find(&asuransi)
	}
	return c.JSON(http.StatusOK, asuransi)
}

func GetAsuransiByID(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "bad request",
		})
	}
	asuransi := entitas.Asuransi{}
	err = config.DB.Where("id = ?", id).Preload(clause.Associations).First(&asuransi).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, asuransi)
}

func getNextNomorPolis(db *gorm.DB) (string, error) {
	var lastAsuransi entitas.Asuransi
	result := db.Where("nomor_polis LIKE ?", "K.001.%").Order("nomor_polis DESC").First(&lastAsuransi)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return "K.001.00001", nil
	} else if result.Error != nil {
		return "", result.Error
	}

	re := regexp.MustCompile(`K\.001\.(\d{5})`)
	matches := re.FindStringSubmatch(lastAsuransi.NomorPolis)
	if len(matches) < 2 {
		return "", fmt.Errorf("invalid nomor_polis format: %s", lastAsuransi.NomorPolis)
	}
	num, err := strconv.Atoi(matches[1])
	if err != nil {
		return "", err
	}
	num++
	newNomorPolis := fmt.Sprintf("K.001.%05d", num)
	return newNomorPolis, nil
}
