package client

import (
	"context"
	"github.com/BernsteinMond/gorecengine/inference/src/internal/core"
	"net/http"
)

type TrackingServerClient struct {
	cl *http.Client
}

var _ core.TrackingServerClient = (*TrackingServerClient)(nil)

func New() *TrackingServerClient {
	return &TrackingServerClient{
		cl: http.DefaultClient,
	}
}

func (t *TrackingServerClient) SendNewPost(ctx context.Context, post *core.Post) error {
	// TODO: implement
	return nil
}
