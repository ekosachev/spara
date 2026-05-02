package dto

import "time"

type Model struct {
	ID        uint
	CreatedAt time.Time
	UpdatedAt time.Time
}

type ApiResponse struct {
	Success bool   `json:"success"`
	Error   string `json:"error,omitempty"`
	Data    any    `json:"data,omitempty"`
}
