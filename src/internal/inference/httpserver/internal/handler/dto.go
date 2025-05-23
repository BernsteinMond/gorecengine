package handler

import "github.com/BernsteinMond/gorecengine/src/internal/inference/service"

type postDTO struct {
	ID        string `json:"id"`
	CreatedAt string `json:"created_at"`
	Title     string `json:"title"`
	Text      string `json:"text"`
}

func fromDomainToPostDTO(post *service.Post) postDTO {
	return postDTO{
		ID:        post.ID.String(),
		CreatedAt: post.CreatedAt.String(),
		Title:     post.Title,
		Text:      post.Text,
	}
}
