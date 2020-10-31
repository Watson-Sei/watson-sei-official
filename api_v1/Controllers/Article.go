package Controllers

import (
	"net/http"

	"github.com/Watson-Sei/watson-sei-official/api_v1/Models"
	"github.com/gin-gonic/gin"
)

// GetArticle ... Get all article
func GetArticle(context *gin.Context) {
	var articles []Models.Article
	err := Models.GetAllArticle(&articles)
	if err != nil {
		context.AbortWithStatus(http.StatusNotFound)
	} else {
		context.JSON(http.StatusOK, articles)
	}
}