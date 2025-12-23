package repository

import (
	"context"

	"github.com/The-United-Nation-of-Monkeys/db_lessons/internal/dto"
	"github.com/jackc/pgx/v5"
)

type SQLExecuteRepositoryInterface interface {
	ExecuteSQL(ctx context.Context, conn *pgx.Conn, query string) (*dto.ExecuteSQLResponseDTO, error)
}

type SQLExecuteRepository struct{}

func NewSQLExecuteRepository() *SQLExecuteRepository {
	return &SQLExecuteRepository{}
}

func (r *SQLExecuteRepository) ExecuteSQL(ctx context.Context, conn *pgx.Conn, query string) (*dto.ExecuteSQLResponseDTO, error) {
	// Определяем тип запроса по первому слову
	queryUpper := query
	if len(query) > 0 {
		// Берем первые 10 символов для проверки
		if len(query) > 10 {
			queryUpper = query[:10]
		}
	}

	// Проверяем, является ли это SELECT запросом
	isSelect := false
	if len(queryUpper) >= 6 {
		firstWord := queryUpper[:6]
		if firstWord == "SELECT" || firstWord == "select" || firstWord == "Select" {
			isSelect = true
		}
	}

	if isSelect {
		// Для SELECT используем Query
		rows, err := conn.Query(ctx, query)
		if err != nil {
			return nil, err
		}
		defer rows.Close()

		// Получаем имена колонок
		fieldDescriptions := rows.FieldDescriptions()
		columns := make([]string, len(fieldDescriptions))
		for i, fd := range fieldDescriptions {
			columns[i] = string(fd.Name)
		}

		// Читаем результаты
		var resultRows []map[string]interface{}
		rowsAffected := int64(0)

		for rows.Next() {
			values := make([]interface{}, len(columns))
			valuePtrs := make([]interface{}, len(columns))
			for i := range values {
				valuePtrs[i] = &values[i]
			}

			if err := rows.Scan(valuePtrs...); err != nil {
				return nil, err
			}

			row := make(map[string]interface{})
			for i, col := range columns {
				row[col] = values[i]
			}
			resultRows = append(resultRows, row)
			rowsAffected++
		}

		if err := rows.Err(); err != nil {
			return nil, err
		}

		return &dto.ExecuteSQLResponseDTO{
			RowsAffected: rowsAffected,
			Columns:      columns,
			Rows:         resultRows,
		}, nil
	} else {
		// Для INSERT, UPDATE, DELETE, CREATE, DROP и т.д. используем Exec
		commandTag, err := conn.Exec(ctx, query)
		if err != nil {
			return nil, err
		}

		return &dto.ExecuteSQLResponseDTO{
			RowsAffected: commandTag.RowsAffected(),
			Columns:      []string{},
			Rows:         []map[string]interface{}{},
		}, nil
	}
}
