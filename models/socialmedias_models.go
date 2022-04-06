package models

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type SocialMedia struct {
	gorm.Model
	Nama           string `json:"nama" gorm:"not null" valid:"required~Input Name"`
	SocialMediaUrl string `json:"social_media_url" gorm:"not null" valid:"required~Input Social Media URL"`
	UserID         uint   `json:"user_id"`
	User           *User  `json:"user" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

func (sm *SocialMedia) BeforeCreate(tx *gorm.DB) (err error) {
	if _, errCreate := govalidator.ValidateStruct(sm); errCreate != nil {
		err = errCreate
		return
	}

	err = nil
	return
}
