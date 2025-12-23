package auth

import "fmt"

// GetPermissions returns route configurations for authentication
// Format: "METHOD /path" or "/path" (for all methods)
// Routes not in this map are public (no auth required)
func GetPermissions(apiVersion int) map[string][]string {
	routeConfigs := map[string][]string{
		// Students - GET requires auth, POST is public (registration)
		fmt.Sprintf("GET /api/v%d/students", apiVersion):        {"student", "teacher", "admin"},
		fmt.Sprintf("PUT /api/v%d/students/:id", apiVersion):    {"student", "admin"},
		fmt.Sprintf("GET /api/v%d/students/:id", apiVersion):    {"student", "teacher", "admin"},
		fmt.Sprintf("DELETE /api/v%d/students/:id", apiVersion): {"student", "admin"},

		// Teachers - GET requires auth, POST is public (registration)
		fmt.Sprintf("GET /api/v%d/teachers", apiVersion):        {"student", "teacher", "admin"},
		fmt.Sprintf("PUT /api/v%d/teachers/:id", apiVersion):    {"teacher", "admin"},
		fmt.Sprintf("GET /api/v%d/teachers/:id", apiVersion):    {"student", "teacher", "admin"},
		fmt.Sprintf("DELETE /api/v%d/teachers/:id", apiVersion): {"teacher", "admin"},

		// Categories
		fmt.Sprintf("GET /api/v%d/categories", apiVersion):        {"student", "teacher", "admin"},
		fmt.Sprintf("POST /api/v%d/categories", apiVersion):       {"teacher", "admin"},
		fmt.Sprintf("GET /api/v%d/categories/:id", apiVersion):    {"student", "teacher", "admin"},
		fmt.Sprintf("PUT /api/v%d/categories/:id", apiVersion):    {"teacher", "admin"},
		fmt.Sprintf("DELETE /api/v%d/categories/:id", apiVersion): {"teacher", "admin"},

		// Currencies
		fmt.Sprintf("GET /api/v%d/currencies", apiVersion):        {"student", "teacher", "admin"},
		fmt.Sprintf("POST /api/v%d/currencies", apiVersion):       {"teacher", "admin"},
		fmt.Sprintf("GET /api/v%d/currencies/:id", apiVersion):    {"student", "teacher", "admin"},
		fmt.Sprintf("PUT /api/v%d/currencies/:id", apiVersion):    {"teacher", "admin"},
		fmt.Sprintf("DELETE /api/v%d/currencies/:id", apiVersion): {"teacher", "admin"},

		// Levels
		fmt.Sprintf("GET /api/v%d/levels", apiVersion):        {"student", "teacher", "admin"},
		fmt.Sprintf("POST /api/v%d/levels", apiVersion):       {"teacher", "admin"},
		fmt.Sprintf("GET /api/v%d/levels/:id", apiVersion):    {"student", "teacher", "admin"},
		fmt.Sprintf("PUT /api/v%d/levels/:id", apiVersion):    {"teacher", "admin"},
		fmt.Sprintf("DELETE /api/v%d/levels/:id", apiVersion): {"teacher", "admin"},

		// Subcategories
		fmt.Sprintf("GET /api/v%d/subcategories", apiVersion):        {"student", "teacher", "admin"},
		fmt.Sprintf("POST /api/v%d/subcategories", apiVersion):       {"teacher", "admin"},
		fmt.Sprintf("GET /api/v%d/subcategories/:id", apiVersion):    {"student", "teacher", "admin"},
		fmt.Sprintf("PUT /api/v%d/subcategories/:id", apiVersion):    {"teacher", "admin"},
		fmt.Sprintf("DELETE /api/v%d/subcategories/:id", apiVersion): {"teacher", "admin"},

		// Materials
		fmt.Sprintf("GET /api/v%d/materials", apiVersion):        {"student", "teacher", "admin"},
		fmt.Sprintf("POST /api/v%d/materials", apiVersion):       {"teacher", "admin"},
		fmt.Sprintf("GET /api/v%d/materials/:id", apiVersion):    {"student", "teacher", "admin"},
		fmt.Sprintf("PUT /api/v%d/materials/:id", apiVersion):    {"teacher", "admin"},
		fmt.Sprintf("DELETE /api/v%d/materials/:id", apiVersion): {"teacher", "admin"},

		// Homeworks
		fmt.Sprintf("GET /api/v%d/homeworks", apiVersion):        {"student", "teacher", "admin"},
		fmt.Sprintf("POST /api/v%d/homeworks", apiVersion):       {"teacher", "admin"},
		fmt.Sprintf("GET /api/v%d/homeworks/:id", apiVersion):    {"student", "teacher", "admin"},
		fmt.Sprintf("PUT /api/v%d/homeworks/:id", apiVersion):    {"teacher", "admin"},
		fmt.Sprintf("DELETE /api/v%d/homeworks/:id", apiVersion): {"teacher", "admin"},

		// Lessons
		fmt.Sprintf("GET /api/v%d/lessons", apiVersion):        {"student", "teacher", "admin"},
		fmt.Sprintf("POST /api/v%d/lessons", apiVersion):       {"teacher", "admin"},
		fmt.Sprintf("GET /api/v%d/lessons/:id", apiVersion):    {"student", "teacher", "admin"},
		fmt.Sprintf("PUT /api/v%d/lessons/:id", apiVersion):    {"teacher", "admin"},
		fmt.Sprintf("DELETE /api/v%d/lessons/:id", apiVersion): {"teacher", "admin"},

		// Courses
		fmt.Sprintf("GET /api/v%d/courses", apiVersion):        {"student", "teacher", "admin"},
		fmt.Sprintf("POST /api/v%d/courses", apiVersion):       {"teacher", "admin"},
		fmt.Sprintf("GET /api/v%d/courses/:id", apiVersion):    {"student", "teacher", "admin"},
		fmt.Sprintf("PUT /api/v%d/courses/:id", apiVersion):    {"teacher", "admin"},
		fmt.Sprintf("DELETE /api/v%d/courses/:id", apiVersion): {"teacher", "admin"},

		// Tasks
		fmt.Sprintf("GET /api/v%d/tasks", apiVersion):        {"student", "teacher", "admin"},
		fmt.Sprintf("POST /api/v%d/tasks", apiVersion):       {"teacher", "admin"},
		fmt.Sprintf("GET /api/v%d/tasks/:id", apiVersion):    {"student", "teacher", "admin"},
		fmt.Sprintf("PUT /api/v%d/tasks/:id", apiVersion):    {"teacher", "admin"},
		fmt.Sprintf("DELETE /api/v%d/tasks/:id", apiVersion): {"teacher", "admin"},

		// Status Homeworks
		fmt.Sprintf("GET /api/v%d/status-homeworks", apiVersion):        {"student", "teacher", "admin"},
		fmt.Sprintf("POST /api/v%d/status-homeworks", apiVersion):       {"teacher", "admin"},
		fmt.Sprintf("GET /api/v%d/status-homeworks/:id", apiVersion):    {"student", "teacher", "admin"},
		fmt.Sprintf("PUT /api/v%d/status-homeworks/:id", apiVersion):    {"teacher", "admin"},
		fmt.Sprintf("DELETE /api/v%d/status-homeworks/:id", apiVersion): {"teacher", "admin"},

		// Status Answers
		fmt.Sprintf("GET /api/v%d/status-answers", apiVersion):        {"student", "teacher", "admin"},
		fmt.Sprintf("POST /api/v%d/status-answers", apiVersion):       {"student", "teacher", "admin"},
		fmt.Sprintf("GET /api/v%d/status-answers/:id", apiVersion):    {"student", "teacher", "admin"},
		fmt.Sprintf("PUT /api/v%d/status-answers/:id", apiVersion):    {"student", "teacher", "admin"},
		fmt.Sprintf("DELETE /api/v%d/status-answers/:id", apiVersion): {"student", "teacher", "admin"},

		// Status Transactions
		fmt.Sprintf("GET /api/v%d/status-transactions", apiVersion):        {"student", "teacher", "admin"},
		fmt.Sprintf("POST /api/v%d/status-transactions", apiVersion):       {"admin"},
		fmt.Sprintf("GET /api/v%d/status-transactions/:id", apiVersion):    {"student", "teacher", "admin"},
		fmt.Sprintf("PUT /api/v%d/status-transactions/:id", apiVersion):    {"admin"},
		fmt.Sprintf("DELETE /api/v%d/status-transactions/:id", apiVersion): {"admin"},

		// Student Answers
		fmt.Sprintf("GET /api/v%d/student-answers", apiVersion):        {"student", "teacher", "admin"},
		fmt.Sprintf("POST /api/v%d/student-answers", apiVersion):       {"student", "teacher", "admin"},
		fmt.Sprintf("GET /api/v%d/student-answers/:id", apiVersion):    {"student", "teacher", "admin"},
		fmt.Sprintf("PUT /api/v%d/student-answers/:id", apiVersion):    {"student", "teacher", "admin"},
		fmt.Sprintf("DELETE /api/v%d/student-answers/:id", apiVersion): {"student", "teacher", "admin"},

		// Homework Results
		fmt.Sprintf("GET /api/v%d/homework-results", apiVersion):        {"student", "teacher", "admin"},
		fmt.Sprintf("POST /api/v%d/homework-results", apiVersion):       {"student", "teacher", "admin"},
		fmt.Sprintf("GET /api/v%d/homework-results/:id", apiVersion):    {"student", "teacher", "admin"},
		fmt.Sprintf("PUT /api/v%d/homework-results/:id", apiVersion):    {"student", "teacher", "admin"},
		fmt.Sprintf("DELETE /api/v%d/homework-results/:id", apiVersion): {"student", "teacher", "admin"},

		// Transactions
		fmt.Sprintf("GET /api/v%d/transactions", apiVersion):                     {"student", "teacher", "admin"},
		fmt.Sprintf("POST /api/v%d/transactions", apiVersion):                    {"student", "teacher", "admin"},
		fmt.Sprintf("GET /api/v%d/transactions/report", apiVersion):              {"student", "teacher", "admin"},
		fmt.Sprintf("POST /api/v%d/transactions/bulk-update-status", apiVersion): {"admin"},
		fmt.Sprintf("GET /api/v%d/transactions/:id", apiVersion):                 {"student", "teacher", "admin"},
		fmt.Sprintf("PUT /api/v%d/transactions/:id", apiVersion):                 {"student", "teacher", "admin"},
		fmt.Sprintf("DELETE /api/v%d/transactions/:id", apiVersion):              {"student", "teacher", "admin"},

		// Reports
		fmt.Sprintf("GET /api/v%d/reports/student-category-stats", apiVersion):                {"student", "teacher", "admin"},
		fmt.Sprintf("GET /api/v%d/reports/student-category-stats/students/:id", apiVersion):   {"student", "teacher", "admin"},
		fmt.Sprintf("GET /api/v%d/reports/student-category-stats/categories/:id", apiVersion): {"student", "teacher", "admin"},

		// Relations - Teachers Courses
		fmt.Sprintf("GET /api/v%d/teachers-courses", apiVersion):                           {"student", "teacher", "admin"},
		fmt.Sprintf("POST /api/v%d/teachers-courses", apiVersion):                          {"student", "teacher", "admin"},
		fmt.Sprintf("GET /api/v%d/teachers-courses/:course_id/:teacher_id", apiVersion):    {"student", "teacher", "admin"},
		fmt.Sprintf("DELETE /api/v%d/teachers-courses/:course_id/:teacher_id", apiVersion): {"student", "teacher", "admin"},

		// Relations - Course Lessons
		fmt.Sprintf("GET /api/v%d/course-lessons", apiVersion):        {"student", "teacher", "admin"},
		fmt.Sprintf("POST /api/v%d/course-lessons", apiVersion):       {"student", "teacher", "admin"},
		fmt.Sprintf("GET /api/v%d/course-lessons/:id", apiVersion):    {"student", "teacher", "admin"},
		fmt.Sprintf("DELETE /api/v%d/course-lessons/:id", apiVersion): {"student", "teacher", "admin"},

		// Relations - Lessons Materials
		fmt.Sprintf("GET /api/v%d/lessons-materials", apiVersion):                            {"student", "teacher", "admin"},
		fmt.Sprintf("POST /api/v%d/lessons-materials", apiVersion):                           {"student", "teacher", "admin"},
		fmt.Sprintf("GET /api/v%d/lessons-materials/:lesson_id/:material_id", apiVersion):    {"student", "teacher", "admin"},
		fmt.Sprintf("DELETE /api/v%d/lessons-materials/:lesson_id/:material_id", apiVersion): {"student", "teacher", "admin"},

		// Relations - Lesson Homeworks
		fmt.Sprintf("GET /api/v%d/lesson-homeworks", apiVersion):                            {"student", "teacher", "admin"},
		fmt.Sprintf("POST /api/v%d/lesson-homeworks", apiVersion):                           {"student", "teacher", "admin"},
		fmt.Sprintf("GET /api/v%d/lesson-homeworks/:lesson_id/:homework_id", apiVersion):    {"student", "teacher", "admin"},
		fmt.Sprintf("DELETE /api/v%d/lesson-homeworks/:lesson_id/:homework_id", apiVersion): {"student", "teacher", "admin"},

		// Relations - Homeworks Tasks
		fmt.Sprintf("GET /api/v%d/homeworks-tasks", apiVersion):                          {"student", "teacher", "admin"},
		fmt.Sprintf("POST /api/v%d/homeworks-tasks", apiVersion):                         {"student", "teacher", "admin"},
		fmt.Sprintf("GET /api/v%d/homeworks-tasks/:task_id/:homework_id", apiVersion):    {"student", "teacher", "admin"},
		fmt.Sprintf("DELETE /api/v%d/homeworks-tasks/:task_id/:homework_id", apiVersion): {"student", "teacher", "admin"},

		// Relations - Transactions Courses
		fmt.Sprintf("GET /api/v%d/transactions-courses", apiVersion):                               {"student", "teacher", "admin"},
		fmt.Sprintf("POST /api/v%d/transactions-courses", apiVersion):                              {"student", "teacher", "admin"},
		fmt.Sprintf("GET /api/v%d/transactions-courses/:transaction_id/:course_id", apiVersion):    {"student", "teacher", "admin"},
		fmt.Sprintf("DELETE /api/v%d/transactions-courses/:transaction_id/:course_id", apiVersion): {"student", "teacher", "admin"},
	}

	return routeConfigs
}
