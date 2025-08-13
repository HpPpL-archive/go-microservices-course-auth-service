package auth

import (
	"github.com/HpPpL/microservices_course_auth/internal/repository"
	"github.com/HpPpL/microservices_course_auth/internal/service"
)

type serv struct {
	authRepository repository.AuthRepository
}

func NewService(authRepository repository.AuthRepository) service.AuthService {
	return &serv{authRepository: authRepository}
}
