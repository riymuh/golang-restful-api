package user

import "encoding/json"

type UserRequest struct {
	Name       string `json:"name" binding:"required"`
	Email      string
	StatusName string      `json:"status_name"`
	Phone      json.Number `json:"phone" binding:"required,number"`
}
