package models

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Photo struct {
	gorm.Model
	Tittle   string `json:"tittle" gorm:"not null" form:"tittle" valid:"required~Input Tittle"`
	Caption  string `json:"caption" form:"caption"`
	PhotoUrl string `json:"photo_url" gorm:"not null" valid:"required~Input Photo URL"`
	UserID   uint   `json:"user_id"`
	User     *User
}

type PhotoBody struct {
	ID       uint   `json:"id"`
	Tittle   string `json:"tittle"`
	Caption  string `json:"caption"`
	PhotoUrl string `json:"photo_url"`
	UserID   uint   `json:"user_id"`
}

func (ph *Photo) BeforeCreate(tx *gorm.DB) (err error) {
	if _, errCreate := govalidator.ValidateStruct(ph); errCreate != nil {
		err = errCreate
		return
	}

	err = nil
	return
}
