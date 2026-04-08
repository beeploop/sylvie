package models

import (
	"database/sql"
	"time"
)

type VideoStatus string

var (
	STATUS_UPLOADED   VideoStatus = "uploaded"
	STATUS_PROCESSING VideoStatus = "processing"
	STATUS_READY      VideoStatus = "ready"
	STATUS_FAILED     VideoStatus = "failed"
)

type Video struct {
	ID     string      `db:"id" json:"id"`
	Title  string      `db:"title" json:"title"`
	Status VideoStatus `db:"status" json:"status"`

	OriginalPath       string         `db:"original_path" json:"original_path"`
	MasterPlaylistPath sql.NullString `db:"master_playlist_path" json:"master_playlist_path"`
	ThumbnailPath      sql.NullString `db:"thumbnail_path" json:"thumbnail_path"`

	DurationSeconds sql.NullFloat64 `db:"duration_seconds" json:"duration_seconds"`
	Width           sql.NullInt64   `db:"width" json:"width"`
	Height          sql.NullInt64   `db:"height" json:"height"`

	CreatedAt   time.Time    `db:"created_at" json:"created_at"`
	ProcessedAt sql.NullTime `db:"processed_at" json:"processed_at"`
}

type NewVideo struct {
	ID           string      `db:"id"`
	Title        string      `db:"title"`
	OriginalPath string      `db:"original_path"`
	Status       VideoStatus `db:"status"`
}

type UpdateVideo struct {
	MasterPlaylistPath *string
	ThumbnailPath      *string
	DurationSeconds    *float64
	Width              *int
	Height             *int
	Status             *VideoStatus
	ProcessedAt        *time.Time
}
