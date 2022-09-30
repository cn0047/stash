package quizengine

import (
	"context"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"

	"quizengine/app/payload"
)

const (
	// TokenExpirationThreshold contains TTL for token.
	TokenExpirationThreshold = time.Hour * 1
)

// CreateUser creates user.
func (s *Service) CreateUser(ctx context.Context, input *payload.CreateUserRequest) error {
	// @TODO: Implement additional password validation.

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("failed to hash password, err: %v", err)
	}

	user := &payload.UserEntity{
		ID:        getUUID(),
		Email:     input.Email,
		FirstName: input.FirstName,
		LastName:  input.LastName,
		Password:  string(hashedPassword),
		CreatedAt: time.Now().Unix(),
	}

	err = s.storage.CreateUser(ctx, user)
	if err != nil {
		return fmt.Errorf("failed to create user, err: %v", err)
	}

	return nil
}

// AuthenticateUser performs user authentication and returns JWT token.
func (s *Service) AuthenticateUser(
	ctx context.Context, input *payload.AuthenticateUserRequest,
) (res payload.AuthenticateUserResponse, err error) {
	user, err := s.storage.GetUserByEmail(ctx, input.Email)
	if err != nil {
		return res, fmt.Errorf("failed to get user, err: %v", err)
	}

	// Compare password with hash.
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password))
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return res, fmt.Errorf("invalid authentication credentials")
	}

	claims := &payload.AuthenticationToken{
		UserID:    user.ID,
		Email:     user.Email,
		ExpiresAt: time.Now().Add(TokenExpirationThreshold).Unix(),
	}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), claims)
	res.Token, err = token.SignedString([]byte("secret"))
	if err != nil {
		return res, fmt.Errorf("failed to create signed token, err: %v", err)
	}
	res.UserID = user.ID

	return res, nil
}
