package register

import (
	v1user "github.com/linzhengen/mii-go/protobuf/go/user/v1"
	"google.golang.org/grpc"
)

func New(
	userHandler v1user.UserServiceServer,
) (s *grpc.Server) {
	s = grpc.NewServer()
	v1user.RegisterUserServiceServer(s, userHandler)
	return s
}
