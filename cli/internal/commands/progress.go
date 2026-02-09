package commands

import (
	"fmt"

	"interview-copilot/cli/internal/client"

	"github.com/spf13/cobra"
)

var progressCmd = &cobra.Command{
	Use:   "progress",
	Short: "Show progress stats",
	RunE: func(cmd *cobra.Command, args []string) error {
		c, _ := client.New()
		var res struct {
			Total   int     `json:"total"`
			Correct int     `json:"correct"`
			Ratio   float64 `json:"ratio"`
		}
		if err := c.Do("GET", "/api/progress", nil, &res); err != nil {
			return err
		}
		fmt.Printf("Answered: %d\nCorrect: %d\nAccuracy: %.1f%%\n",
			res.Total, res.Correct, res.Ratio*100)
		return nil
	},
}
