package main

import "context"

type Service struct {
	store OrderStore
}

func NewService(store OrderStore) *Service {
	return &Service{store}

}

func (s *Service) CreateOrder(ctx context.Context) *Service {
	return nil
}