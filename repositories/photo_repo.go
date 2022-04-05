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
	photoApp = "application/json"
)

type photoRepo struct {
	DB *gorm.DB
}

func NewphotoRepo(db *gorm.DB) *photoRepo {
	return &photoRepo{
		DB: db,
	}
}

func (pht *photoRepo) UploadPhoto(c *gin.Context) {
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := handlers.GetContentType(c)

	Photo := models.Photo{}
	userID := uint(userData["id"].(float64))

	if contentType == photoApp {
		c.ShouldBindJSON(&Photo)
	} else {
		c.ShouldBind(&Photo)
	}

	Photo.UserID = userID

	if err := pht.DB.Debug().Create(&Photo).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"id":         Photo.ID,
		"tittle":     Photo.Tittle,
		"caption":    Photo.Caption,
		"photo_url":  Photo.PhotoUrl,
		"user_id":    Photo.UserID,
		"created_at": Photo.CreatedAt,
	})
}

func (pht *photoRepo) GetPhotos(c *gin.Context) {
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := handlers.GetContentType(c)

	Photo := models.Photo{}
	User := models.User{}
	UserBody := models.UserBody{}

	userID := uint(userData["id"].(float64))

	if contentType == photoApp {
		c.ShouldBindJSON(&Photo)
	} else {
		c.ShouldBind(&Photo)
	}

	Photo.UserID = userID

	if err := pht.DB.Debug().Find(&Photo).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	if err := pht.DB.Debug().Select("Email", "Username").Find(&User).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	UserBody.Email = User.Email
	UserBody.Username = User.Username

	c.JSON(http.StatusCreated, gin.H{
		"id":         Photo.ID,
		"tittle":     Photo.Tittle,
		"caption":    Photo.Caption,
		"photo_url":  Photo.PhotoUrl,
		"user_id":    Photo.UserID,
		"created_at": Photo.CreatedAt,
		"updated_at": Photo.UpdatedAt,
		"user":       UserBody,
	})
}
