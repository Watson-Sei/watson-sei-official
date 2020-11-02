package Models

type User struct {
	Username	string	`form:"username" binding:"required" gorm:"unique;not null"`
	Password	string	`form:"password" binding:"required"`
}

func (b *User) TableName() string {
	return "user"
}