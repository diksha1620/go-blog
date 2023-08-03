package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	// ID       int `json:"id" gorm:"primary key"`
	RoleId   int
	Username string `json:"Username" validate:"required"`
	Email    string `json:"email" gorm:"unique"`
	Password string `json:"password"`
}

type Post struct {
	gorm.Model
	// ID        int `json:"id" gorm:"primary key"`
	UserID  int    `json:"user_id" gorm:"foreignKey:User" referance:"user"`
	Title   string `json:"title" gorm:"not null"`
	Content string `json:"content"`
	// CreatedAt time.Time `json:"created_at" gorm:"not null" sql:"DEFAULT:CURRENT_TIMESTAMP"`
	// UpdatedAt time.Time `json:"updated_at" gorm:"not null" sql:"DEFAULT:CURRENT_TIMESTAMP"`
}
type Comment struct {
	gorm.Model
	// ID        int `json:"id" gorm:"primary key"`
	UserID  int
	PostID  int
	Content string `json:"content"`
	// CreatedAt time.Time `json:"created_at"`
}
type Role struct {
	gorm.Model
	// Id   int    `json:"id" gorm:"primary key"`
	Role string `json:"role"`
}
