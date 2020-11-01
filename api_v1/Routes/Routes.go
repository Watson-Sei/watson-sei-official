package Routes

import (
	"github.com/Watson-Sei/watson-sei-official/api_v1/Controllers"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/", Controllers.GetArticle)
	router.POST("/create", Controllers.CreateArticle)
	router.GET("/:id", Controllers.GetArticleByID)
	router.PUT("/:id", Controllers.UpdateArticle)
	router.DELETE("/:id", Controllers.DeleteArticle)
	return router
}