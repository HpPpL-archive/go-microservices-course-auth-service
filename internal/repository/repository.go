package repository

import (
	"context"
	"github.com/HpPpL/microservices_course_auth/internal/model"
	desc "github.com/HpPpL/microservices_course_auth/pkg/auth_v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

type AuthRepository interface {
	Create(ctx context.Context, info *desc.UserDataInfo) (int64, error)
	Get(ctx context.Context, id int64) (*model.User, error)
	Update(ctx context.Context, updateInfo *desc.UpdateRequest) (*emptypb.Empty, error)
	Delete(ctx context.Context, id int64) (*emptypb.Empty, error)
}
