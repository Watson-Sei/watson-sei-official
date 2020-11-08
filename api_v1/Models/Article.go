package Models

import (
	"fmt"

	"github.com/Watson-Sei/watson-sei-official/api_v1/Config"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// GetAllArticles Fetch all article data
func GetAllArticle(article *[]Article) (err error) {
	db := Config.DbConnect()
	defer db.Close()
	if err = db.Find(article).Error; err != nil {
		return err
	}
	return nil
}

// CreateArticle ... Insert New data
func CreateArticle(article *Article) (err error) {
	db := Config.DbConnect()
	defer db.Close()
	if err = db.Create(article).Error; err != nil {
		return err
	}
	return nil
}

// GetArticleByID ... Fetch only one article by Id
func GetArticleByID(article *Article, id string) (err error) {
	db := Config.DbConnect()
	defer db.Close()
	if err = db.Where("id = ?", id).First(article).Error; err != nil {
		return err
	}
	return nil
}

// UpdateArticle ... Update article
func UpdateArticle(article *Article, id string) (err error) {
	db := Config.DbConnect()
	defer db.Close()
	fmt.Println(article)
	db.Save(article)
	return nil
}

// DeleteArticle ... Delete article
func DeleteArticle(article *Article, id string) (err error) {
	db := Config.DbConnect()
	defer db.Close()
	db.Where("id = ?", id).Delete(article)
	return nil
}