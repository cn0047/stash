package storage

import (
	"context"
	"fmt"
	"strings"

	"cloud.google.com/go/spanner"
	"github.com/mitchellh/mapstructure"
	"google.golang.org/api/iterator"

	"quizengine/app/payload"
)

const (
	// TableQuizzes holds quizzes table name.
	TableQuizzes = "Quizzes"

	// QueryInsertUser holds SQL query to insert user.
	QueryInsertUser = `
		INSERT INTO Users (ID, Email, FirstName, LastName, Password, CreatedAt)
		VALUES (@ID, @Email, @FirstName, @LastName, @Password, @CreatedAt)
	`

	// QuerySelectUserByEmail holds SQL query to select user by email.
	QuerySelectUserByEmail = `
		SELECT ID, Email, FirstName, LastName, Password, CreatedAt
		FROM Users
		WHERE Email = @Email
	`

	// QuerySelectQuizzes holds SQL query to select quizzes.
	QuerySelectQuizzes = `
		SELECT ID, Questions, Author, Published, UpdatedAt, Deleted, DeletedAt
		FROM Quizzes
		WHERE 1 = 1
	`

	// QueryMarkQuizAsDeleted holds SQL query to mark quiz as deleted.
	QueryMarkQuizAsDeleted = `
		UPDATE Quizzes
		SET Deleted = @Deleted, DeletedAt = @DeletedAt
		WHERE ID = @ID
	`

	// QueryInsertSubmission holds SQL query to insert submission.
	QueryInsertSubmission = `
		INSERT INTO Submissions (ID, UserID, QuizID, QuizAuthor, Score, SubmittedAt)
		VALUES (@ID, @UserID, @QuizID, @QuizAuthor, @Score, @SubmittedAt)
	`

	// QuerySelectSubmissions holds SQL query to select submission.
	QuerySelectSubmissions = `
		SELECT ID, UserID, QuizID, QuizAuthor, Score, SubmittedAt
		FROM Submissions
		WHERE 1 = 1
	`
)

// SpannerStorage represents GCP Spanner storage.
type SpannerStorage struct {
	client *spanner.Client
}

// NewSpannerStorage creates new GCP Spanner instance.
func NewSpannerStorage(db string) (*SpannerStorage, error) {
	c, err := spanner.NewClient(context.Background(), db)
	if err != nil {
		return nil, fmt.Errorf("failed to create new spanner client, err: %w", err)
	}

	s := &SpannerStorage{client: c}

	return s, nil
}

// update represents basic wrapper for Spanner's update method.
func (s *SpannerStorage) update(ctx context.Context, sql string, params map[string]interface{}) error {
	cb := func(ctx context.Context, txn *spanner.ReadWriteTransaction) error {
		stmt := spanner.Statement{SQL: sql, Params: params}
		rowCount, err := txn.Update(ctx, stmt)
		if err != nil {
			return fmt.Errorf("failed to exec query, err: %w", err)
		}
		if rowCount < 1 {
			return fmt.Errorf("got unexpected result, rowCount: %d", rowCount)
		}
		return nil
	}
	_, err := s.client.ReadWriteTransaction(ctx, cb)
	if err != nil {
		return fmt.Errorf("failed to perform ReadWriteTransaction, err: %w", err)
	}

	return nil
}

// CreateUser {@inheritdoc}.
func (s *SpannerStorage) CreateUser(ctx context.Context, input *payload.UserEntity) error {
	params := map[string]interface{}{
		"ID":        input.ID,
		"Email":     input.Email,
		"FirstName": input.FirstName,
		"LastName":  input.LastName,
		"Password":  input.Password,
		"CreatedAt": input.CreatedAt,
	}
	err := s.update(ctx, QueryInsertUser, params)
	if err != nil {
		return fmt.Errorf("failed to insert update, err: %w", err)
	}

	return nil
}

// GetUserByEmail {@inheritdoc}.
func (s *SpannerStorage) GetUserByEmail(ctx context.Context, email string) (*payload.UserEntity, error) {
	data := make([]*payload.UserEntity, 0)

	params := map[string]interface{}{"Email": email}
	stmt := spanner.Statement{SQL: QuerySelectUserByEmail, Params: params}
	itr := s.client.Single().Query(ctx, stmt)
	defer itr.Stop()
	for {
		row, err := itr.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, fmt.Errorf("failed to get next row, err: %w", err)
		}

		r := &payload.UserEntity{}
		err = row.Columns(
			&r.ID, &r.Email, &r.FirstName, &r.LastName, &r.Password, &r.CreatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to parse row, err: %w", err)
		}

		data = append(data, r)
	}

	return data[0], nil
}

// GetQuiz {@inheritdoc}.
func (s *SpannerStorage) GetQuiz(ctx context.Context, id string) (*payload.QuizEntity, error) {
	data, err := s.ListQuizzes(ctx, &payload.ListQuizzesRequest{ID: id})
	if err != nil {
		return nil, fmt.Errorf("failed to get quiz, err: %w", err)
	}
	if len(data) == 0 {
		return nil, fmt.Errorf("quiz not found")
	}

	return data[0], nil
}

