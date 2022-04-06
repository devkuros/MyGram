package models

import (
	"time"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Comment struct {
	gorm.Model
	UserID  uint   `json:"user_id"`
	User    *User  `json:"user" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	PhotoID uint   `json:"photo_id"`
	Photo   *Photo `json:"photo" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Message string `json:"message" gorm:"not null" valid:"required~Leave a comment messages"`
}

type CommentBody struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

type CommentPhotoBody struct {
	ID        uint      `json:"id"`
	Tittle    string    `json:"tittle"`
	Caption   string    `json:"caption"`
	PhotoUrl  string    `json:"photo_url"`
	UserID    uint      `json:"user_id"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (cmn *Comment) BeforeCreate(tx *gorm.DB) (err error) {
	if _, errCreate := govalidator.ValidateStruct(cmn); errCreate != nil {
		err = errCreate
		return
	}

	err = nil
	return
}
