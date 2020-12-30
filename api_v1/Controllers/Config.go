package Controllers

import (
	"github.com/Watson-Sei/watson-sei-official/api_v1/Models"
)

type Controller struct {
	Model Models.ModelInterface
}

type GetArticleController struct {
	Model Models.ModelInterfaceForGetArticle
}

type CreateArticleController struct {
	Model Models.ModelInterfaceForCreateArticle
}

type GetArticleByIdController struct {
	Model Models.ModelInterfaceForGetArticleById
}

type UpdateArticleController struct {
	Model Models.ModelInterfaceForUpdateArticle
}