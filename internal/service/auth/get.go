package auth

import (
	"context"
	"github.com/HpPpL/microservices_course_auth/internal/model"
)

func (s *serv) Get(ctx context.Context, id int64) (*model.User, error) {
	user, err := s.authRepository.Get(ctx, id)
	if err != nil {
		return &model.User{}, err
	}

	return user, nil
}
