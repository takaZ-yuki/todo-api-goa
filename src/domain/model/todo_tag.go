package model

import (
)

type TodoTag struct {
	ID    uint    `json:"id" param:"id"`
	TodoId uint `json:"todo_id"`
	TagId uint `json:"tag_id"`
	Tag Tag `json:"tag"`
	CreatedAt time.Time `json:"created_at"`
}
