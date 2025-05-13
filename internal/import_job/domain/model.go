package domain

import "time"

type ImportJobFile struct {
	JobID int
	Path  string
	Time  time.Time
}
type WorkerJob struct {
	ID    string
	RunID string
}
