package Models

import "time"

type Article struct {
	ID 			uint64		`gorm:"primaryKey;AUTO_INCREMENT;not null;" json:"id"`
	Title		string		`gorm:"size:50;not null;" json:"title"`
	Overview	string		`gorm:"size:90;" json:"overview"`
	Text 		string		`gorm:"type: Text;" json:"text"`
	Tags		[]Tag		`gorm:"foreignKey:ArticleID" json:"tags"`
	CreatedAt	time.Time	`gorm:"autoCreateTime;" json:"created_at"`
}

type Tag struct {
	Name 		string
	ArticleID	uint64
}

func (b *Article) TableName() string {
	return "article"
}

func (b *Tag) TableName() string {
	return "tag"
}