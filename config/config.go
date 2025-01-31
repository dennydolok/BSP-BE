package config

import (
	"dennydolok/BSP-BE/model/entitas"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

var DB *gorm.DB
var Secret = GetEnv("SECRET", "secret")

func InitConnection() {
	host := GetEnv("DB_HOST", "localhost")
	user := GetEnv("DB_USER", "root")
	port := GetEnv("DB_PORT", "3307")
	db := GetEnv("DB_NAME", "BSP_BE")
	//dsn1 := "root:@tcp(host.docker.internal:3307)/library_of_bears?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := fmt.Sprintf("%s:@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, host, port, db)
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}
	InitMigrate()
}

func InitMigrate() {
	err := DB.AutoMigrate(&entitas.User{})
	if err != nil {
		PrintLog(err)
		return
	}
	err = DB.AutoMigrate(&entitas.TipeBangunan{})
	if err != nil {
		PrintLog(err)
		return
	}
	err = DB.AutoMigrate(&entitas.Cabang{})
	if err != nil {
		PrintLog(err)
		return
	}
	err = DB.AutoMigrate(&entitas.Asuransi{})
	if err != nil {
		PrintLog(err)
		return
	}
}

func GetEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func PrintLog(text ...any) {
	fmt.Println(text...)
}
