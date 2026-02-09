package commands

import (
	"fmt"

	"interview-copilot/cli/internal/client"

	"github.com/spf13/cobra"
)

var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Login",
	RunE: func(cmd *cobra.Command, args []string) error {
		var email, password string
		fmt.Print("Email: ")
		fmt.Scanln(&email)
		fmt.Print("Password: ")
		fmt.Scanln(&password)

		c, _ := client.New()
		var res struct {
			Token string `json:"token"`
		}
		if err := c.Do("POST", "/auth/login", map[string]string{
			"email":    email,
			"password": password,
		}, &res); err != nil {
			return err
		}
		c.SetToken(res.Token)
		fmt.Println("Logged in")
		return nil
	},
}
