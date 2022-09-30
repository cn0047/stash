package payload

import (
	"errors"
	"time"
)

// AuthenticationToken holds authentication data for JWT token.
type AuthenticationToken struct {
	UserID    string
	Email     string
	ExpiresAt int64
}

// Valid validates AuthenticationToken.
func (a *AuthenticationToken) Valid() error {
	if a.ExpiresAt < time.Now().Unix() {
		return errors.New("token expired")
	}

	return nil
}

// Quiz represents struct without protected data, it can be publicly exposed.
type Quiz struct {
	ID        string          `json:"id"`
	Questions []*QuizQuestion `json:"questions"`
	Author    string          `json:"author"`
	Published bool            `json:"published"`
	Deleted   bool            `json:"deleted"`
}

// QuizQuestion holds data about question for public quiz struct.
type QuizQuestion struct {
	Title   string                `json:"title"`
	Type    QuizQuestionType      `json:"type"`
	Answers []*QuizQuestionAnswer `json:"answers"`
}

// QuizQuestionAnswer holds data about answer for public quiz struct.
type QuizQuestionAnswer struct {
	Title string `json:"title"`
}
