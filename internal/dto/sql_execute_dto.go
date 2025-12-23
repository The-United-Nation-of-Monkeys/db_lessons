package dto

// ExecuteSQLDTO represents a request to execute a SQL query
type ExecuteSQLDTO struct {
	Query string `json:"query" validate:"required" example:"SELECT * FROM student LIMIT 10"`
}

// ExecuteSQLResponseDTO represents the result of SQL query execution
type ExecuteSQLResponseDTO struct {
	RowsAffected int64                    `json:"rows_affected" example:"10"`
	Columns      []string                 `json:"columns,omitempty" example:"student_id,name,surname,email"`
	Rows         []map[string]interface{} `json:"rows,omitempty"`
}
