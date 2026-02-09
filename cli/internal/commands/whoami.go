package commands

import (
	"fmt"

	"interview-copilot/cli/internal/client"

	"github.com/spf13/cobra"
)

var whoamiCmd = &cobra.Command{
	Use:   "whoami",
	Short: "Show current user id",
	RunE: func(cmd *cobra.Command, args []string) error {
		c, _ := client.New()
		var res struct {
			UserID int `json:"user_id"`
		}
		if err := c.Do("GET", "/api/me", nil, &res); err != nil {
			return err
		}
		fmt.Println("User ID:", res.UserID)
		return nil
	},
}
