package core

import (
	"context"
	"github.com/google/uuid"
	"time"
)

type Service interface {
	GetRecommendationByUserID(ctx context.Context, id uuid.UUID) (*Recommendation, error)
}

type TrackingServerClient interface {
	SendNewPost(ctx context.Context, post *Post) error
}

type Recommendation struct {
	Posts []Post
}

type Post struct {
	ID        uuid.UUID
	CreatedAt time.Time
	Title     string
	Text      string
}

type service struct {
	trackingServerClient TrackingServerClient
}

var _ Service = (*service)(nil)

func New(trackingServerclient TrackingServerClient) Service {
	return &service{
		trackingServerClient: trackingServerclient,
	}
}

func (s *service) GetRecommendationByUserID(ctx context.Context, id uuid.UUID) (*Recommendation, error) {
	// TODO: implement me
	return nil, nil
}
