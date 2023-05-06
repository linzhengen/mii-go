package register

import (
	"github.com/linzhengen/mii-go/internal/interface/cmd/handler"
	"github.com/spf13/cobra"
)

func New(
	rootCmd *cobra.Command,
	handler handler.UserHandler,
) {
	rootCmd.AddCommand(handler.Command())
}
