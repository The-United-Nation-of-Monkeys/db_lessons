package dto

import "time"

type CreateCourseDTO struct {
	Name        string     `json:"name" validate:"required"`
	Description string     `json:"description"`
	CategoryID  *int       `json:"category_id"`
	StartDate   *time.Time `json:"start_date"`
	EndDate     *time.Time `json:"end_date"`
	Price       int        `json:"price" validate:"gte=0"`
	CurrencyID  *int       `json:"currency_id"`
}

type UpdateCourseDTO struct {
	Name        *string    `json:"name"`
	Description *string    `json:"description"`
	CategoryID  *int       `json:"category_id"`
	StartDate   *time.Time `json:"start_date"`
	EndDate     *time.Time `json:"end_date"`
	Price       *int       `json:"price" validate:"omitempty,gte=0"`
	CurrencyID  *int       `json:"currency_id"`
}

type CourseDTO struct {
	CourseID    int        `json:"course_id"`
	Name        string     `json:"name"`
	Description string     `json:"description"`
	CategoryID  *int       `json:"category_id"`
	StartDate   *time.Time `json:"start_date"`
	EndDate     *time.Time `json:"end_date"`
	Price       int        `json:"price"`
	CurrencyID  *int       `json:"currency_id"`
}


