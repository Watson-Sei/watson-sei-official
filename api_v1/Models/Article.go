package Models

import (
	"fmt"

	"github.com/Watson-Sei/watson-sei-official/api_v1/Config"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// GetAllArticles Fetch all article data
func GetAllArticle(article *[]Article) (err error) {
	db := Config.DBConnect()
	defer db.Close()
	if err = db.Preload("Tags").Find(&article).Error; err != nil {
		return err
	}
	return nil
}

// CreateArticle ... Insert New data
func CreateArticle(article *Article) (err error) {
	db := Config.DBConnect()
	defer db.Close()
	if err = db.Create(article).Error; err != nil {
		return err
	}
	return nil
}

// GetArticleByID ... Fetch only one article by Id
func GetArticleByID(article *Article, id string) (err error) {
	db := Config.DBConnect()
	defer db.Close()
	if err = db.Where("id = ?", id).Preload("Tags").First(article).Error; err != nil {
		return err
	}
	return nil
}

// UpdateArticle ... Update article
func UpdateArticle(article *Article) (err error) {
	db := Config.DBConnect()
	defer db.Close()
	fmt.Println(article)
	db.Omit("Tags").Save(article)
	return nil
}

// DeleteArticle ... Delete article
func DeleteArticle(article *Article, id string) (err error) {
	var tag Tag
	db := Config.DBConnect()
	defer db.Close()
	db.Where("id = ?", id).Delete(article)
	db.Where("article_id = ?", id).Delete(tag)
	return nil
}


// Tagの処理
func GetAllTag(tag *[]Tag) (err error) {
	db := Config.DBConnect()
	defer db.Close()
	if err = db.Select("name").Group("name").Find(&tag).Error; err != nil {
		return err
	}
	return nil
}

// 特定タグの記事一覧
func GetArticleByTag(article *[]Article, tagParam string) (err error) {
	db := Config.DBConnect()
	defer db.Close()
	if err = db.Joins("inner join tag on article.id = tag.article_id").Where("tag.name = ?", tagParam).Preload("Tags").Find(&article).Error; err != nil {
		return err
	}
	return nil
}