package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rohanshrestha09/go-template/models"
	"github.com/rohanshrestha09/go-template/service/database"
)

func User() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {

		args := database.GetByIDArgs{
			Exclude: []string{"Email"},
		}

		data, err := database.GetByID[models.User](ctx.Param("id"), args)

		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}

		ctx.Set("user", data)

		ctx.Next()

	}
}
