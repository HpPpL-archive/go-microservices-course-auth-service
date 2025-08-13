package converter

import (
	"github.com/HpPpL/microservices_course_auth/internal/model"
	modelRepo "github.com/HpPpL/microservices_course_auth/internal/repository/auth/model"
	desc "github.com/HpPpL/microservices_course_auth/pkg/auth_v1"
	"time"
)

const (
	roleUnspecified = "unspecified"
	roleUser        = "user"
	roleAdmin       = "admin"
)

func ToUserFromRepo(user *modelRepo.User) *model.User {
	var UpdatedAt time.Time
	if user.UpdatedAt.Valid {
		UpdatedAt = user.UpdatedAt.Time
	}

	return &model.User{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		Role:      user.Role,
		CreatedAt: user.CreatedAt,
		UpdatedAt: UpdatedAt,
	}
}

func ToUserInfoFromService(userInfo *model.UserDataInfo) *modelRepo.UserDataInfo {
	var roleStr string
	switch userInfo.Role {
	case desc.Role_ROLE_UNSPECIFIED:
		roleStr = roleUnspecified
	case desc.Role_ROLE_USER:
		roleStr = roleUser
	case desc.Role_ROLE_ADMIN:
		roleStr = roleAdmin
	}
	return &modelRepo.UserDataInfo{
		Name:            userInfo.Name,
		Email:           userInfo.Email,
		Password:        userInfo.Password,
		PasswordConfirm: userInfo.PasswordConfirm,
		Role:            roleStr,
	}
}

func ToUpdateUserFromService(userUpdate *model.UpdateUser) *modelRepo.UpdateUser {
	return &modelRepo.UpdateUser{
		ID:    userUpdate.ID,
		Name:  userUpdate.Name,
		Email: userUpdate.Email,
	}
}
