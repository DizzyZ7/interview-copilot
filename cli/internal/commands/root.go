package commands

import "github.com/spf13/cobra"

var rootCmd = &cobra.Command{
	Use:   "copilot",
	Short: "Interview Copilot CLI",
}

func Execute() {
	rootCmd.Execute()
}

func init() {
	rootCmd.AddCommand(loginCmd)
	rootCmd.AddCommand(registerCmd)
	rootCmd.AddCommand(questionsCmd)
	rootCmd.AddCommand(randomCmd)
	rootCmd.AddCommand(quizCmd)
	rootCmd.AddCommand(progressCmd)
	rootCmd.AddCommand(whoamiCmd)
}
