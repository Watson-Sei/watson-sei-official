package Controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Home(context *gin.Context)  {
	context.JSON(http.StatusOK, gin.H{
		"message":  "main json",
	})
}
