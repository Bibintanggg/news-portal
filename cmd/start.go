package cmd

import (
	"bwa-news/internal/app"

	"github.com/spf13/cobra"
)

var startCmd = &cobra.Command{ // kalau command `go run main.go start` yang akan running dari sini
	Use:   "start",
	Short: "start",
	Long:  `start`,
	Run: func(cmd *cobra.Command, args []string) {
		app.RunServer()
	},
} // sampai sini - akan ngerunning RunServer

func init() {
	rootCmd.AddCommand(startCmd)
}
