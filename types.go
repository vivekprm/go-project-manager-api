package main

type ErrorResponse struct {
	Error string `json:"error"`
}

type Task struct {
	ID           int64
	Name         string
	ProjectID    int
	AssignedToID int
	Status       string
	CreatedAt    string
}
