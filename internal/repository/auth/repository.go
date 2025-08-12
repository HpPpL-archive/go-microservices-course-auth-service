package auth

import (
	"context"
	"errors"
	desc "github.com/HpPpL/microservices_course_auth/pkg/auth_v1"
	"github.com/jackc/pgx/v4/pgxpool"
	"google.golang.org/protobuf/types/known/emptypb"
	"log"
	"time"

	sq "github.com/Masterminds/squirrel"

	"github.com/HpPpL/microservices_course_auth/internal/model"
	"github.com/HpPpL/microservices_course_auth/internal/repository"
	"github.com/HpPpL/microservices_course_auth/internal/repository/auth/converter"
	modelRepo "github.com/HpPpL/microservices_course_auth/internal/repository/auth/model"
)

const (
	tableName = "auth"

	idColumn        = "id"
	nameColumn      = "name"
	emailColumn     = "email"
	roleColumn      = "role"
	createdAtColumn = "created_at"
	updatedAtColumn = "updated_at"
)

var (
	// General PG errors
	errFailedBuildQuery = errors.New("failed to build query")
	errUserDoesntExist  = errors.New("user with current id doesn't exist")

	// Create errors
	errPasswordDoesntMatch = errors.New("password doesn't match")
	errInvalidRole         = errors.New("invalid role value")
	errFailedInsertUser    = errors.New("failed to insert user")

	// Get errors
	errFailedSelectUser = errors.New("failed to select user")

	// Update errors
	errFailedUpdateUser = errors.New("failed to update user data")

	// Delete errors
	errFailedDeleteUser = errors.New("failed to delete user")
)

const (
	roleUnspecified = "unspecified"
	roleUser        = "user"
	roleAdmin       = "admin"
)

type repo struct {
	db *pgxpool.Pool
}

func NewRepository(db *pgxpool.Pool) repository.AuthRepository {
	return &repo{db: db}
}

func (r *repo) Create(ctx context.Context, info *model.UserDataInfo) (int64, error) {
	log.Print("There is create request!")
	if info.Password != info.PasswordConfirm {
		log.Print("Passwords do not match!")
		return 0, errPasswordDoesntMatch
	}

	var roleStr string
	switch info.Role {
	case desc.Role_ROLE_UNSPECIFIED:
		roleStr = roleUnspecified
	case desc.Role_ROLE_USER:
		roleStr = roleUser
	case desc.Role_ROLE_ADMIN:
		roleStr = roleAdmin
	default:
		log.Printf("invalid role value: %v", info.Role)
		return 0, errInvalidRole
	}

	// Вынести имя таблицы можно попробовать потом
	builderInsert := sq.Insert(tableName).
		PlaceholderFormat(sq.Dollar).
		Columns(nameColumn, emailColumn, roleColumn).
		Values(info.Name, info.Email, roleStr).
		Suffix("RETURNING id")

	query, args, err := builderInsert.ToSql()
	if err != nil {
		log.Printf("failed to build query: %v", err)
		return 0, errFailedBuildQuery
	}

	var UserID int64
	err = r.db.QueryRow(ctx, query, args...).Scan(&UserID)
	if err != nil {
		log.Printf("failed to insert user: %v", err)
		return 0, errFailedInsertUser
	}

	log.Printf("inserted user with id: %v", UserID)
	return UserID, nil
}

func (r *repo) Get(ctx context.Context, userId int64) (*model.User, error) {
	log.Print("There is get request!")

	builderSelectOne := sq.Select(idColumn, nameColumn, emailColumn, roleColumn, createdAtColumn, updatedAtColumn).
		From(tableName).
		PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{idColumn: userId}).
		Limit(1)

	query, args, err := builderSelectOne.ToSql()
	if err != nil {
		log.Printf("failed to build query: %v", err)
		return &model.User{}, errFailedBuildQuery
	}

	var user modelRepo.User
	var roleStr string

	err = r.db.QueryRow(ctx, query, args...).Scan(&user.ID, &user.Name, &user.Email, &roleStr, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		log.Printf("failed to select user: %v", err)
		return &model.User{}, errFailedSelectUser
	}

	switch roleStr {
	case roleUnspecified:
		user.Role = desc.Role_ROLE_UNSPECIFIED
	case roleUser:
		user.Role = desc.Role_ROLE_USER
	case roleAdmin:
		user.Role = desc.Role_ROLE_ADMIN
	default:
		log.Print(errInvalidRole.Error())
		return &model.User{}, errInvalidRole
	}

	log.Printf("ID: %v, Name: %v, Email: %v, Role: %v, CreatedAt: %v, UpdatedAt: %v",
		user.ID, user.Name, user.Email, user.Role, user.CreatedAt, user.UpdatedAt)

	return converter.ToAuthFromRepo(&user), nil
}

func (r *repo) Update(ctx context.Context, updateInfo *model.UpdateUser) (*emptypb.Empty, error) {
	log.Print("There is update request!")

	userID := updateInfo.ID

	builderUpdate := sq.Update(tableName).
		PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{idColumn: userID})

	nameWrapper := updateInfo.Name
	if nameWrapper != nil {
		builderUpdate = builderUpdate.Set(nameColumn, nameWrapper.GetValue())
	}

	emailWrapper := updateInfo.Email
	if emailWrapper != nil {
		builderUpdate = builderUpdate.Set(emailColumn, emailWrapper.GetValue())
	}

	builderUpdate = builderUpdate.Set(updatedAtColumn, time.Now())

	query, args, err := builderUpdate.ToSql()
	if err != nil {
		log.Printf("failed to build query: %v", err)
		return &emptypb.Empty{}, errFailedBuildQuery
	}

	res, err := r.db.Exec(ctx, query, args...)
	if err != nil {
		log.Printf("failed to update user data: %v", err)
		return &emptypb.Empty{}, errFailedUpdateUser
	}

	if res.RowsAffected() == 0 {
		return &emptypb.Empty{}, errUserDoesntExist
	}

	log.Printf("updated %d rows", res.RowsAffected())

	return &emptypb.Empty{}, nil
}

func (r *repo) Delete(ctx context.Context, userID int64) (*emptypb.Empty, error) {
	builderDelete := sq.Delete(tableName).
		PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{idColumn: userID})

	query, args, err := builderDelete.ToSql()
	if err != nil {
		log.Printf("failed to build query: %v", err)
		return &emptypb.Empty{}, errFailedBuildQuery
	}

	res, err := r.db.Exec(ctx, query, args...)
	if err != nil {
		log.Printf("failed to delete user: %v", err)
		return &emptypb.Empty{}, errFailedDeleteUser
	}

	if res.RowsAffected() == 0 {
		return &emptypb.Empty{}, errUserDoesntExist
	}

	log.Printf("Deleted %d rows", res.RowsAffected())
	return &emptypb.Empty{}, nil
}
