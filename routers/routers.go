package routers

import (
	"mygram/configs"
	"mygram/middlewares"
	"mygram/repositories"

	"github.com/gin-gonic/gin"
)

func StartServer() *gin.Engine {
	r := gin.Default()
	db := configs.Conns()

	ctrlUser := repositories.NewUserRepo(db)
	ctrlPhoto := repositories.NewphotoRepo(db)
	authorizations := middlewares.NewUserAuthorization(db)

	userRouter := r.Group("users")
	{
		userRouter.POST("/register", ctrlUser.RegisterUsers)
		userRouter.POST("/login", ctrlUser.LoginUsers)
	}

	userMiddlewares := r.Group("users")
	{
		userMiddlewares.Use(middlewares.Authentication())
		userMiddlewares.PUT("/update/:userId", authorizations.UserAuthorizations(), ctrlUser.UpdateUsers)
		userMiddlewares.DELETE("/deleted/:userId", authorizations.UserAuthorizations(), ctrlUser.DeleteUsers)
	}

	photoMiddlewares := r.Group("photos")
	{
		photoMiddlewares.Use(middlewares.Authentication())
		photoMiddlewares.GET("/", ctrlPhoto.GetPhotos)
		photoMiddlewares.POST("/upload", ctrlPhoto.UploadPhoto)
	}

	return r
}
