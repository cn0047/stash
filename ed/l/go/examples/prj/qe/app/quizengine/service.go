package quizengine

import (
	"strings"

	"github.com/google/uuid"

	"quizengine/app/storage"
)

// Service represents main service in quizengine.
type Service struct {
	storage storage.Storage
}

// NewService returns new instance of quizengine service.
func NewService(s storage.Storage) (*Service, error) {
	svc := &Service{
		storage: s,
	}

	return svc, nil
}

// getUUID gets UUID, represents single function for service to generate UUID.
// IMPORTANT: It won't return real UUID,
// but only 1st chunk of it, it's done only for simplicity of this demo project,
// and it shouldn't generate issues because returned ID by this function is unique enough to avoid collisions.
func getUUID() string {
	return strings.Split(uuid.New().String(), "-")[0]
}
