package Controllers

import (
	"net/http"
	"strconv"

	"github.com/Watson-Sei/watson-sei-official/api_v1/Models"
	"github.com/gin-gonic/gin"
)

// GetArticle ... Get all article
func (c GetArticleController) GetArticle(context *gin.Context) {
	var articles []Models.Article
	err := c.Model.GetAllArticle(&articles)
	if err != nil {
		context.AbortWithStatus(http.StatusNotFound)
	} else {
		context.JSON(http.StatusOK, articles)
	}
}

// CreateArticle ... Create all article
func (c Controller) CreateArticle(context *gin.Context) {
	var article Models.Article
	context.BindJSON(&article)
	err := c.Model.CreateArticle(&article)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"err":err})
		return
	} else {
		context.JSON(http.StatusOK, article)
	}
}

// GetArticleByID ... Get the article by id
func (c Controller) GetArticleByID(context *gin.Context) {
	id := context.Params.ByName("id")
	covertedStrUint64, _ := strconv.ParseUint(id, 10, 64)
	var article Models.Article
	err := c.Model.GetArticleByID(&article, covertedStrUint64)
	if err != nil {
		context.AbortWithStatus(http.StatusNotFound)
	} else {
		context.JSON(http.StatusOK, article)
	}
}

// UpdateArticle ... Update the article information
func (c Controller) UpdateArticle(context *gin.Context) {
	var article Models.Article
	id := context.Params.ByName("id")
	covertedStrUint64, _ := strconv.ParseUint(id, 10, 64)
	err := c.Model.GetArticleByID(&article, covertedStrUint64)
	if err != nil {
		context.JSON(http.StatusNotFound, article)
		return
	}
	context.BindJSON(&article)
	err = c.Model.UpdateArticle(&article)
	if err != nil {
		context.AbortWithStatus(http.StatusNotFound)
	} else {
		context.JSON(http.StatusOK, article)
	}
}

// DeleteArticle ... Delete the article
func (c Controller) DeleteArticle(context *gin.Context) {
	var article Models.Article
	id := context.Params.ByName("id")
	covertedStrUint64, _ := strconv.ParseUint(id, 10, 64)
	err := c.Model.DeleteArticle(&article, covertedStrUint64)
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
func (c Controller) GetAllTag(context *gin.Context) {
	var tags []Models.Tag
	err := c.Model.GetAllTag(&tags)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"err":err})
		return
	}
	context.JSON(http.StatusOK, tags)
}

func (c Controller) GetArticleByTag(context *gin.Context) {
	tagParam := context.Params.ByName("tag")
	var articles []Models.Article
	err := c.Model.GetArticleByTag(&articles, tagParam)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"err":err})
		return
	}
	context.JSON(http.StatusOK, articles)
}
