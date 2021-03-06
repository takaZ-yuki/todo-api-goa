package model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID    int    `json:"id" param:"id"`
	Mail  string `json:"mail"`
	Nickname  string `json:"nickname"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}
