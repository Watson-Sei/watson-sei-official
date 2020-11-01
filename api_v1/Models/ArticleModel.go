package Models

import "time"

type Article struct {
	ID 			uint		`gorm:"primaryKey;not null;" json:"id"`
	Title		string		`gorm:"size:50;not null;" json:"title"`
	Text 		string		`gorm:"not null;" json:"text"`
	CreatedAt	time.Time	`gorm:"autoCreateTime;" json:"created_at"`
}

func (b *Article) TableName() string {
	return "article"
}