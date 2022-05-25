package model

import (
)

type Tag struct {
	ID    uint    `json:"id" param:"id"`
	TagId uint `json:"tag_id"`
	Text string `json:"text"`
	CreatedAt time.Time `json:"created_at"`
}
