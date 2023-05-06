package handler

import (
	"github.com/linzhengen/mii-go/internal/usecase"
	"github.com/linzhengen/mii-go/pkg/logger"
	"github.com/spf13/cobra"
)

func NewUserHandler(userUseCase usecase.UserUseCase) UserHandler {
	var userCmd = &cobra.Command{
		Use: "user",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.HelpFunc()(cmd, args)
		},
	}
	return &userHandler{userCmd: userCmd, userUseCase: userUseCase}
}

type UserHandler interface {
	Command() *cobra.Command
	GetUser()
}

type userHandler struct {
	userCmd     *cobra.Command
	userUseCase usecase.UserUseCase
}

func (u userHandler) Command() *cobra.Command {
	return u.userCmd
}

func (u userHandler) GetUser() {
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
	u.userCmd.AddCommand(getUserCmd)
}
