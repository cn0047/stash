package quizengine

import (
	"testing"

	"quizengine/app/payload"

	"github.com/stretchr/testify/assert"
)

func TestService_initQuiz(tc *testing.T) {
	tc.Run("2 answers", func(t *testing.T) {
		input := &payload.QuizEntity{
			Questions: []*payload.QuizQuestionEntity{
				{
					Title: "Is Earth flat?",
					Type:  payload.QuizQuestionTypeSingleAnswer,
					Answers: []*payload.QuizQuestionAnswerEntity{
						{Title: "yes", CorrectAnswer: false},
						{Title: "no", CorrectAnswer: true},
					},
				},
			},
		}

		initQuiz(input)

		assert.Equal(t, input.Questions[0].Answers[0].Score, float64(-1))
		assert.Equal(t, input.Questions[0].Answers[1].Score, float64(+1))
	})

	tc.Run("3 wrong answers", func(t *testing.T) {
		input := &payload.QuizEntity{
			Questions: []*payload.QuizQuestionEntity{
				{
					Title: "Is Earth flat?",
					Type:  payload.QuizQuestionTypeSingleAnswer,
					Answers: []*payload.QuizQuestionAnswerEntity{
						{Title: "yes", CorrectAnswer: false},
						{Title: "sure", CorrectAnswer: false},
						{Title: "absolutely yes", CorrectAnswer: false},
						{Title: "no", CorrectAnswer: true},
					},
				},
			},
		}

		initQuiz(input)

		assert.Equal(t, input.Questions[0].Answers[0].Score, float64(-0.333))
		assert.Equal(t, input.Questions[0].Answers[1].Score, float64(-0.333))
		assert.Equal(t, input.Questions[0].Answers[2].Score, float64(-0.333))
		assert.Equal(t, input.Questions[0].Answers[3].Score, float64(+1))
	})

	tc.Run("2 correct answers", func(t *testing.T) {
		input := &payload.QuizEntity{
			Questions: []*payload.QuizQuestionEntity{
				{
					Title: "Is Earth flat?",
					Type:  payload.QuizQuestionTypeMultipleAnswers,
					Answers: []*payload.QuizQuestionAnswerEntity{
						{Title: "yes", CorrectAnswer: false},
						{Title: "absolutely yes", CorrectAnswer: false},
						{Title: "of course", CorrectAnswer: false},
						{Title: "of course, yes", CorrectAnswer: false},
						{Title: "no", CorrectAnswer: true},
						{Title: "absolutely no", CorrectAnswer: true},
					},
				},
			},
		}

		initQuiz(input)

		assert.Equal(t, input.Questions[0].Answers[0].Score, float64(-0.25))
		assert.Equal(t, input.Questions[0].Answers[1].Score, float64(-0.25))
		assert.Equal(t, input.Questions[0].Answers[2].Score, float64(-0.25))
		assert.Equal(t, input.Questions[0].Answers[3].Score, float64(-0.25))
		assert.Equal(t, input.Questions[0].Answers[4].Score, float64(+0.5))
		assert.Equal(t, input.Questions[0].Answers[5].Score, float64(+0.5))
	})
}
