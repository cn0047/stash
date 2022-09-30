package payload

// QuizQuestionType represents dedicated datatype for quiz question type.
type QuizQuestionType int

const (
	// QuizQuestionTypeSingleAnswer means question with 1 correct answer.
	QuizQuestionTypeSingleAnswer QuizQuestionType = 1
	// QuizQuestionTypeMultipleAnswers means question with more than 1 correct answer.
	QuizQuestionTypeMultipleAnswers QuizQuestionType = 2
)

// QuizEntity represents internal struct with protected data about quiz, which can't be publicly exposed.
type QuizEntity struct {
	ID        string                `json:"id" mapstructure:"ID"`
	Questions []*QuizQuestionEntity `json:"questions" mapstructure:"Questions"`
	Author    string                `json:"author" mapstructure:"Author"`
	Published bool                  `json:"published" mapstructure:"Published"`
	UpdatedAt int64                 `json:"updated_at" mapstructure:"UpdatedAt"`
	Deleted   bool                  `json:"deleted" mapstructure:"Deleted"`
	DeletedAt int64                 `json:"deleted_at" mapstructure:"DeletedAt"`
}

// QuizQuestionEntity holds data about question for internal quiz struct.
type QuizQuestionEntity struct {
	Title   string                      `json:"title" mapstructure:"Title"`
	Type    QuizQuestionType            `json:"type" mapstructure:"Type"`
	Answers []*QuizQuestionAnswerEntity `json:"answers" mapstructure:"Answers"`
}

// QuizQuestionAnswerEntity holds data about answer for internal quiz struct.
type QuizQuestionAnswerEntity struct {
	Title         string  `json:"title" mapstructure:"Title"`
	CorrectAnswer bool    `json:"correct_answer" mapstructure:"CorrectAnswer"`
	Score         float64 `json:"score" mapstructure:"Score"`
}

// ToDTO converts internal quiz entity to public quiz.
func (q *QuizEntity) ToDTO() *Quiz {
	quizDTO := &Quiz{
		ID:        q.ID,
		Questions: make([]*QuizQuestion, 0, len(q.Questions)),
		Author:    q.Author,
		Published: q.Published,
		Deleted:   q.Deleted,
	}
	for _, question := range q.Questions {
		questionDTO := &QuizQuestion{
			Title:   question.Title,
			Type:    question.Type,
			Answers: make([]*QuizQuestionAnswer, 0, len(question.Answers)),
		}
		for _, answer := range question.Answers {
			answerDTO := &QuizQuestionAnswer{Title: answer.Title}
			questionDTO.Answers = append(questionDTO.Answers, answerDTO)
		}
		quizDTO.Questions = append(quizDTO.Questions, questionDTO)
	}

	return quizDTO
}

// SubmissionEntity represents internal struct about submission.
type SubmissionEntity struct {
	ID          string                 `json:"id" mapstructure:"ID"`
	UserID      string                 `json:"user_id" mapstructure:"UserID"`
	QuizID      string                 `json:"quiz_id" mapstructure:"QuizID"`
	QuizAuthor  string                 `json:"quiz_author" mapstructure:"QuizAuthor"`
	Score       *SubmissionScoreEntity `json:"score" mapstructure:"Score"`
	SubmittedAt int64                  `json:"submitted_at" mapstructure:"SubmittedAt"`
}

// SubmissionScoreEntity holds data about score for internal submission struct.
type SubmissionScoreEntity struct {
	TotalScore float64   `json:"total_score" mapstructure:"TotalScore"`
	Questions  []float64 `json:"questions" mapstructure:"Questions"`
}

// UserEntity represents internal struct about user.
type UserEntity struct {
	ID        string `json:"id" mapstructure:"ID"`
	Email     string `json:"email" mapstructure:"Email"`
	FirstName string `json:"first_name" mapstructure:"FirstName"`
	LastName  string `json:"last_name" mapstructure:"LastName"`
	Password  string `json:"password" mapstructure:"Password"`
	CreatedAt int64  `json:"created_at" mapstructure:"CreatedAt"`
}
