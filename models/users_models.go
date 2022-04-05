package models

import (
	"mygram/handlers"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string  `json:"username" gorm:"not null;unique" form:"username" valid:"required~Input Username"`
	Email    string  `json:"email" gorm:"not null;unique" form:"email" valid:"required~Email required,email~Invalid email format"`
	Password string  `json:"password" gorm:"not null" form:"password" valid:"required~Input Password,minstringlength(6)~Password has to have a minimum length of 6 characters"`
	Age      int     `json:"age" gorm:"not null" form:"age" valid:"required~Input Age,range(8|130)~Min Age 8,numeric"`
	Photo    []Photo `json:"photo" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

type UserBody struct {
	Username string `json:"username"`
	Email    string `json:"email"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	if _, errCreate := govalidator.ValidateStruct(u); errCreate != nil {
		err = errCreate
		return
	}

	u.Password = handlers.HashPassword(u.Password)
	err = nil
	return
}
