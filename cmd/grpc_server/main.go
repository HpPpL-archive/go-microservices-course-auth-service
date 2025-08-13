package main

import (
	"context"
	"github.com/HpPpL/microservices_course_auth/internal/converter"
	repoAuth "github.com/HpPpL/microservices_course_auth/internal/repository/auth"
	"github.com/HpPpL/microservices_course_auth/internal/service"
	serviceAuth "github.com/HpPpL/microservices_course_auth/internal/service/auth"

	"flag"
	"github.com/jackc/pgx/v4/pgxpool"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/types/known/emptypb"
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

// server implements the AuthV1 gRPC service by embedding UnimplementedAuthV1Server.
type server struct {
	desc.UnimplementedAuthV1Server
	authService service.AuthService
}

// Create part
func (s *server) Create(ctx context.Context, req *desc.CreateRequest) (*desc.CreateResponse, error) {
	userID, err := s.authService.Create(ctx, converter.ToUserInfoFromDesc(req.Info))
	if err != nil {
		return nil, err
	}

	log.Printf("inserted user with id: %v", userID)
	return &desc.CreateResponse{
		Id: userID,
	}, nil
}

// Get part
func (s *server) Get(ctx context.Context, req *desc.GetRequest) (*desc.GetResponse, error) {
	user, err := s.authService.Get(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	return converter.ToUserFromService(user), nil
}

// Update part
func (s *server) Update(ctx context.Context, req *desc.UpdateRequest) (*emptypb.Empty, error) {
	_, err := s.authService.Update(ctx, converter.ToUpdateUserFromDesc(req))
	if err != nil {
		return &emptypb.Empty{}, err
	}

	return &emptypb.Empty{}, nil
}

// Delete part
func (s *server) Delete(ctx context.Context, req *desc.DeleteRequest) (*emptypb.Empty, error) {
	_, err := s.authService.Delete(ctx, req.Id)

	if err != nil {
		return &emptypb.Empty{}, err
	}
	return &emptypb.Empty{}, nil
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

	s := grpc.NewServer()
	reflection.Register(s)
	authRepository := repoAuth.NewRepository(pool)
	desc.RegisterAuthV1Server(s, &server{authService: serviceAuth.NewService(authRepository)})
	log.Printf("server listening at %v", lis.Addr())

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
