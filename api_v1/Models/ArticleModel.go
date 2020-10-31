package Models

import "time"

type Article struct {
	ID 			uint		`gorm:"primaryKey;not null;"`
	Title		string		`gorm:"size:50;not null;"`
	Text 		string		`gorm:"not null;"`
	CreatedAt	time.Time	`gorm:"autoCreateTime;"`
}

func (b *Article) TableName() string {
	return "article"
}