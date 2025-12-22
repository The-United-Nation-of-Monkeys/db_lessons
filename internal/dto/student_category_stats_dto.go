package dto

type StudentCategoryStatsDTO struct {
	StudentID        int     `json:"student_id"`
	StudentName      *string `json:"student_name"`
	StudentSurname   *string `json:"student_surname"`
	CategoryID       *int    `json:"category_id"`
	CategoryName     *string `json:"category_name"`
	SolvedTasksCount int     `json:"solved_tasks_count"`
	PointsEarned     int     `json:"points_earned"`
	PointsPossible   int     `json:"points_possible"`
	SuccessPercent   float64 `json:"success_percent"`
}

