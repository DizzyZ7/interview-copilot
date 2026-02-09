package commands

import (
	"fmt"

	"interview-copilot/cli/internal/client"

	"github.com/spf13/cobra"
)

var randomCmd = &cobra.Command{
	Use:   "random",
	Short: "Get random questions",
	RunE: func(cmd *cobra.Command, args []string) error {
		c, _ := client.New()
		var res []struct {
			ID    int    `json:"id"`
			Title string `json:"title"`
			Body  string `json:"body"`
		}
		if err := c.Do("GET", "/api/random?limit=3", nil, &res); err != nil {
			return err
		}
		for _, q := range res {
			fmt.Printf("\n# %s\n%s\n", q.Title, q.Body)
		}
		return nil
	},
}
