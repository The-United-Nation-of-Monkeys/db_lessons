package auth

import "fmt"

func GetPermissions(apiVersion int) map[string][]string {
	routeConfigs := map[string][]string{
		fmt.Sprintf("GET /api/v%d/students", apiVersion):        {"student", "teacher", "admin"},
		fmt.Sprintf("PUT /api/v%d/students/:id", apiVersion):    {"student", "admin"},
		fmt.Sprintf("GET /api/v%d/students/:id", apiVersion):    {"student", "teacher", "admin"},
		fmt.Sprintf("DELETE /api/v%d/students/:id", apiVersion): {"student", "admin"},

		fmt.Sprintf("GET /api/v%d/teachers", apiVersion):        {"student", "teacher", "admin"},
		fmt.Sprintf("PUT /api/v%d/teachers/:id", apiVersion):    {"teacher", "admin"},
		fmt.Sprintf("GET /api/v%d/teachers/:id", apiVersion):    {"student", "teacher", "admin"},
		fmt.Sprintf("DELETE /api/v%d/teachers/:id", apiVersion): {"teacher", "admin"},

		fmt.Sprintf("GET /api/v%d/categories", apiVersion):        {"student", "teacher", "admin"},
		fmt.Sprintf("POST /api/v%d/categories", apiVersion):       {"teacher", "admin"},
		fmt.Sprintf("GET /api/v%d/categories/:id", apiVersion):    {"student", "teacher", "admin"},
		fmt.Sprintf("PUT /api/v%d/categories/:id", apiVersion):    {"teacher", "admin"},
		fmt.Sprintf("DELETE /api/v%d/categories/:id", apiVersion): {"teacher", "admin"},

		fmt.Sprintf("GET /api/v%d/currencies", apiVersion):        {"student", "teacher", "admin"},
		fmt.Sprintf("POST /api/v%d/currencies", apiVersion):       {"teacher", "admin"},
		fmt.Sprintf("GET /api/v%d/currencies/:id", apiVersion):    {"student", "teacher", "admin"},
		fmt.Sprintf("PUT /api/v%d/currencies/:id", apiVersion):    {"teacher", "admin"},
		fmt.Sprintf("DELETE /api/v%d/currencies/:id", apiVersion): {"teacher", "admin"},

		fmt.Sprintf("GET /api/v%d/levels", apiVersion):        {"student", "teacher", "admin"},
		fmt.Sprintf("POST /api/v%d/levels", apiVersion):       {"teacher", "admin"},
		fmt.Sprintf("GET /api/v%d/levels/:id", apiVersion):    {"student", "teacher", "admin"},
		fmt.Sprintf("PUT /api/v%d/levels/:id", apiVersion):    {"teacher", "admin"},
		fmt.Sprintf("DELETE /api/v%d/levels/:id", apiVersion): {"teacher", "admin"},

		fmt.Sprintf("GET /api/v%d/subcategories", apiVersion):        {"student", "teacher", "admin"},
		fmt.Sprintf("POST /api/v%d/subcategories", apiVersion):       {"teacher", "admin"},
		fmt.Sprintf("GET /api/v%d/subcategories/:id", apiVersion):    {"student", "teacher", "admin"},
		fmt.Sprintf("PUT /api/v%d/subcategories/:id", apiVersion):    {"teacher", "admin"},
		fmt.Sprintf("DELETE /api/v%d/subcategories/:id", apiVersion): {"teacher", "admin"},

		fmt.Sprintf("GET /api/v%d/materials", apiVersion):        {"student", "teacher", "admin"},
		fmt.Sprintf("POST /api/v%d/materials", apiVersion):       {"teacher", "admin"},
		fmt.Sprintf("GET /api/v%d/materials/:id", apiVersion):    {"student", "teacher", "admin"},
		fmt.Sprintf("PUT /api/v%d/materials/:id", apiVersion):    {"teacher", "admin"},
		fmt.Sprintf("DELETE /api/v%d/materials/:id", apiVersion): {"teacher", "admin"},

		fmt.Sprintf("GET /api/v%d/homeworks", apiVersion):        {"student", "teacher", "admin"},
		fmt.Sprintf("POST /api/v%d/homeworks", apiVersion):       {"teacher", "admin"},
		fmt.Sprintf("GET /api/v%d/homeworks/:id", apiVersion):    {"student", "teacher", "admin"},
		fmt.Sprintf("PUT /api/v%d/homeworks/:id", apiVersion):    {"teacher", "admin"},
		fmt.Sprintf("DELETE /api/v%d/homeworks/:id", apiVersion): {"teacher", "admin"},

		fmt.Sprintf("GET /api/v%d/lessons", apiVersion):        {"student", "teacher", "admin"},
		fmt.Sprintf("POST /api/v%d/lessons", apiVersion):       {"teacher", "admin"},
		fmt.Sprintf("GET /api/v%d/lessons/:id", apiVersion):    {"student", "teacher", "admin"},
		fmt.Sprintf("PUT /api/v%d/lessons/:id", apiVersion):    {"teacher", "admin"},
		fmt.Sprintf("DELETE /api/v%d/lessons/:id", apiVersion): {"teacher", "admin"},

		fmt.Sprintf("GET /api/v%d/courses", apiVersion):        {"student", "teacher", "admin"},
		fmt.Sprintf("POST /api/v%d/courses", apiVersion):       {"teacher", "admin"},
		fmt.Sprintf("GET /api/v%d/courses/:id", apiVersion):    {"student", "teacher", "admin"},
		fmt.Sprintf("PUT /api/v%d/courses/:id", apiVersion):    {"teacher", "admin"},
		fmt.Sprintf("DELETE /api/v%d/courses/:id", apiVersion): {"teacher", "admin"},

		fmt.Sprintf("GET /api/v%d/tasks", apiVersion):        {"student", "teacher", "admin"},
		fmt.Sprintf("POST /api/v%d/tasks", apiVersion):       {"teacher", "admin"},
		fmt.Sprintf("GET /api/v%d/tasks/:id", apiVersion):    {"student", "teacher", "admin"},
		fmt.Sprintf("PUT /api/v%d/tasks/:id", apiVersion):    {"teacher", "admin"},
		fmt.Sprintf("DELETE /api/v%d/tasks/:id", apiVersion): {"teacher", "admin"},

		fmt.Sprintf("GET /api/v%d/status-homeworks", apiVersion):        {"student", "teacher", "admin"},
		fmt.Sprintf("POST /api/v%d/status-homeworks", apiVersion):       {"teacher", "admin"},
		fmt.Sprintf("GET /api/v%d/status-homeworks/:id", apiVersion):    {"student", "teacher", "admin"},
		fmt.Sprintf("PUT /api/v%d/status-homeworks/:id", apiVersion):    {"teacher", "admin"},
		fmt.Sprintf("DELETE /api/v%d/status-homeworks/:id", apiVersion): {"teacher", "admin"},

		fmt.Sprintf("GET /api/v%d/status-answers", apiVersion):        {"student", "teacher", "admin"},
		fmt.Sprintf("POST /api/v%d/status-answers", apiVersion):       {"teacher", "admin"},
		fmt.Sprintf("GET /api/v%d/status-answers/:id", apiVersion):    {"student", "teacher", "admin"},
		fmt.Sprintf("PUT /api/v%d/status-answers/:id", apiVersion):    {"teacher", "admin"},
		fmt.Sprintf("DELETE /api/v%d/status-answers/:id", apiVersion): {"teacher", "admin"},

		fmt.Sprintf("GET /api/v%d/status-transactions", apiVersion):        {"student", "teacher", "admin"},
		fmt.Sprintf("POST /api/v%d/status-transactions", apiVersion):       {"admin"},
		fmt.Sprintf("GET /api/v%d/status-transactions/:id", apiVersion):    {"student", "teacher", "admin"},
		fmt.Sprintf("PUT /api/v%d/status-transactions/:id", apiVersion):    {"admin"},
		fmt.Sprintf("DELETE /api/v%d/status-transactions/:id", apiVersion): {"admin"},

		fmt.Sprintf("GET /api/v%d/student-answers", apiVersion):        {"student", "teacher", "admin"},
		fmt.Sprintf("POST /api/v%d/student-answers", apiVersion):       {"student", "admin"},
		fmt.Sprintf("GET /api/v%d/student-answers/:id", apiVersion):    {"student", "teacher", "admin"},
		fmt.Sprintf("PUT /api/v%d/student-answers/:id", apiVersion):    {"student", "admin"},
		fmt.Sprintf("DELETE /api/v%d/student-answers/:id", apiVersion): {"student", "admin"},

		fmt.Sprintf("GET /api/v%d/homework-results", apiVersion):        {"student", "teacher", "admin"},
		fmt.Sprintf("POST /api/v%d/homework-results", apiVersion):       {"teacher", "admin"},
		fmt.Sprintf("GET /api/v%d/homework-results/:id", apiVersion):    {"student", "teacher", "admin"},
		fmt.Sprintf("PUT /api/v%d/homework-results/:id", apiVersion):    {"teacher", "admin"},
		fmt.Sprintf("DELETE /api/v%d/homework-results/:id", apiVersion): {"teacher", "admin"},

		fmt.Sprintf("GET /api/v%d/transactions", apiVersion):                     {"student", "teacher", "admin"},
		fmt.Sprintf("POST /api/v%d/transactions", apiVersion):                    {"student", "admin"},
		fmt.Sprintf("GET /api/v%d/transactions/report", apiVersion):              {"student", "teacher", "admin"},
		fmt.Sprintf("POST /api/v%d/transactions/bulk-update-status", apiVersion): {"admin"},
		fmt.Sprintf("GET /api/v%d/transactions/:id", apiVersion):                 {"student", "teacher", "admin"},
		fmt.Sprintf("PUT /api/v%d/transactions/:id", apiVersion):                 {"student", "admin"},
		fmt.Sprintf("DELETE /api/v%d/transactions/:id", apiVersion):              {"student", "admin"},

		fmt.Sprintf("GET /api/v%d/reports/student-category-stats", apiVersion):                {"student", "teacher", "admin"},
		fmt.Sprintf("GET /api/v%d/reports/student-category-stats/students/:id", apiVersion):   {"student", "teacher", "admin"},
		fmt.Sprintf("GET /api/v%d/reports/student-category-stats/categories/:id", apiVersion): {"student", "teacher", "admin"},

		fmt.Sprintf("GET /api/v%d/teachers-courses", apiVersion):                           {"student", "teacher", "admin"},
		fmt.Sprintf("POST /api/v%d/teachers-courses", apiVersion):                          {"teacher", "admin"},
		fmt.Sprintf("GET /api/v%d/teachers-courses/:course_id/:teacher_id", apiVersion):    {"student", "teacher", "admin"},
		fmt.Sprintf("DELETE /api/v%d/teachers-courses/:course_id/:teacher_id", apiVersion): {"teacher", "admin"},

		fmt.Sprintf("GET /api/v%d/course-lessons", apiVersion):        {"student", "teacher", "admin"},
		fmt.Sprintf("POST /api/v%d/course-lessons", apiVersion):       {"teacher", "admin"},
		fmt.Sprintf("GET /api/v%d/course-lessons/:id", apiVersion):    {"student", "teacher", "admin"},
		fmt.Sprintf("DELETE /api/v%d/course-lessons/:id", apiVersion): {"teacher", "admin"},

		fmt.Sprintf("GET /api/v%d/lessons-materials", apiVersion):                            {"student", "teacher", "admin"},
		fmt.Sprintf("POST /api/v%d/lessons-materials", apiVersion):                           {"teacher", "admin"},
		fmt.Sprintf("GET /api/v%d/lessons-materials/:lesson_id/:material_id", apiVersion):    {"student", "teacher", "admin"},
		fmt.Sprintf("DELETE /api/v%d/lessons-materials/:lesson_id/:material_id", apiVersion): {"teacher", "admin"},

		fmt.Sprintf("GET /api/v%d/lesson-homeworks", apiVersion):                            {"student", "teacher", "admin"},
		fmt.Sprintf("POST /api/v%d/lesson-homeworks", apiVersion):                           {"teacher", "admin"},
		fmt.Sprintf("GET /api/v%d/lesson-homeworks/:lesson_id/:homework_id", apiVersion):    {"student", "teacher", "admin"},
		fmt.Sprintf("DELETE /api/v%d/lesson-homeworks/:lesson_id/:homework_id", apiVersion): {"teacher", "admin"},

		fmt.Sprintf("GET /api/v%d/homeworks-tasks", apiVersion):                          {"student", "teacher", "admin"},
		fmt.Sprintf("POST /api/v%d/homeworks-tasks", apiVersion):                         {"teacher", "admin"},
		fmt.Sprintf("GET /api/v%d/homeworks-tasks/:task_id/:homework_id", apiVersion):    {"student", "teacher", "admin"},
		fmt.Sprintf("DELETE /api/v%d/homeworks-tasks/:task_id/:homework_id", apiVersion): {"teacher", "admin"},

		fmt.Sprintf("GET /api/v%d/transactions-courses", apiVersion):                               {"student", "teacher", "admin"},
		fmt.Sprintf("POST /api/v%d/transactions-courses", apiVersion):                              {"student", "teacher", "admin"},
		fmt.Sprintf("GET /api/v%d/transactions-courses/:transaction_id/:course_id", apiVersion):    {"student", "teacher", "admin"},
		fmt.Sprintf("DELETE /api/v%d/transactions-courses/:transaction_id/:course_id", apiVersion): {"student", "teacher", "admin"},

		fmt.Sprintf("POST /api/v%d/sql/execute", apiVersion): {"admin"},
	}

	return routeConfigs
}
