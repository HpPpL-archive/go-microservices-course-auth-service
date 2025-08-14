package auth

import (
	"context"
	"github.com/HpPpL/microservices_course_auth/internal/converter"
	desc "github.com/HpPpL/microservices_course_auth/pkg/auth_v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (i *Implementation) Update(ctx context.Context, req *desc.UpdateRequest) (*emptypb.Empty, error) {
	_, err := i.authService.Update(ctx, converter.ToUpdateUserFromDesc(req))
	if err != nil {
		return &emptypb.Empty{}, err
	}

	return &emptypb.Empty{}, nil
}
