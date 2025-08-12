package converter

import (
	"github.com/HpPpL/microservices_course_auth/internal/model"
	modelRepo "github.com/HpPpL/microservices_course_auth/internal/repository/auth/model"
)

// Нам же вообще не важен этот метод, так как мы его не возврашаем
func ToAuthFromRepo(user *modelRepo.User) *model.User {
	//var UpdatedAt *timestamppb.Timestamp
	//if user.UpdatedAt.Valid {
	//	UpdatedAt = timestamppb.New(user.UpdatedAt.Time)
	//}

	return &model.User{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		Role:      user.Role,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		//CreatedAt: timestamppb.New(user.CreatedAt),
		//UpdatedAt: UpdatedAt,
	}
}