// DeleteQuiz {@inheritdoc}.
func (s *SpannerStorage) DeleteQuiz(ctx context.Context, input *payload.DeleteQuizRequest) error {
	params := map[string]interface{}{
		"ID":        input.QuizID,
		"Deleted":   true,
		"DeletedAt": input.DeletedAt,
	}
	err := s.update(ctx, QueryMarkQuizAsDeleted, params)
	if err != nil {
		return fmt.Errorf("failed to perform update, err: %w", err)
	}

	return err
}

// ListQuizzes {@inheritdoc}.
func (s *SpannerStorage) ListQuizzes(
	ctx context.Context, input *payload.ListQuizzesRequest,
) ([]*payload.QuizEntity, error) {
	result := make([]*payload.QuizEntity, 0)

	where := ""
	if input.ID != "" {
		where += " AND ID = @ID"
	}
	if input.Author != "" {
		where += " AND Author = @Author"
	}

	params := map[string]interface{}{
		"ID":     input.ID,
		"Author": input.Author,
	}
	stmt := spanner.Statement{SQL: QuerySelectQuizzes + where, Params: params}
	itr := s.client.Single().Query(ctx, stmt)
	defer itr.Stop()
	for {
		row, err := itr.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, fmt.Errorf("failed to get next row, err: %w", err)
		}

		q := &payload.QuizEntity{}
		rawQuestions := spanner.NullJSON{}
		err = row.Columns(&q.ID, &rawQuestions, &q.Author, &q.Published, &q.UpdatedAt, &q.Deleted, &q.DeletedAt)
		if err != nil {
			return nil, fmt.Errorf("failed to parse row, err: %w", err)
		}

		questions := make([]*payload.QuizQuestionEntity, 0)
		err = mapstructure.Decode(rawQuestions.Value, &questions)
		if err != nil {
			return nil, fmt.Errorf("failed to decode questions, err: %w", err)
		}
		q.Questions = questions

		result = append(result, q)
	}

	return result, nil
}

// UpsertQuiz {@inheritdoc}.
func (s *SpannerStorage) UpsertQuiz(ctx context.Context, input *payload.QuizEntity) error {
	cols := []string{"ID", "Questions", "Author", "Published", "UpdatedAt", "Deleted", "DeletedAt"}
	vals := []interface{}{
		input.ID,
		spanner.NullJSON{Value: input.Questions, Valid: true},
		input.Author,
		input.Published,
		input.UpdatedAt,
		false,
		0,
	}
	m := []*spanner.Mutation{spanner.InsertOrUpdate(TableQuizzes, cols, vals)}
	_, err := s.client.Apply(ctx, m)
	if err != nil {
		return fmt.Errorf("failed to upsert quiz, err: %w", err)
	}

	return nil
}

// CreateSubmission {@inheritdoc}.
func (s *SpannerStorage) CreateSubmission(ctx context.Context, input *payload.SubmissionEntity) error {
	params := map[string]interface{}{
		"ID":          input.ID,
		"UserID":      input.UserID,
		"QuizID":      input.QuizID,
		"QuizAuthor":  input.QuizAuthor,
		"Score":       spanner.NullJSON{Value: input.Score, Valid: true},
		"SubmittedAt": input.SubmittedAt,
	}
	err := s.update(ctx, QueryInsertSubmission, params)
	if err != nil {
		if strings.Contains(err.Error(), "duplicate key") {
			return fmt.Errorf("submission already exists")
		}
		return fmt.Errorf("failed to insert submission, err: %w", err)
	}

	return nil
}

// ListSubmissions {@inheritdoc}.
func (s *SpannerStorage) ListSubmissions(
	ctx context.Context, input *payload.ListSubmissionsRequest,
) ([]*payload.SubmissionEntity, error) {
	result := make([]*payload.SubmissionEntity, 0)

	where := ""
	if input.UserID != "" {
		where += " AND UserID = @UserID"
	}
	if input.QuizAuthor != "" {
		where += " AND QuizAuthor = @QuizAuthor"
	}

	params := map[string]interface{}{
		"UserID":     input.UserID,
		"QuizAuthor": input.QuizAuthor,
	}
	stmt := spanner.Statement{SQL: QuerySelectSubmissions + where, Params: params}
	itr := s.client.Single().Query(ctx, stmt)
	defer itr.Stop()
	for {
		row, err := itr.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, fmt.Errorf("failed to get next row, err: %w", err)
		}

		s := &payload.SubmissionEntity{}
		rawScore := spanner.NullJSON{}
		err = row.Columns(&s.ID, &s.UserID, &s.QuizID, &s.QuizAuthor, &rawScore, &s.SubmittedAt)
		if err != nil {
			return nil, fmt.Errorf("failed to parse row, err: %w", err)
		}

		score := &payload.SubmissionScoreEntity{}
		err = mapstructure.Decode(rawScore.Value, &score)
		if err != nil {
			return nil, fmt.Errorf("failed to decode score, err: %w", err)
		}
		s.Score = score

		result = append(result, s)
	}

	return result, nil
}
