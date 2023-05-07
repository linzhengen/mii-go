package cmd

import (
	"github.com/linzhengen/mii-go/internal/usecase"
	"github.com/linzhengen/mii-go/pkg/logger"
	"github.com/spf13/cobra"
)

func NewUserCmd(userUseCase usecase.UserUseCase) *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "user",
		Short: "manage user",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.HelpFunc()(cmd, args)
		},
	}
	c := &userCmd{userUseCase: userUseCase}
	cmd.AddCommand(c.GetUser())
	return cmd
}

type UserCmd interface {
	GetUser() *cobra.Command
}

type userCmd struct {
	userUseCase usecase.UserUseCase
}

func (u userCmd) GetUser() *cobra.Command {
	var userID string
	var getUserCmd = &cobra.Command{
		Use: "get",
		Run: func(cmd *cobra.Command, args []string) {
			if user, err := u.userUseCase.GetUser(cmd.Context(), userID); err != nil {
				logger.Errorf("get user error: %v", err)
			} else {
				cmd.Println(user)
			}
		},
	}
	getUserCmd.Flags().StringVar(&userID, "userId", "", "The ID of the user")
	return getUserCmd
}
