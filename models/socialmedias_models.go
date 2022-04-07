package models

import (
	"time"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type SocialMedia struct {
	gorm.Model
	Nama           string `json:"name" gorm:"not null" valid:"required~Input Name"`
	SocialMediaUrl string `json:"social_media_url" gorm:"not null" valid:"required~Input Social Media URL"`
	UserID         uint   `json:"user_id"`
	User           *User  `json:"user" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

type GetSocialMediaBody struct {
	ID             uint      `json:"id"`
	Nama           string    `json:"name"`
	SocialMediaUrl string    `json:"social_media_url"`
	UserID         uint      `json:"user_id"`
	UpdatedAt      time.Time `json:"updated_at"`
}

type UserSocialMediaBody struct {
	StatsSocialMedia struct {
		ID             uint      `json:"id"`
		Nama           string    `json:"name"`
		SocialMediaUrl string    `json:"social_media_url"`
		UserID         uint      `json:"user_id"`
		CreatedAt      time.Time `json:"created_at"`
		UpdatedAt      time.Time `json:"updated_at"`
		StatsUsers     struct {
			ID            uint   `json:"id"`
			Username      string `json:"username"`
			ProfileImgUrl string `json:"profile_image_url"`
		} `json:"user"`
	} `json:"social_medias"`
}

func (sm *SocialMedia) BeforeCreate(tx *gorm.DB) (err error) {
	if _, errCreate := govalidator.ValidateStruct(sm); errCreate != nil {
		err = errCreate
		return
	}

	err = nil
	return
}
