package dto

type CreateTransactionDTO struct {
	StudentID  int  `json:"student_id" validate:"required"`
	StatusID   *int `json:"status_id"`
	TotalPrice int  `json:"total_price" validate:"gte=0"`
}

type UpdateTransactionDTO struct {
	StudentID  *int `json:"student_id"`
	StatusID   *int `json:"status_id"`
	TotalPrice *int `json:"total_price" validate:"omitempty,gte=0"`
}

type TransactionDTO struct {
	TransactionID int  `json:"transaction_id"`
	StudentID     int  `json:"student_id"`
	StatusID      *int `json:"status_id"`
	TotalPrice    int  `json:"total_price"`
}

type TransactionReportRequestDTO struct {
	StatusName string `json:"status_name,omitempty"`
	MinTotal   int    `json:"min_total,omitempty"`
	MaxTotal   int    `json:"max_total,omitempty"`
}

type TransactionReportDTO struct {
	TransactionsID int    `json:"transactions_id"`
	StudentID      int    `json:"student_id"`
	StudentName    string `json:"student_name"`
	StudentSurname string `json:"student_surname"`
	StatusName     string `json:"status_name"`
	TotalPrice     int    `json:"total_price"`
	CoursesNames   string `json:"courses_names"`
}
