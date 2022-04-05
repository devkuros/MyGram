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
	userApp = "application/json"
)

type userRepo struct {
	DB *gorm.DB
}

func NewUserRepo(db *gorm.DB) *userRepo {
	return &userRepo{
		DB: db,
	}
}

func (usr *userRepo) RegisterUsers(c *gin.Context) {
	contentType := handlers.GetContentType(c)
	_, _ = usr.DB, contentType
	User := models.User{}

	if contentType == userApp {
		c.ShouldBindJSON(&User)
	} else {
		c.ShouldBind(&User)
	}

	if err := usr.DB.Debug().Create(&User).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"id":       User.ID,
		"email":    User.Email,
		"username": User.Username,
		"age":      User.Age,
	})
}

func (usr *userRepo) LoginUsers(c *gin.Context) {
	contentType := handlers.GetContentType(c)
	_, _ = usr.DB, contentType

	User := models.User{}
	password := ""

	if contentType == userApp {
		c.ShouldBindJSON(&User)
	} else {
		c.ShouldBind(&User)
	}

	password = User.Password

	if err := usr.DB.Debug().Where("email = ?", User.Email).Take(&User).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error":   "Unauthorized",
			"message": "invalid email/password",
		})
		return
	}

	comparePass := handlers.ComparePassword([]byte(User.Password), []byte(password))

	if !comparePass {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error":   "Unauthorized",
			"message": "invalid email/password",
		})
		return
	}

	token := handlers.GenerateToken(User.ID, User.Email)
	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}

func (usr *userRepo) UpdateUsers(c *gin.Context) {
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := handlers.GetContentType(c)

	User := models.User{}

	getId, _ := strconv.Atoi(c.Param("userId"))
	userID := uint(userData["id"].(float64))

	if contentType == userApp {
		c.ShouldBindJSON(&User)
	} else {
		c.ShouldBind(&User)
	}

	_, User.ID = userID, uint(getId)

	if err := usr.DB.Model(&User).Where("id = ?", getId).Updates(models.User{Email: User.Email, Username: User.Username}).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":        User.ID,
		"email":     User.Email,
		"username":  User.Username,
		"age":       User.Age,
		"update_at": User.UpdatedAt,
	})

}
