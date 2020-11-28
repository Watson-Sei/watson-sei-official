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

// CreateArticle ... Create all article
func CreateArticle(context *gin.Context) {
	var article Models.Article
	context.BindJSON(&article)
	err := Models.CreateArticle(&article)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"err":err})
		return
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
	err = Models.UpdateArticle(&article)
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

// 画像アップロード
func UploadImage(context *gin.Context)  {
	form, _ := context.MultipartForm()
	file := form.File["file"][0]
	uuid := context.PostForm("uuid")
	err := context.SaveUploadedFile(file, "assets/" + uuid + ".png")
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"err":err.Error()})
		context.Abort()
	}
	context.JSON(http.StatusOK, gin.H{"url": "assets/" + uuid + ".png"})
}

// タグ一覧
func GetAllTag(context *gin.Context) {
	var tags []Models.Tag
	err := Models.GetAllTag(&tags)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"err":err})
		return
	}
	context.JSON(http.StatusOK, tags)
}

func GetArticleByTag(context *gin.Context) {
	tagParam := context.Params.ByName("tag")
	var articles []Models.Article
	err := Models.GetArticleByTag(&articles, tagParam)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"err":err})
		return
	}
	context.JSON(http.StatusOK, articles)
}