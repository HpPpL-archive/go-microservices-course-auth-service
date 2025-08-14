package auth

import (
	"context"
	"github.com/HpPpL/microservices_course_auth/internal/converter"
	desc "github.com/HpPpL/microservices_course_auth/pkg/auth_v1"
	"log"
)

func (i *Implementation) Create(ctx context.Context, req *desc.CreateRequest) (*desc.CreateResponse, error) {
	userID, err := i.authService.Create(ctx, converter.ToUserInfoFromDesc(req.Info))
	if err != nil {
		return nil, err
	}

	log.Printf("inserted user with id: %v", userID)
	return &desc.CreateResponse{
		Id: userID,
	}, nil
}
