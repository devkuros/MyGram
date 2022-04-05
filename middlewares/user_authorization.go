package middlewares

import (
	"mygram/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"gorm.io/gorm"
)

type userAuthorization struct {
	DB *gorm.DB
}

func NewUserAuthorization(db *gorm.DB) *userAuthorization {
	return &userAuthorization{
		DB: db,
	}
}

func (ua *userAuthorization) UserAuthorizations() gin.HandlerFunc {
	return func(c *gin.Context) {
		getId, err := strconv.Atoi(c.Param("userId"))
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error":   "Bad Request",
				"message": "invalid parameter",
			})
			return
		}

		userData := c.MustGet("userData").(jwt.MapClaims)
		userID := uint(userData["id"].(float64))

		User := models.User{}

		err = ua.DB.Select("id").First(&User, uint(getId)).Error
		if err != nil {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"error":   "Data not found",
				"message": "data not exist",
			})
			return
		}

		if User.ID != userID {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":   "Unauthorized",
				"message": "please try again!",
			})
			return
		}
		c.Next()
	}
}
