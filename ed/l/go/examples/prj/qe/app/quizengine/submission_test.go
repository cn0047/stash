package quizengine

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"quizengine/app/payload"
)

func TestService_calcScore(tc *testing.T) {
	quiz := &payload.QuizEntity{
		Questions: []*payload.QuizQuestionEntity{
			{
				Title: "Is Earth flat?",
				Type:  payload.QuizQuestionTypeSingleAnswer,
				Answers: []*payload.QuizQuestionAnswerEntity{
					{Title: "yes", Score: float64(-0.5)},
					{Title: "yes", Score: float64(-0.5)},
					{Title: "no", Score: float64(+1)},
				},
			},
			{
				Title: "Is Earth flat??",
				Type:  payload.QuizQuestionTypeMultipleAnswers,
				Answers: []*payload.QuizQuestionAnswerEntity{
					{Title: "yes", Score: float64(-0.25)},
					{Title: "yes", Score: float64(-0.25)},
					{Title: "yes", Score: float64(-0.25)},
					{Title: "yes", Score: float64(-0.25)},
					{Title: "no", Score: float64(+0.5)},
					{Title: "no", Score: float64(+0.5)},
				},
			},
		},
	}

	tc.Run("2 wrong answers", func(t *testing.T) {
		submission := &payload.CreateSubmissionRequest{Answers: [][]int{{0}, {0}}}
		score, err := calcScore(submission, quiz)
		assert.Equal(t, err, nil)
		assert.Equal(t, score.TotalScore, float64(-0.75))
	})

	tc.Run("4 wrong answers", func(t *testing.T) {
		submission := &payload.CreateSubmissionRequest{Answers: [][]int{{0, 1}, {0, 1}}}
		score, err := calcScore(submission, quiz)
		assert.Equal(t, err, nil)
		assert.Equal(t, score.TotalScore, float64(-1.5))
	})

	tc.Run("max score", func(t *testing.T) {
		submission := &payload.CreateSubmissionRequest{Answers: [][]int{{2}, {4, 5}}}
		score, err := calcScore(submission, quiz)
		assert.Equal(t, err, nil)
		assert.Equal(t, score.TotalScore, float64(+2))
	})

	tc.Run("min score", func(t *testing.T) {
		submission := &payload.CreateSubmissionRequest{Answers: [][]int{{0, 1}, {0, 1, 2, 3}}}
		score, err := calcScore(submission, quiz)
		assert.Equal(t, err, nil)
		assert.Equal(t, score.TotalScore, float64(-2))
	})

	tc.Run("zero score", func(t *testing.T) {
		submission := &payload.CreateSubmissionRequest{Answers: [][]int{{0, 1}, {4, 5}}}
		score, err := calcScore(submission, quiz)
		assert.Equal(t, err, nil)
		assert.Equal(t, score.TotalScore, float64(0))
	})
}
