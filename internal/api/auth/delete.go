package auth

import (
	"context"
	desc "github.com/HpPpL/microservices_course_auth/pkg/auth_v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (i *Implementation) Delete(ctx context.Context, req *desc.DeleteRequest) (*emptypb.Empty, error) {
	_, err := i.authService.Delete(ctx, req.Id)

	if err != nil {
		return &emptypb.Empty{}, err
	}
	return &emptypb.Empty{}, nil
}
