package payload

// AuthenticateUserRequest holds input parameters to authenticate user.
type AuthenticateUserRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

// AuthenticateUserResponse holds output result of user authentication.
type AuthenticateUserResponse struct {
	Token  string `json:"token"`
	UserID string `json:"user_id"`
}

// CreateUserRequest holds input parameters to create user.
type CreateUserRequest struct {
	Email     string `json:"email" validate:"required,email"`
	FirstName string `json:"first_name" validate:"required,alpha"`
	LastName  string `json:"last_name" validate:"required,alpha"`
	Password  string `json:"password" validate:"required"` // Additional validation in service.
}

// CreateUserResponse holds output result of user creation.
type CreateUserResponse struct{}

// ListQuizzesRequest holds input parameters to list quizzes.
type ListQuizzesRequest struct {
	ID     string `json:"-" validate:"-"`      // Only for internal usage.
	Author string `json:"author" validate:"-"` // Optional value, validation in service.
}

// ListQuizzesResponse holds output quizzes list.
type ListQuizzesResponse struct{}

// UpsertQuizRequest holds input parameters to upsert (insert/update) quiz.
type UpsertQuizRequest struct {
	ID        string                      `json:"quiz" validate:"required,alphanum"`
	Published bool                        `json:"published" validate:"-"` // Optional, false by default.
	Questions []UpsertQuizRequestQuestion `json:"questions" validate:"required"`
	Author    string                      `json:"author" validate:"required,alphanum"`
}

// UpsertQuizRequestQuestion holds data about question for quiz upsert.
type UpsertQuizRequestQuestion struct {
	Title   string                    `json:"title" validate:"required,alphanum"`
	Answers []UpsertQuizRequestAnswer `json:"answers" validate:"required"`
}

// UpsertQuizRequestAnswer holds data about answer for quiz upsert.
type UpsertQuizRequestAnswer struct {
	Title         string `json:"title" validate:"required,alphanum"`
	CorrectAnswer bool   `json:"correct_answer" validate:"-"` // Optional, false by default.
}

// UpsertQuizResponse holds output result of quiz upsert.
type UpsertQuizResponse struct{}

// DeleteQuizRequest holds input parameters to delete quiz.
type DeleteQuizRequest struct {
	Requester string `json:"requester" validate:"required,alphanum"`
	QuizID    string `json:"quiz_id" validate:"required"`
	DeletedAt int64  `json:"-" validate:"-"` // Only for internal usage.
}

// DeleteQuizResponse holds output result of quiz deletion.
type DeleteQuizResponse struct{}

// ListSubmissionsRequest holds input parameters to list submissions.
type ListSubmissionsRequest struct {
	Requester  string `json:"requester" validate:"required,alphanum"`
	UserID     string `json:"user_id" validate:"-"`     // Optional value, validation in service.
	QuizAuthor string `json:"quiz_author" validate:"-"` // Optional value, validation in service.
}

// ListSubmissionsResponse holds output submissions list.
type ListSubmissionsResponse struct{}

// CreateSubmissionRequest holds input parameters to create submissions.
type CreateSubmissionRequest struct {
	UserID  string  `json:"user_id" validate:"required"`
	QuizID  string  `json:"quiz_id" validate:"required"`
	Answers [][]int `json:"answers" validate:"required"`
}

// CreateSubmissionResponse holds output result of submission creation.
type CreateSubmissionResponse struct {
	TotalScore float64   `json:"total_score"`
	Questions  []float64 `json:"questions"`
}
