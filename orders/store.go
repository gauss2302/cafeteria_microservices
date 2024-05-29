package main

import "context"

type Store struct {
	// For Mongo Db

}

func NewStore() *Store {
	return &Store{}
}

func (s *Store) Create(ctx context.Context) error {
	return nil
}
