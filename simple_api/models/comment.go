package models

type CommentModel struct {
	Subject string `json:"subject" validate:"required"`
	Comment string `json:"comment" validate:"required"`
}
