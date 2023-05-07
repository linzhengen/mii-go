package register

import (
	"context"
	"fmt"
	"net/http"

	"google.golang.org/grpc/connectivity"

	v1user "github.com/linzhengen/mii-go/protobuf/go/user/v1"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/linzhengen/mii-go/config"
	"github.com/linzhengen/mii-go/pkg/logger"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func New(
	envCfg config.EnvConfig,
) *runtime.ServeMux {
	ctx := context.Background()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}

	mux := runtime.NewServeMux()
	if err := mux.HandlePath(http.MethodGet, "/healthz", healthzServer(ctx, envCfg.Grpc.Addr(), opts)); err != nil {
		logger.Severe(err)
	}

	mustRegisterGWHandler(v1user.RegisterUserServiceHandlerFromEndpoint, ctx, mux, envCfg.Grpc.Addr(), opts)

	return mux
}

func healthzServer(ctx context.Context, endpoint string, opts []grpc.DialOption) runtime.HandlerFunc {
	conn, err := grpc.DialContext(ctx, endpoint, opts...)
	if err != nil {
		logger.Severe(err)
	}
	return func(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
		w.Header().Set("Content-Type", "text/plain")
		if s := conn.GetState(); s != connectivity.Ready {
			http.Error(w, fmt.Sprintf("grpc server is %s", s), http.StatusBadGateway)
			return
		}
		fmt.Fprintln(w, "ok")
	}
}

type registerFunc func(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error

func mustRegisterGWHandler(register registerFunc, ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) {
	err := register(ctx, mux, endpoint, opts)
	if err != nil {
		logger.Severe(err)
	}
}
