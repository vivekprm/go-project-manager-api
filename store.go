package main

import "database/sql"

// Make it an interface to easily inject in our services or mock in test
type Store interface {
	// Users
	CreateUser() error
}

// Repository to communicate with database
type Storage struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Storage {
	return &Storage{
		db: db,
	}
}

func (s *Storage) CreateUser() error {
	return nil
}
