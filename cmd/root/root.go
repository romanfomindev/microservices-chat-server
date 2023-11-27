package root

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "app",
	Short: "chat application",
}

var authCmd = &cobra.Command{
	Use:   "auth",
	Short: "Auth commands",
}

var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Login",
	Run: func(cmd *cobra.Command, args []string) {
		//
	},
}

var accessTokenCmd = &cobra.Command{
	Use:   "access-token",
	Short: "Get Access Token",
	Run: func(cmd *cobra.Command, args []string) {
		//
	},
}

var refreshTokenCmd = &cobra.Command{
	Use:   "refresh-token",
	Short: "Get Refresh Token",
	Run: func(cmd *cobra.Command, args []string) {
		//
	},
}

var chatCmd = &cobra.Command{
	Use:   "chat",
	Short: "Chat commands",
}

var createChatCmd = &cobra.Command{
	Use:   "create",
	Short: "Create chat",
	Run: func(cmd *cobra.Command, args []string) {
		//
	},
}

var deleteChatCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete chat",
	Run: func(cmd *cobra.Command, args []string) {
		//
	},
}

var sendMessageCmd = &cobra.Command{
	Use:   "send",
	Short: "Send message",
	Run: func(cmd *cobra.Command, args []string) {
		//
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(authCmd)
	authCmd.AddCommand(loginCmd)
	authCmd.AddCommand(accessTokenCmd)
	authCmd.AddCommand(refreshTokenCmd)
	rootCmd.AddCommand(chatCmd)
	chatCmd.AddCommand(createChatCmd)
	chatCmd.AddCommand(deleteChatCmd)
	chatCmd.AddCommand(sendMessageCmd)

}
