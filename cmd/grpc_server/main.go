package main

import (
	"context"
	authAPI "github.com/HpPpL/microservices_course_auth/internal/api/auth"
	repoAuth "github.com/HpPpL/microservices_course_auth/internal/repository/auth"
	serviceAuth "github.com/HpPpL/microservices_course_auth/internal/service/auth"

	"flag"
	"github.com/jackc/pgx/v4/pgxpool"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"

	"github.com/HpPpL/microservices_course_auth/internal/config"
	desc "github.com/HpPpL/microservices_course_auth/pkg/auth_v1"
)

// Path to config
var configPath string

func init() {
	flag.StringVar(&configPath, "config-path", ".env", "path to config file")
}

func main() {
	flag.Parse()
	ctx := context.Background()

	err := config.Load(configPath)
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	grpcConfig, err := config.NewGRPCConfig()
	if err != nil {
		log.Fatalf("failed to get grpc config")
	}

	pgConfig, err := config.NewPGConfig()
	if err != nil {
		log.Fatalf("failed to get pg config")
	}

	lis, err := net.Listen("tcp", grpcConfig.Address())
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	pool, err := pgxpool.Connect(ctx, pgConfig.DSN())
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	defer pool.Close()

	authRepository := repoAuth.NewRepository(pool)
	authService := serviceAuth.NewService(authRepository)

	s := grpc.NewServer()
	reflection.Register(s)
	desc.RegisterAuthV1Server(s, authAPI.NewImplementation(authService))

	log.Printf("server listening at %v", lis.Addr())

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
