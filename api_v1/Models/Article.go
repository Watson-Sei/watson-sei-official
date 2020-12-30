package Models

import (
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// GetAllArticles Fetch all article data
func (m Model) GetAllArticle() (*[]Article, error) {
	var article []Article
	tx := m.Db.Preload("Tags").Begin()
	err = tx.Find(&article).Error
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	tx.Commit()
	return &article, err
}

// CreateArticle ... Insert New data
func (m Model) CreateArticle(article *Article) error {
	tx := m.Db.Begin()
	err = tx.Create(article).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return err
}

// GetArticleByID ... Fetch only one article by Id
func (m Model) GetArticleById(id uint64) (*Article, error) {
	var article Article
	tx := m.Db.Preload("Tags").Begin()
	err = tx.Where("id = ?", id).First(&article).Error
	if err != nil {
		tx.Rollback()
		return &article, err
	}
	tx.Commit()
	return &article, err
}

// UpdateArticle ... Update article
func (m Model) UpdateArticle(article *Article) error {
	tx := m.Db.Begin()
	err := tx.Omit("Tags").Save(article).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return err
}

// DeleteArticle ... Delete article
func (m Model) DeleteArticle(article *Article, id uint64) error {
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
	tx := m.Db.Preload("Tags").Begin()
	if err = tx.Joins("inner join tag on article.id = tag.article_id").Where("tag.name = ?", tagParam).Preload("Tags").Find(&article).Error; err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return err
}