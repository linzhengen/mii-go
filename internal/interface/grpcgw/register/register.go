package register

import (
	"context"

	v1user "github.com/linzhengen/mii-go/protobuf/go/user/v1"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/linzhengen/mii-go/config"
	"github.com/linzhengen/mii-go/pkg/logger"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func New(
	envCfg *config.EnvConfig,
) *runtime.ServeMux {
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}

	ctx := context.Background()
	mustRegisterGWHandler(v1user.RegisterUserServiceHandlerFromEndpoint, ctx, mux, envCfg.Grpc.Addr(), opts)

	return mux
}

type registerFunc func(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error

func mustRegisterGWHandler(register registerFunc, ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) {
	err := register(ctx, mux, endpoint, opts)
	if err != nil {
		logger.Severe(err)
	}
}
