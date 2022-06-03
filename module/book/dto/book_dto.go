package dto

import (
	"encoding/json"
)

type BookDto struct {
	Title       string      `json:"title" binding:"required"`
	Price       json.Number `json:"price" binding:"required,number"`
	Description string      `json:"description"`
	SubTitle    string      `json:"sub_title"`
}
