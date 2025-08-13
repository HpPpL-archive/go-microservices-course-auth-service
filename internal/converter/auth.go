package converter

import (
	"github.com/HpPpL/microservices_course_auth/internal/model"
	desc "github.com/HpPpL/microservices_course_auth/pkg/auth_v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func ToUserInfoFromDesc(userInfo *desc.UserDataInfo) *model.UserDataInfo {
	return &model.UserDataInfo{
		Name:            userInfo.Name,
		Email:           userInfo.Email,
		Password:        userInfo.Password,
		PasswordConfirm: userInfo.PasswordConfirm,
		Role:            userInfo.Role,
	}
}

func ToUserFromService(user *model.User) *desc.GetResponse {
	return &desc.GetResponse{
		Id:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		Role:      user.Role,
		CreatedAt: timestamppb.New(user.CreatedAt),
		UpdatedAt: timestamppb.New(user.UpdatedAt),
	}
}

func ToUpdateUserFromDesc(request *desc.UpdateRequest) *model.UpdateUser {
	var name string
	nameWrapper := request.GetName()
	if nameWrapper != nil {
		name = nameWrapper.GetValue()
	}

	var email string
	emailWrapper := request.GetEmail()
	if emailWrapper != nil {
		email = emailWrapper.GetValue()
	}

	return &model.UpdateUser{
		ID:    request.Id,
		Name:  name,
		Email: email,
	}
}
