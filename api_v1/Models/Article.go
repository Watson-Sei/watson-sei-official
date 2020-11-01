package Models

import "github.com/Watson-Sei/watson-sei-official/api_v1/Config"
import _ "github.com/jinzhu/gorm/dialects/mysql"

// GetAllArticles Fetch all article data
func GetAllArticle(article *[]Article) (err error) {
	if err = Config.DB.Find(article).Error; err != nil {
		return err
	}
	return nil
}

// CreateArticle ... Insert New data
func CreateArticle(article *Article) (err error) {
	if err = Config.DB.Create(article).Error; err != nil {
		return err
	}
	return nil
}

// GetArticleByID ... Fetch only one article by Id
func GetArticleByID(article *Article, id string) (err error) {
	if err = Config.DB.Where("id = ?", id).First(article).Error; err != nil {
		return err
	}
	return nil
}

