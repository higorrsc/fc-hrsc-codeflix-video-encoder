package domain

import "time"

type Video struct {
	ID         string    `json:"encoded_video_folder"`
	ResourceID string    `json:"resource_id"`
	FilePath   string    `json:"file_path"`
	CreatedAt  time.Time `json:"created_at"`
}
