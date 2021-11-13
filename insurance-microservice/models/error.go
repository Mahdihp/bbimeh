package models

// Error defines the response error
type Error struct {
	Code    int    `json:"code" example:"200"`
	Message string `json:"message" example:"Error message"`
}
