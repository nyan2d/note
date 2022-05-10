package model

import (
	"time"
)

type Note struct {
	ID        int       `json:"id"`
	Folder    int       `json:"folder"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
	EditedAt  time.Time `json:"edited_at"`
}
