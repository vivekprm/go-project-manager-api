package main

import "database/sql"

// Make it an interface to easily inject in our services or mock in test
type Store interface {
	// Users
	CreateUser() error
	CreateTask(t *Task) (*Task, error)
	GetTask(id string) (*Task, error)
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

func (s *Storage) CreateTask(t *Task) (*Task, error) {
	rows, err := s.db.Exec("INSERT INTO tasks (name, status, project_id, assigned_to) VALUES (?, ?, ?, ?)", t.Name, t.Status, t.ProjectID, t.AssignedToID)
	if err != nil {
		return nil, err
	}
	id, err := rows.LastInsertId()
	if err != nil {
		return nil, err
	}
	t.ID = id
	return t, nil
}

func (s *Storage) GetTask(id string) (*Task, error) {
	var t *Task
	err := s.db.QueryRow("SELECT id, name, status, project_id, assigned_to, created_at FROM tasks WHERE id=?", id).Scan(&t.ID, &t.Name, &t.Status, &t.ProjectID, &t.AssignedToID, &t.CreatedAt)
	if err != nil {
		return nil, err
	}
	return t, nil
}
