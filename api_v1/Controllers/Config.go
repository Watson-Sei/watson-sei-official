package Controllers

import (
	"github.com/Watson-Sei/watson-sei-official/api_v1/Models"
)

type Controller struct {
	GetArticleModel Models.ModelInterfaceForGetArticle
	Model Models.ModelInterface
}

// いっぱいあって、どれをモックにしたら良いかわからなくなっちゃう
controller := Controller{
	GetArticleModel: mock
	Model: 本物
}

type GetArticleController struct {
	Model Models.ModelInterfaceForGetArticle
}
