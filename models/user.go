package models

import (
	"backend-tugas-reactjs/utils/token"
	"html"
	"strings"
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type (
	User struct {
		ID        uint   `gorm:"primary_key" json:"id"`
		Username  string `gorm:"type:varchar(255); not null" json:"username"`
		Email     string `gorm:"type:varchar(255); not null" json:"email"`
		Password  string `gorm:"type:varchar(255); not null" json:"password"`
		Role      string `json:"role"`
		CreatedAt time.Time
		UpdatedAt time.Time
		Nilai     []Nilai `json:"-"`
	}
)

func (User) TableName() string {
	return "users"
}

func VerifyPassword(password, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func LoginCheck(username string, password string, db *gorm.DB) (string, User, error) {

	var err error

	u := User{}

	err = db.Model(User{}).Where("username = ?", username).Take(&u).Error

	if err != nil {
		return "", u, err
	}

	err = VerifyPassword(password, u.Password)

	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return "", u, err
	}

	token, err := token.GenerateToken(u.ID)

	if err != nil {
		return "", u, err
	}

	return token, u, nil

}

func (u *User) SaveUser(db *gorm.DB) (*User, error) {
	//turn password into hash
	hashedPassword, errPassword := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if errPassword != nil {
		return &User{}, errPassword
	}
	u.Password = string(hashedPassword)
	//remove spaces in username
	u.Username = html.EscapeString(strings.TrimSpace(u.Username))

	var err error = db.Create(&u).Error
	if err != nil {
		return &User{}, err
	}
	return u, nil
}
