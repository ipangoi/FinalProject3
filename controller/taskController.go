package controller

import "github.com/gin-gonic/gin"

type TaskHandler interface {
	TaskCreate(*gin.Context)
	TaskGet(*gin.Context)
	TaskUpdate(*gin.Context)
	TaskDelete(*gin.Context)
}
