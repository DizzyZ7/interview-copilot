package commands

import (
	"fmt"

	"interview-copilot/cli/internal/client"

	"github.com/spf13/cobra"
)

var quizCmd = &cobra.Command{
	Use:   "quiz",
	Short: "Start quiz mode",
	RunE: func(cmd *cobra.Command, args []string) error {
		c, _ := client.New()
		var qs []struct {
			ID    int    `json:"id"`
			Title string `json:"title"`
			Body  string `json:"body"`
		}
		if err := c.Do("POST", "/api/quiz/start", nil, &qs); err != nil {
			return err
		}

		for _, q := range qs {
			fmt.Printf("\n# %s\n%s\n", q.Title, q.Body)
			fmt.Print("Correct? (y/n): ")
			var ans string
			fmt.Scanln(&ans)
			c.Do("POST", "/api/quiz/answer", map[string]bool{"correct": ans == "y"}, nil)
		}
		fmt.Println("Quiz finished")
		return nil
	},
}
