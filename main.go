package main

import (
	"dennydolok/BSP-BE/config"
	Controller "dennydolok/BSP-BE/controller"
	"net/http"

	echojwt "github.com/labstack/echo-jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	config.InitConnection()
	e := New()
	e.Logger.Fatal(e.Start("localhost:81"))

}

func New() *echo.Echo {
	e := echo.New()

	e.GET("/healthcheck", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"status": "ok",
		})
	})

	cabang := e.Group("/cabang")
	//Cabang
	cabang.Use(echojwt.WithConfig(echojwt.Config{
		SigningKey: []byte(config.Secret),
	}))
	cabang.GET("", Controller.GetAllCabang)
	cabang.GET(":id", Controller.GetCabang)
	cabang.POST("", Controller.AddCabang)
	cabang.PUT(":id", Controller.UpdateCabang)
	cabang.DELETE(":id", Controller.DeleteCabang)

	//Tipe Bangunan
	e.GET("/tipebangunan", Controller.GetAllTipeBangunan)
	e.GET("/tipebangunan/:id", Controller.GetTipeBangunan)
	e.POST("/tipebangunan", Controller.AddTipeBangunan)
	e.PUT("/tipebangunan/:id", Controller.UpdateTipeBangunan)
	e.DELETE("/tipebangunan/:id", Controller.DeleteTipeBangunan)

	//Asuransi
	e.GET("/asuransi", Controller.GetAsuransi)
	e.GET("/asuransi/:id", Controller.GetAsuransiByID)
	e.POST("/asuransi", Controller.AddAsuransi, echojwt.JWT([]byte(config.Secret)))
	e.POST("/asuransi/approve/:id", Controller.Approve)
	e.POST("/asuransi/reject/:id", Controller.Reject)
	e.POST("/asuransi/premi", Controller.KalkulasiPremi)

	//User
	e.POST("/login", Controller.Login, middleware.Logger())
	e.POST("/register", Controller.Register)
	e.PUT("/user/:id", Controller.Update)
	return e
}
