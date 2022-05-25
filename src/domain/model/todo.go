package model

import (
)

type Todo struct {
	ID    uint    `json:"id" param:"id"`
	UserId uint   `json:"user_id"`
	Title  string `json:"title`
	Description string `json:"description"`
	Status string  `json:"status"`
	StartDate string  `json:"start_date"`
	ExpireDate string  `json:"expire_date"`
	Tags []Tag `json:"tags"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
