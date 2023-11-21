package router

import (
	"finalProject3/controller"
	"finalProject3/middleware"

	"github.com/gin-gonic/gin"
)

func StartApp() *gin.Engine {
	r := gin.Default()

	userRouter := r.Group("/user")
	{
		userRouter.POST("/register", controller.NewUserHandlerImpl().UserRegister)

		userRouter.POST("/login", controller.NewUserHandlerImpl().UserLogin)

		userRouter.Use(middleware.Authentication())
		userRouter.PUT("/:userID", middleware.UserAuthorization(), controller.NewUserHandlerImpl().UserUpdate)

		userRouter.DELETE("/:userID", middleware.UserAuthorization(), controller.NewUserHandlerImpl().UserDelete)
	}

	categoryRouter := r.Group("/category")
	{
		categoryRouter.Use(middleware.Authentication())
		categoryRouter.POST("/create", controller.NewCategoryHandlerImpl().CategoryCreate)
		categoryRouter.GET("/get", controller.NewCategoryHandlerImpl().CategoryGet)
		categoryRouter.PATCH("/update/:categoryID", controller.NewCategoryHandlerImpl().CategoryUpdate)
		categoryRouter.DELETE("/delete/:categoryID", controller.NewCategoryHandlerImpl().CategoryDelete)
	}

	taskRouter := r.Group("/task")
	{
		taskRouter.Use(middleware.Authentication())
		taskRouter.POST("/create", controller.NewTaskHandlerImpl().TaskCreate)
		taskRouter.GET("/get", controller.NewTaskHandlerImpl().TaskGet)
		taskRouter.PUT("/update/:taskID", controller.NewTaskHandlerImpl().TaskUpdate)
		taskRouter.DELETE("/delete/:taskID", controller.NewTaskHandlerImpl().TaskDelete)
	}

	return r
}
