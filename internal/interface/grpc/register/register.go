package register

import (
	"github.com/linzhengen/mii-go/protobuf/go/v1/user"
	"google.golang.org/grpc"
)

func New(
	userHandler user.UserServiceServer,
) (s *grpc.Server) {
	s = grpc.NewServer()
	user.RegisterUserServiceServer(s, userHandler)
	return s
}
