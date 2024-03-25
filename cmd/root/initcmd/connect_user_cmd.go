package initcmd

import (
	"log"

	"github.com/spf13/cobra"
)

func ConnectUserExistFlags(connectExistUserCmd *cobra.Command) {
	connectExistUserCmd.Flags().StringP("username", "u", "", "User name")
	err := connectExistUserCmd.MarkFlagRequired("username")
	if err != nil {
		log.Fatalf("failed to mark username flag as required %v", err.Error())
	}
	connectExistUserCmd.Flags().StringP("passwrd", "p", "", "Password")
	err = connectExistUserCmd.MarkFlagRequired("passwrd")
	if err != nil {
		log.Fatalf("failed to mark password flag as required %v", err.Error())
	}
}
