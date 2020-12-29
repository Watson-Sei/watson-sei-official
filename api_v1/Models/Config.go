package Models

import "gorm.io/gorm"

var err error

type ModelInterface interface {
	GetAllArticle(article *[]Article) error
	CreateArticle(article *Article) error
	GetArticleByID(article *Article, id uint64) error
	UpdateArticle(article *Article) error
	DeleteArticle(article *Article, id uint64) error
	GetAllTag(tag *[]Tag) error
	GetArticleByTag(article *[]Article, tagParam string) error
	CreateUser(username string, password string) error
	PasswordChecker(username string, password string) error
	RefreshJWTToken(userId uint, token string, exp int64) (map[string]string, error)
}

type ModelInterfaceForGetArticle interface {
	GetAllArticle(article *[]Article) error
}

type Model struct {
	Db *gorm.DB
}