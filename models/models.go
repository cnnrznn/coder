package models

import "time"

type Problem struct {
	ID        string    `json:"id"`
	Statement string    `json:"statement"`
	CreatedAt time.Time `json:"created_at"`
}

type TestCase struct {
	ID        string      `json:"id"`
	ProblemID string      `json:"problem_id"`
	Input     interface{} `json:"input"`
	Output    interface{} `json:"output"`
}
