package repositories

import (
	"mygram/handlers"
	"mygram/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"gorm.io/gorm"
)

var (
	socialMediatApp = "application/json"
)

type socialMediaRepo struct {
	DB *gorm.DB
}

func NewSocialMediaRepo(db *gorm.DB) *socialMediaRepo {
	return &socialMediaRepo{
		DB: db,
	}
}

func (som *socialMediaRepo) GetSocialMedias(c *gin.Context) {
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := handlers.GetContentType(c)

	Social := models.SocialMedia{}
	User := models.User{}
	Photo := models.Photo{}
	SocialBody := models.UserSocialMediaBody{}

	userID := uint(userData["id"].(float64))
	if contentType == socialMediatApp {
		c.ShouldBindJSON(&Social)
	} else {
		c.ShouldBind(&Social)
	}

	Social.UserID = userID

	if err := som.DB.Where("user_id = ?", userID).First(&Social).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	if err := som.DB.Where("id = ?", Social.UserID).First(&User).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	if err := som.DB.Select("photo_url").Where("user_id = ?", userID).First(&Photo).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error":   "Data not found",
			"message": err.Error(),
		})
		return
	}

	SocialBody.StatsSocialMedia.ID = Social.ID
	SocialBody.StatsSocialMedia.Nama = Social.Nama
	SocialBody.StatsSocialMedia.SocialMediaUrl = Social.SocialMediaUrl
	SocialBody.StatsSocialMedia.UserID = Social.UserID
	SocialBody.StatsSocialMedia.UpdatedAt = Social.UpdatedAt
	SocialBody.StatsSocialMedia.CreatedAt = Social.CreatedAt

	SocialBody.StatsSocialMedia.StatsUsers.ID = User.ID
	SocialBody.StatsSocialMedia.StatsUsers.Username = User.Username
	SocialBody.StatsSocialMedia.StatsUsers.ProfileImgUrl = Photo.PhotoUrl

	c.JSON(http.StatusOK, SocialBody)
}

func (som *socialMediaRepo) CreateSocialMedias(c *gin.Context) {
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := handlers.GetContentType(c)
	Social := models.SocialMedia{}

	userID := uint(userData["id"].(float64))

	if contentType == socialMediatApp {
		c.ShouldBindJSON(&Social)
	} else {
		c.ShouldBind(&Social)
	}

	Social.UserID = userID

	if err := som.DB.Debug().Create(&Social).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(201, gin.H{
		"id":               Social.ID,
		"name":             Social.Nama,
		"social_media_url": Social.SocialMediaUrl,
		"created_at":       Social.CreatedAt,
	})
}
