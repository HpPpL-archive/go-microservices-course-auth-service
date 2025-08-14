package app

import (
	"context"
	"github.com/HpPpL/microservices_course_auth/internal/api/auth"
	"github.com/HpPpL/microservices_course_auth/internal/closer"
	"github.com/HpPpL/microservices_course_auth/internal/config"
	"github.com/HpPpL/microservices_course_auth/internal/repository"
	repoAuth "github.com/HpPpL/microservices_course_auth/internal/repository/auth"
	"github.com/HpPpL/microservices_course_auth/internal/service"
	serviceAuth "github.com/HpPpL/microservices_course_auth/internal/service/auth"
	"log"

	"github.com/jackc/pgx/v4/pgxpool"
)

type serviceProvider struct {
	pgConfig   config.PGConfig
	grpcConfig config.GRPCConfig

	pool           *pgxpool.Pool
	authRepository repository.AuthRepository

	authService service.AuthService

	authImpl *auth.Implementation
}

func newServiceProvider() *serviceProvider {
	return &serviceProvider{}
}

func (s *serviceProvider) PGConfig() config.PGConfig {
	if s.pgConfig == nil {
		cfg, err := config.NewPGConfig()
		if err != nil {
			log.Fatalf("failed to get pg config: %s", err.Error())
		}
		s.pgConfig = cfg
	}

	return s.pgConfig
}

func (s *serviceProvider) GRPCConfig() config.GRPCConfig {
	if s.grpcConfig == nil {
		cfg, err := config.NewGRPCConfig()
		if err != nil {
			log.Fatalf("failed to get grpc config: %s", err.Error())
		}
		s.grpcConfig = cfg
	}
	return s.grpcConfig
}

func (s *serviceProvider) Pool(ctx context.Context) *pgxpool.Pool {
	if s.pool == nil {
		pool, err := pgxpool.Connect(ctx, s.PGConfig().DSN())
		if err != nil {
			log.Fatalf("failed to connect to database: %v", err)
		}
		closer.Add(func() error {
			pool.Close()
			return nil
		})
		s.pool = pool
	}

	return s.pool
}

func (s *serviceProvider) AuthRepository(ctx context.Context) repository.AuthRepository {
	if s.authRepository == nil {
		authRepository := repoAuth.NewRepository(s.Pool(ctx))
		s.authRepository = authRepository
	}

	return s.authRepository
}

func (s *serviceProvider) AuthService(ctx context.Context) service.AuthService {
	if s.authService == nil {
		authService := serviceAuth.NewService(s.AuthRepository(ctx))
		s.authService = authService
	}

	return s.authService
}

func (s *serviceProvider) AuthImpl(ctx context.Context) *auth.Implementation {
	if s.authImpl == nil {
		authImpl := auth.NewImplementation(s.AuthService(ctx))
		s.authImpl = authImpl
	}

	return s.authImpl
}
