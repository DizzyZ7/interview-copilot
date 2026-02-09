package commands

import (
	"fmt"

	"interview-copilot/cli/internal/client"

	"github.com/spf13/cobra"
)

var registerCmd = &cobra.Command{
	Use:   "register",
	Short: "Register",
	RunE: func(cmd *cobra.Command, args []string) error {
		var email, password string
		fmt.Print("Email: ")
		fmt.Scanln(&email)
		fmt.Print("Password: ")
		fmt.Scanln(&password)

		c, _ := client.New()
		if err := c.Do("POST", "/auth/register", map[string]string{
			"email":    email,
			"password": password,
		}, nil); err != nil {
			return err
		}
		fmt.Println("Registered, now login")
		return nil
	},
}
