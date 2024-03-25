package initcmd

import (
	"log"

	"github.com/spf13/cobra"
)

func CreateUserFlags(createUserCmd *cobra.Command) {
	createUserCmd.Flags().StringP("username", "u", "", "User name")
	err := createUserCmd.MarkFlagRequired("username")
	if err != nil {
		log.Fatalf("failed to mark username flag as required %v", err.Error())
	}
	createUserCmd.Flags().StringP("passwrd", "p", "", "Password")
	err = createUserCmd.MarkFlagRequired("passwrd")
	if err != nil {
		log.Fatalf("failed to mark password flag as required %v", err.Error())
	}
	createUserCmd.Flags().StringP("email", "e", "", "Email")
	err = createUserCmd.MarkFlagRequired("email")
	if err != nil {
		log.Fatalf("failed to mark email flag as required %v", err.Error())
	}
}
