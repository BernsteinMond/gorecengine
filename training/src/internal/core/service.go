package core

type Service interface{}

type service struct{}

var _ Service = (*service)(nil)
