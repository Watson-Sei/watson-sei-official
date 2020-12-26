package Models

import (
	"fmt"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// GetAllArticles Fetch all article data
func (m Model) GetAllArticle(article *[]Article) error {
	tx := m.Db.Preload("Tags").Begin()
	err = tx.Find(&article).Commit().Error
	return err
}

// CreateArticle ... Insert New data
func (m Model) CreateArticle(article *Article) error {
	err = m.Db.Create(article).Error
	return err
}

// GetArticleByID ... Fetch only one article by Id
func (m Model) GetArticleByID(article *Article, id string) error {
	err = m.Db.Where("id = ?", id).Preload("Tags").First(article).Error
	return err
}

// UpdateArticle ... Update article
func (m Model) UpdateArticle(article *Article) error {
	fmt.Println(article)
	m.Db.Omit("Tags").Save(article)
	return nil
}

// DeleteArticle ... Delete article
func (m Model) DeleteArticle(article *Article, id string) error {
	var tag Tag
	m.Db.Where("id = ?", id).Delete(article)
	m.Db.Where("article_id = ?", id).Delete(tag)
	return nil
}


// Tagの処理
func (m Model) GetAllTag(tag *[]Tag) error {
	if err = m.Db.Select("name").Group("name").Find(&tag).Error; err != nil {
		return err
	}
	return nil
}

// 特定タグの記事一覧
func (m Model) GetArticleByTag(article *[]Article, tagParam string) error {
	if err = m.Db.Joins("inner join tag on article.id = tag.article_id").Where("tag.name = ?", tagParam).Preload("Tags").Find(&article).Error; err != nil {
		return err
	}
	return nil
}