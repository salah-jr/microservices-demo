package DB

import "gorm.io/gorm"

type Token struct {
	gorm.Model
	Token  string `json:"token"`
	UserId int    `json:"user_id"`
}

type RemoveToken struct {
	AuthToken string `json:"auth_token" binding:"required"`
	UserToken string `json:"user_token" binding:"required"`
	UserId    int    `json:"user_id" binding:"required"`
}

type Post struct {
	gorm.Model
	Title       string `json:"title"`
	Description string `json:"description"`
	UserId      int    `json:"user_id"`
}

type User struct {
	gorm.Model
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Token    string `json:"token"`
}
type Login struct {
	gorm.Model
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}
