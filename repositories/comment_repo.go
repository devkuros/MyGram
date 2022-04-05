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
	commentApp = "application/json"
)

type commentRepo struct {
	DB *gorm.DB
}

func NewCommentRepo(db *gorm.DB) *commentRepo {
	return &commentRepo{
		DB: db,
	}
}

func (cmt *commentRepo) CreateComment(c *gin.Context) {
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := handlers.GetContentType(c)
	Comment := models.Comment{}

	userID := uint(userData["id"].(float64))
	if contentType == commentApp {
		c.ShouldBindJSON(&Comment)
	} else {
		c.ShouldBind(&Comment)
	}

	Comment.UserID = userID
	Comment.PhotoID = userID

	// get id comment from tabel comment > samain user_id di tabel comment dan user_id diJWT > kalau sama proses.

	if err := cmt.DB.Debug().Create(&Comment).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"id":         Comment.ID,
		"message":    Comment.Message,
		"photo_id":   Comment.PhotoID,
		"user_id":    Comment.UserID,
		"created_at": Comment.CreatedAt,
	})
}
