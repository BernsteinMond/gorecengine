package client

import (
	"context"
	"github.com/BernsteinMond/gorecengine/src/internal/inference/service"
	"net/http"
)

type TrackingServerClient struct {
	cl *http.Client
}

var _ service.TrackingServerClient = (*TrackingServerClient)(nil)

func (t *TrackingServerClient) SendNewPost(ctx context.Context, post *service.Post) error {
	// TODO: implement
	return nil
}
