package dto

import "time"

type CreateLessonDTO struct {
	Name        string     `json:"name" validate:"required"`
	Description string     `json:"description"`
	OpenTime    *time.Time `json:"open_time"`
	VideoLink   string     `json:"video_link"`
	LessonText  string     `json:"lesson_text"`
}

type UpdateLessonDTO struct {
	Name        *string    `json:"name"`
	Description *string    `json:"description"`
	OpenTime    *time.Time `json:"open_time"`
	VideoLink   *string    `json:"video_link"`
	LessonText  *string    `json:"lesson_text"`
}

type LessonDTO struct {
	LessonID    int        `json:"lesson_id"`
	Name        string     `json:"name"`
	Description string     `json:"description"`
	OpenTime    *time.Time `json:"open_time"`
	VideoLink   string     `json:"video_link"`
	LessonText  string     `json:"lesson_text"`
}


