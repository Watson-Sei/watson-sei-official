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

// GetArticleByID ... Get the article by id
func GetArticleByID(context *gin.Context) {
	id := context.Params.ByName("id")
	var article Models.Article
	err := Models.GetArticleByID(&article, id)
	if err != nil {
		context.AbortWithStatus(http.StatusNotFound)
	} else {
		context.JSON(http.StatusOK, article)
	}
}

// UpdateArticle ... Update the article information
func UpdateArticle(context *gin.Context) {
	var article Models.Article
	id := context.Params.ByName("id")
	err := Models.GetArticleByID(&article, id)
	if err != nil {
		context.JSON(http.StatusNotFound, article)
		return
	}
	context.BindJSON(&article)
	err = Models.UpdateArticle(&article, id)
	if err != nil {
		context.AbortWithStatus(http.StatusNotFound)
	} else {
		context.JSON(http.StatusOK, article)
	}
}

// DeleteArticle ... Delete the article
func DeleteArticle(context *gin.Context) {
	var article Models.Article
	id := context.Params.ByName("id")
	err := Models.DeleteArticle(&article, id)
	if err != nil {
		context.AbortWithStatus(http.StatusNotFound)
	} else {
		context.JSON(http.StatusOK, gin.H{"id" + id: "is deleted"})
	}
}