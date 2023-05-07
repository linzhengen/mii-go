package register

import (
	"github.com/linzhengen/mii-go/internal/interface/cmd/handler"
	"github.com/spf13/cobra"
)

type Commands []*cobra.Command

func New(
	userHandler handler.UserHandler,
) Commands {
	return []*cobra.Command{
		userHandler.Command(),
	}
}
