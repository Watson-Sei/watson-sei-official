package Controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Main(context *gin.Context)  {
	context.JSON(http.StatusOK, gin.H{
		"message":  "main json",
		"exp": context.MustGet("exp"),
	})
}
