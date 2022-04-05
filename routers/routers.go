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
	userAuth := middlewares.NewUserAuthorization(db)

	userRouter := r.Group("users")
	{
		userRouter.POST("/register", ctrlUser.RegisterUsers)
		userRouter.POST("/login", ctrlUser.LoginUsers)
	}

	userMiddlewares := r.Group("users")
	{
		userMiddlewares.Use(middlewares.Authentication())
		userMiddlewares.PUT("/update", userAuth.UserAuthorizations(), ctrlUser.UpdateUsers)
	}

	return r
}
