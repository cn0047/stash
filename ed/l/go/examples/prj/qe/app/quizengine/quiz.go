package quizengine

import (
	"context"
	"fmt"
	"math"
	"time"

	"quizengine/app/payload"
)

const (
	MaxQuizQuestionsCountThreshold = 10
	MaxQuizAnswersCountThreshold   = 5
)

// ListQuizzes gets quizzes.
func (s *Service) ListQuizzes(
	ctx context.Context, input *payload.ListQuizzesRequest,
) ([]*payload.Quiz, error) {
	// @TODO: Implement author validation.

	quizzes, err := s.storage.ListQuizzes(ctx, input)
	if err != nil {
		return nil, fmt.Errorf("failed to Upsert quiz, err: %v", err)
	}

	result := make([]*payload.Quiz, 0, len(quizzes))
	for _, q := range quizzes {
		result = append(result, q.ToDTO())
	}

	return result, nil
}

// GetQuiz gets quiz by ID.
func (s *Service) GetQuiz(ctx context.Context, ID string) (*payload.Quiz, error) {
	quiz, err := s.storage.GetQuiz(ctx, ID)
	if err != nil {
		return nil, fmt.Errorf("failed to upsert quiz, err: %v", err)
	}

	return quiz.ToDTO(), nil
}

// DeleteQuiz deletes quiz.
func (s *Service) DeleteQuiz(ctx context.Context, input *payload.DeleteQuizRequest) error {
	// Check correctness.
	quiz, err := s.storage.GetQuiz(ctx, input.QuizID)
	if err != nil {
		return fmt.Errorf("failed to find quiz")
	}
	if quiz.Author != input.Requester {
		return fmt.Errorf("upsert forbidden, can delete only own quiz")
	}

	input.DeletedAt = time.Now().Unix()

	err = s.storage.DeleteQuiz(ctx, input)
	if err != nil {
		return fmt.Errorf("failed to delete quiz, err: %v", err)
	}

	return nil
}

// UpsertQuiz performs quiz upsert (insert/update).
func (s *Service) UpsertQuiz(ctx context.Context, input *payload.UpsertQuizRequest) error {
	// Check correctness.
	if input.ID != "" {
		quiz, err := s.storage.GetQuiz(ctx, input.ID)
		if err == nil { // Works only for existing quiz.
			if quiz.Published {
				return fmt.Errorf("published quiz can't be updated")
			}
			if quiz.Author != input.Author {
				return fmt.Errorf("upsert forbidden, can update only own quiz")
			}
		}
	}

	// Additional validations.
	if len(input.Questions) == 0 {
		return fmt.Errorf("at least 1 question must be specified")
	}
	if len(input.Questions) > MaxQuizQuestionsCountThreshold {
		return fmt.Errorf("questions count must be less than %v", MaxQuizQuestionsCountThreshold)
	}

	// Prepare quiz.
	if input.ID == "" {
		input.ID = getUUID()
	}
	quiz := &payload.QuizEntity{
		ID:        input.ID,
		Questions: make([]*payload.QuizQuestionEntity, 0, len(input.Questions)),
		Author:    input.Author,
		Published: input.Published,
		UpdatedAt: time.Now().Unix(),
	}
	for _, question := range input.Questions {
		if len(question.Answers) == 0 {
			return fmt.Errorf("at least 1 answer must be specified")
		}
		if len(question.Answers) > MaxQuizAnswersCountThreshold {
			return fmt.Errorf("answers count must be less than %v", MaxQuizAnswersCountThreshold)
		}
		answers := make([]*payload.QuizQuestionAnswerEntity, 0, len(question.Answers))
		for _, answer := range question.Answers {
			a := &payload.QuizQuestionAnswerEntity{
				Title:         answer.Title,
				CorrectAnswer: answer.CorrectAnswer,
			}
			answers = append(answers, a)
		}
		q := &payload.QuizQuestionEntity{
			Title:   question.Title,
			Answers: answers,
		}
		quiz.Questions = append(quiz.Questions, q)
	}
	initQuiz(quiz)

	// Upsert quiz.
	err := s.storage.UpsertQuiz(ctx, quiz)
	if err != nil {
		return fmt.Errorf("failed to Upsert quiz, err: %v", err)
	}

	return nil
}

func initQuiz(input *payload.QuizEntity) {
	for questionIndex, question := range input.Questions {
		// Calculate counts.
		totalAnswersCount := len(question.Answers)
		totalCorrectAnswersCount := 0
		for _, answer := range question.Answers {
			if answer.CorrectAnswer {
				totalCorrectAnswersCount++
			}
		}
		totalWrongAnswersCount := totalAnswersCount - totalCorrectAnswersCount

		// Calculate scores.
		correctAnswerScore := +round(1 / float64(totalCorrectAnswersCount))
		wrongAnswerScore := -round(1 / float64(totalWrongAnswersCount))

		// Assign scores.
		for answerIndex, answer := range question.Answers {
			score := wrongAnswerScore
			if answer.CorrectAnswer {
				score = correctAnswerScore
			}
			input.Questions[questionIndex].Answers[answerIndex].Score = score
		}

		// Set question type.
		question.Type = payload.QuizQuestionTypeMultipleAnswers
		if totalCorrectAnswersCount == 1 {
			question.Type = payload.QuizQuestionTypeSingleAnswer
		}
	}
}

func round(x float64) float64 {
	return math.Round(x*1000) / 1000
}
