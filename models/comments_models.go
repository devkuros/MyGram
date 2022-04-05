package models

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Comment struct {
	gorm.Model
	UserID  uint
	User    *User `json:"user" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	PhotoID uint
	Photo   *Photo `json:"photo" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Message string `json:"message" gorm:"not null" valid:"required~Leave a comment messages"`
}

func (cmn *Comment) BeforeCreate(tx *gorm.DB) (err error) {
	if _, errCreate := govalidator.ValidateStruct(cmn); errCreate != nil {
		err = errCreate
		return
	}

	err = nil
	return
}
