package repository

import (
	"context"
	"github.com/HpPpL/microservices_course_auth/internal/model"
	modelRepo "github.com/HpPpL/microservices_course_auth/internal/repository/auth/model"
	"google.golang.org/protobuf/types/known/emptypb"
)

type AuthRepository interface {
	Create(ctx context.Context, info *modelRepo.UserDataInfo) (int64, error)
	Get(ctx context.Context, id int64) (*model.User, error)
	Update(ctx context.Context, updateInfo *modelRepo.UpdateUser) (*emptypb.Empty, error)
	Delete(ctx context.Context, id int64) (*emptypb.Empty, error)
}
