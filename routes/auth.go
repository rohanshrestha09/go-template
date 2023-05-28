package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/rohanshrestha09/go-template/controllers/auth"
	"github.com/rohanshrestha09/go-template/middleware"
)

func authRouter(router *gin.RouterGroup) {

	router.Use(middleware.Auth())
	{
		router.GET("/", auth.Get)

		router.PATCH("/", auth.Update)
	}

}
