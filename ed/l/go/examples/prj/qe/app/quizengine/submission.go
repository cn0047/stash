package quizengine

import (
	"context"
	"fmt"
	"time"

	"quizengine/app/payload"
)

// ListSubmissions gets submissions.
func (s *Service) ListSubmissions(
	ctx context.Context, input *payload.ListSubmissionsRequest,
) ([]*payload.SubmissionEntity, error) {
	// @TODO: Implement user_id and quiz_author validation.

	if input.UserID == "" && input.QuizAuthor == "" {
		return nil, fmt.Errorf("either user_id or quiz_author must be provided")
	}
	if input.UserID != "" && input.Requester != input.UserID {
		return nil, fmt.Errorf("can review only own submissions")
	}
	if input.QuizAuthor != "" && input.Requester != input.QuizAuthor {
		return nil, fmt.Errorf("can review only submissions to own quizzes")
	}

	res, err := s.storage.ListSubmissions(ctx, input)
	if err != nil {
		return nil, fmt.Errorf("failed to list submissions, err: %v", err)
	}

	return res, err
}

// CreateSubmission creates submission.
func (s *Service) CreateSubmission(
	ctx context.Context, input *payload.CreateSubmissionRequest,
) (*payload.CreateSubmissionResponse, error) {
	submission := &payload.SubmissionEntity{
		ID:          getUUID(),
		UserID:      input.UserID,
		QuizID:      input.QuizID,
		SubmittedAt: time.Now().Unix(),
	}

	// Check correctness.
	quiz, err := s.storage.GetQuiz(ctx, input.QuizID)
	if err != nil {
		return nil, fmt.Errorf("failed to get quiz, err: %v", err)
	}
	submission.QuizAuthor = quiz.Author
	if !quiz.Published {
		return nil, fmt.Errorf("can't submit unpublished quiz")
	}
	if quiz.Deleted {
		return nil, fmt.Errorf("can't submit deleted quiz")
	}
	if quiz.Author == input.UserID {
		return nil, fmt.Errorf("can't submit own quiz")
	}

	submission.Score, err = calcScore(input, quiz)
	if err != nil {
		return nil, fmt.Errorf("failed to create submission, err: %v", err)
	}

	// Save submission.
	err = s.storage.CreateSubmission(ctx, submission)
	if err != nil {
		return nil, fmt.Errorf("failed to create submission, err: %v", err)
	}

	res := &payload.CreateSubmissionResponse{
		TotalScore: submission.Score.TotalScore,
		Questions:  submission.Score.Questions,
	}

	return res, nil
}

func calcScore(
	input *payload.CreateSubmissionRequest, quiz *payload.QuizEntity,
) (*payload.SubmissionScoreEntity, error) {
	score := &payload.SubmissionScoreEntity{
		TotalScore: 0,
		Questions:  make([]float64, 0),
	}

	for questionIndex, answers := range input.Answers {
		if questionIndex > len(quiz.Questions) {
			return nil, fmt.Errorf("submitted more questions than quiz has")
		}
		var questionScore float64
		question := quiz.Questions[questionIndex]
		// @TODO: Check question.Type and accept only 1 answer for payload.QuizQuestionTypeSingleAnswer.
		for _, answerIndex := range answers {
			if answerIndex > len(question.Answers) {
				return nil, fmt.Errorf("submitted more answers than quiz questions have")
			}
			questionScore += question.Answers[answerIndex].Score
		}
		score.TotalScore += questionScore
		score.Questions = append(score.Questions, questionScore)
	}

	return score, nil
}
