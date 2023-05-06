package register

import (
	v1user "github.com/linzhengen/mii-go/protobuf/go/user/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	healthpb "google.golang.org/grpc/health/grpc_health_v1"
)

func New(
	userHandler v1user.UserServiceServer,
) (s *grpc.Server) {
	s = grpc.NewServer()
	healthServer := health.NewServer()
	v1user.RegisterUserServiceServer(s, userHandler)
	healthpb.RegisterHealthServer(s, healthServer)
	return s
}
