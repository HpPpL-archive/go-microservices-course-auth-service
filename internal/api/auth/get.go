package auth

import (
	"context"
	"github.com/HpPpL/microservices_course_auth/internal/converter"
	desc "github.com/HpPpL/microservices_course_auth/pkg/auth_v1"
)

func (i *Implementation) Get(ctx context.Context, req *desc.GetRequest) (*desc.GetResponse, error) {
	user, err := i.authService.Get(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	return converter.ToUserFromService(user), nil
}
