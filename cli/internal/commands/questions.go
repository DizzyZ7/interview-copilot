package commands

import (
	"fmt"

	"interview-copilot/cli/internal/client"

	"github.com/spf13/cobra"
)

var questionsCmd = &cobra.Command{
	Use:   "questions",
	Short: "List questions",
	RunE: func(cmd *cobra.Command, args []string) error {
		c, _ := client.New()
		var res []struct {
			ID    int      `json:"id"`
			Title string   `json:"title"`
			Tags  []string `json:"tags"`
		}
		if err := c.Do("GET", "/api/questions", nil, &res); err != nil {
			return err
		}
		for _, q := range res {
			fmt.Printf("%d. %s %v\n", q.ID, q.Title, q.Tags)
		}
		return nil
	},
}
