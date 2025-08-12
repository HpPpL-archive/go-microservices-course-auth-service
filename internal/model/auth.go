package model

import (
	"database/sql"
	desc "github.com/HpPpL/microservices_course_auth/pkg/auth_v1"
	"google.golang.org/protobuf/types/known/wrapperspb"
	"time"
)

type UserDataInfo struct {
	Name            string
	Email           string
	Password        string
	PasswordConfirm string
	Role            desc.Role
}

type User struct {
	ID        int64
	Name      string
	Email     string
	Role      desc.Role
	CreatedAt time.Time
	UpdatedAt sql.NullTime
}

type UpdateUser struct {
	ID    int64
	Name  *wrapperspb.StringValue
	Email *wrapperspb.StringValue
}
