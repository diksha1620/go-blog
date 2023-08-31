package models

import (
	"html"
	"strings"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	// ID       int `json:"id" gorm:"primary key"`
	FirstName string `json:"firstname" binding:"required"`
	LastName  string `json:"lastname" binding:"required"`
	Username  string `json:"Username" binding:"required"`
	RoleId    int
	Email     string `json:"email" binding:"required,email"`
	Password  string `json:"password"`
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

func (u *User) SaveUser() (*User, error) {

	err := DB.Create(&u).Error
	if err != nil {
		return &User{}, err
	}
	return u, nil
}

func (u *User) BeforeSave() error {

	//turn password into hash
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)

	//remove spaces in username
	u.Username = html.EscapeString(strings.TrimSpace(u.Username))

	return nil

}
