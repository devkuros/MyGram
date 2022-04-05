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
	ctrlPhoto := repositories.NewPhotoRepo(db)
	authorizations := middlewares.NewAuthorization(db)

	userRouter := r.Group("users")
	{
		userRouter.POST("/register", ctrlUser.RegisterUsers)
		userRouter.POST("/login", ctrlUser.LoginUsers)
	}

	userMiddlewares := r.Group("users")
	{
		userMiddlewares.Use(middlewares.Authentication())
		userMiddlewares.PUT("/:userId", authorizations.UserAuthorizations(), ctrlUser.UpdateUsers)
		userMiddlewares.DELETE("/:userId", authorizations.UserAuthorizations(), ctrlUser.DeleteUsers)
	}

	photoMiddlewares := r.Group("photos")
	{
		photoMiddlewares.Use(middlewares.Authentication())
		photoMiddlewares.GET("/", ctrlPhoto.GetPhotos)
		photoMiddlewares.POST("/upload", ctrlPhoto.UploadPhoto)
		photoMiddlewares.PUT("/:photoId", authorizations.PhotoAuthorizations(), ctrlPhoto.UpdatePhotos)
		photoMiddlewares.DELETE("/:photoId", authorizations.PhotoAuthorizations(), ctrlPhoto.DeletePhotos)
	}

	return r
}
