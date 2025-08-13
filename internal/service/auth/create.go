package auth

import (
	"context"
	"github.com/HpPpL/microservices_course_auth/internal/model"
	"github.com/HpPpL/microservices_course_auth/internal/repository/auth/converter"
)

func (s *serv) Create(ctx context.Context, req *model.UserDataInfo) (int64, error) {
	userID, err := s.authRepository.Create(ctx, converter.ToUserInfoFromService(req))
	if err != nil {
		return 0, err
	}

	return userID, nil
}
