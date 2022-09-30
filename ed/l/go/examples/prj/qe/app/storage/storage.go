package storage

import (
	"context"

	"quizengine/app/payload"
)

// Storage describes data storage.
type Storage interface {
	// CreateUser create user.
	CreateUser(ctx context.Context, input *payload.UserEntity) error
	// GetUserByEmail gets user by email address.
	GetUserByEmail(ctx context.Context, email string) (*payload.UserEntity, error)

	// GetQuiz gets quiz.
	GetQuiz(ctx context.Context, id string) (*payload.QuizEntity, error)
	// ListQuizzes gets quizzes.
	ListQuizzes(ctx context.Context, input *payload.ListQuizzesRequest) ([]*payload.QuizEntity, error)
	// UpsertQuiz performs insert/update for quiz.
	UpsertQuiz(ctx context.Context, input *payload.QuizEntity) error
	// DeleteQuiz deletes quiz.
	DeleteQuiz(ctx context.Context, input *payload.DeleteQuizRequest) error

	// CreateSubmission creates submission.
	CreateSubmission(ctx context.Context, input *payload.SubmissionEntity) error
	// ListSubmissions gets submissions.
	ListSubmissions(ctx context.Context, input *payload.ListSubmissionsRequest) ([]*payload.SubmissionEntity, error)
}
