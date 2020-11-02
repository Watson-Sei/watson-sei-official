package Routes

import (
	"github.com/Watson-Sei/watson-sei-official/api_v1/Controllers"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	router.LoadHTMLGlob("Templates/*.html")
	v1 := router.Group("/v1")
	{
		v1.GET("/", Controllers.GetArticle)
		v1.POST("/create", Controllers.CreateArticle)
		v1.GET("/:id", Controllers.GetArticleByID)
		v1.PUT("/:id", Controllers.UpdateArticle)
		v1.DELETE("/:id", Controllers.DeleteArticle)
	}
	admin := router.Group("/admin")
	{
		admin.GET("/signup", Controllers.SignupGet)
		admin.POST("/signup", Controllers.SignupPost)
	}

	return router
}