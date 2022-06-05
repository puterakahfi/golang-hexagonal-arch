package dto

import (
	"encoding/json"
)

type BookRequestDto struct {
	Title       string      `json:"title" binding:"required"`
	Price       json.Number `json:"price" binding:"required,number"`
	Description string      `json:"description"`
	SubTitle    string      `json:"sub_title"`
}

type BookResponseDto struct {
	ID          string      `json:"id"`
	Title       string      `json:"title"`
	Price       json.Number `json:"price"`
	Discount    int         `json:"discount"`
	Description string      `json:"description"`
	SubTitle    string      `json:"sub_title"`
}

type BookFilter struct {
	Title       string      `json:"title"`
	Price       json.Number `json:"price"`
	Description string      `json:"description"`
	SubTitle    string      `json:"sub_title"`
}
