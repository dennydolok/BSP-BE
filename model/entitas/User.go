package entitas

import "time"

type User struct {
	ID          int       `json:"id" gorm:"Column:id;primaryKey;autoIncrement"`
	Email       string    `json:"email" gorm:"Column:email;unique"`
	DateOfBirth time.Time `json:"date_of_birth" gorm:"Column:date_of_birth"`
	Nama        string    `json:"nama" gorm:"Column:nama"`
	Role        int       `json:"role" gorm:"Column:role"`
	Password    string    `json:"password" gorm:"Column:password"`
}

type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Register struct {
	Email          string `json:"email"`
	Nama           string `json:"nama"`
	VerifyEmail    string `json:"verify_email"`
	Password       string `json:"password"`
	VerifyPassword string `json:"verify_password"`
}
