package converter

import (
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/HpPpL/microservices_course_auth/internal/model"
	desc "github.com/HpPpL/microservices_course_auth/pkg/auth_v1"
)

func ToUserFromService(user *model.User) *desc.UserDataInfo
