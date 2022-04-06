package repositories

import (
	"mygram/handlers"
	"mygram/models"
	"net/http"
	"strconv"

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

func (cmt *commentRepo) UpdateComments(c *gin.Context) {
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := handlers.GetContentType(c)

	Comment := models.Comment{}
	Photo := models.Photo{}
	CommentPhotoBody := models.CommentPhotoBody{}

	commentId, _ := strconv.Atoi(c.Param("commentId"))
	userID := uint(userData["id"].(float64))

	if contentType == commentApp {
		c.ShouldBindJSON(&Comment)
	} else {
		c.ShouldBind(&Comment)
	}

	Comment.UserID = userID
	Comment.ID = uint(commentId)

	if err := cmt.DB.Model(&Comment).Where("id = ?", commentId).Updates(models.Comment{Message: Comment.Message}).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     "Bad Request",
			"message": err.Error(),
		})
		return
	}

	if err := cmt.DB.Debug().Find(&Comment).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	getPhotoId := Comment.PhotoID

	if err := cmt.DB.Debug().Where("id = ?", getPhotoId).First(&Photo).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	CommentPhotoBody.ID = getPhotoId
	CommentPhotoBody.Tittle = Photo.Tittle
	CommentPhotoBody.Caption = Photo.Caption
	CommentPhotoBody.PhotoUrl = Photo.PhotoUrl
	CommentPhotoBody.UserID = Photo.UserID
	CommentPhotoBody.UpdatedAt = Photo.UpdatedAt

	c.JSON(http.StatusOK, gin.H{
		"id":         CommentPhotoBody.ID,
		"tittle":     CommentPhotoBody.Tittle,
		"caption":    CommentPhotoBody.Caption,
		"photo_url":  CommentPhotoBody.PhotoUrl,
		"user_id":    CommentPhotoBody.UserID,
		"updated_at": CommentPhotoBody.UpdatedAt,
	})
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

func (cmt *commentRepo) GetComments(c *gin.Context) {
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := handlers.GetContentType(c)

	Comment := models.Comment{}
	User := models.User{}
	Photo := models.Photo{}
	CommentBody := models.CommentBody{}
	PhotoBody := models.PhotoBody{}

	userID := uint(userData["id"].(float64))

	if contentType == commentApp {
		c.ShouldBindJSON(&Comment)
	} else {
		c.ShouldBind(&Comment)
	}

	Comment.UserID = userID

	if err := cmt.DB.Debug().Find(&Comment).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	if err := cmt.DB.Debug().Find(&User).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	if err := cmt.DB.Debug().Find(&Photo).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	CommentBody.ID = User.ID
	CommentBody.Email = User.Email
	CommentBody.Username = User.Username

	PhotoBody.ID = Photo.ID
	PhotoBody.Tittle = Photo.Tittle
	PhotoBody.Caption = Photo.Caption
	PhotoBody.PhotoUrl = Photo.PhotoUrl
	PhotoBody.UserID = Photo.UserID

	c.JSON(http.StatusCreated, gin.H{
		"id":         Comment.ID,
		"message":    Comment.Message,
		"photo_id":   Comment.PhotoID,
		"user_id":    Comment.UserID,
		"created_at": Comment.CreatedAt,
		"updated_at": Comment.UpdatedAt,
		"user":       CommentBody,
		"photo":      PhotoBody,
	})
}

func (cmt *commentRepo) DeleteComments(c *gin.Context) {
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := handlers.GetContentType(c)
	Comment := models.Comment{}

	commentId, _ := strconv.Atoi(c.Param("commentId"))
	userID := uint(userData["id"].(float64))

	if contentType == commentApp {
		c.ShouldBindJSON(&Comment)
	} else {
		c.ShouldBind(&Comment)
	}

	Comment.UserID = userID
	Comment.ID = uint(commentId)

	if err := cmt.DB.Model(&Comment).Where("id = ?", commentId).Delete(&Comment).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "your comment has been successfully deleted",
	})
}
