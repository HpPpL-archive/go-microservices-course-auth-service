package auth

import (
	"github.com/HpPpL/microservices_course_auth/internal/service"
	desc "github.com/HpPpL/microservices_course_auth/pkg/auth_v1"
)

type Implementation struct {
	desc.UnimplementedAuthV1Server
	authService service.AuthService
}

func NewImplementation(authService service.AuthService) *Implementation {
	return &Implementation{
		authService: authService,
	}
}
