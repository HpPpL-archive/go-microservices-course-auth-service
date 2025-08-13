package auth

import (
	"context"
	"github.com/HpPpL/microservices_course_auth/internal/model"
	repoConverter "github.com/HpPpL/microservices_course_auth/internal/repository/auth/converter"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *serv) Update(ctx context.Context, req *model.UpdateUser) (*emptypb.Empty, error) {
	_, err := s.authRepository.Update(ctx, repoConverter.ToUpdateUserFromService(req))
	if err != nil {
		return &emptypb.Empty{}, err
	}

	return &emptypb.Empty{}, nil
}
