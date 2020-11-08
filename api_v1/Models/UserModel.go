package Models

type User struct {
	Username	string	`form:"username" json:"username" binding:"required" gorm:"unique;not null"`
	Password	string	`form:"password" json:"password" binding:"required" gorm:"not null"`
}

func (b *User) TableName() string {
	return "user"
}