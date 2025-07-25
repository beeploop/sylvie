package models

import "time"

type Job struct {
	ID        string
	FilePath  string
	MediaType MediaType
	CreatedAt time.Time
}
