package handler

import (
	"github.com/BernsteinMond/gorecengine/inference/src/internal/core"
)

type postDTO struct {
	ID        string `json:"id"`
	CreatedAt string `json:"created_at"`
	Title     string `json:"title"`
	Text      string `json:"text"`
}

func fromDomainToPostDTO(post *core.Post) postDTO {
	return postDTO{
		ID:        post.ID.String(),
		CreatedAt: post.CreatedAt.String(),
		Title:     post.Title,
		Text:      post.Text,
	}
}
