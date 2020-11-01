package Controllers

import (
	"fmt"
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

// CreateArticle ... Create all article
func CreateArticle(context *gin.Context) {
	var article Models.Article
	context.BindJSON(&article)
	err := Models.CreateArticle(&article)
	if err != nil {
		fmt.Println(err.Error())
		context.AbortWithStatus(http.StatusNotFound)
	} else {
		context.JSON(http.StatusOK, article)
	}
}