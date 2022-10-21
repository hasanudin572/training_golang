package models

import (
	"time"
)

type User struct {
	ID        uint      `json:"user_id" gorm:"primary_key"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	Age       int64     `json:"age"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}
type Photo struct {
	ID        uint      `json:"photo_id" gorm:"primary_key"`
	Title     string    `json:"title"`
	Caption   string    `json:"caption"`
	User_id   int64     `json:"user_id"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}
type Comment struct {
	ID        uint      `json:"comment_id" gorm:"primary_key"`
	User_id   int64     `json:"user_id"`
	Photo_id  int64     `json:"photo_id"`
	Message   string    `json:"message"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}
type SosialMedia struct {
	ID               uint   `json:"sosialmedia_id" gorm:"primary_key"`
	Name             string `json:"name"`
	Sosial_media_url string `json:"sosial_media_url"`
	User_id          int64  `json:"user_id"`
}
