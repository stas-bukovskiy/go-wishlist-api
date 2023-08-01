package entity

type User struct {
	Base
	Name     string `json:"name" binging:"required"`
	Email    string `json:"email" binging:"required" gorm:"uniqueIndex"`
	Password string `json:"-" gorm:"column:password_hash"`
}
