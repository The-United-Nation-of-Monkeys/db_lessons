package dto

type ExecuteSQLDTO struct {
	Query string `json:"query" validate:"required" example:"SELECT * FROM student LIMIT 10"`
}

type ExecuteSQLResponseDTO struct {
	RowsAffected int64                    `json:"rows_affected" example:"10"`
	Columns      []string                 `json:"columns,omitempty" example:"student_id,name,surname,email"`
	Rows         []map[string]interface{} `json:"rows,omitempty"`
}
