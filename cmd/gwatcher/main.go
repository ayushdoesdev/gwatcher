package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/ayushdoesdev/gwatcher/internal/watcher"
)

var (
	cmdFlag   string
	dirFlag   string
	delayFlag int
)

func main() {
	rootCmd := &cobra.Command{
		Use:   "gwatcher",
		Short: "gwatcher 🦫 - watch .go files and auto-run commands",
		Run: func(cmd *cobra.Command, args []string) {
			if cmdFlag == "" {
				cmdFlag = "go test ./..."
			}
			if dirFlag == "" {
				dirFlag = "."
			}
			fmt.Printf("👀 Watching %s | Running: %s\n", dirFlag, cmdFlag)
			if err := watcher.Start(dirFlag, cmdFlag, delayFlag); err != nil {
				fmt.Println("❌ Error:", err)
				os.Exit(1)
			}
		},
	}

	rootCmd.Flags().StringVarP(&cmdFlag, "cmd", "c", "", "Command to run (default: go test ./...)")
	rootCmd.Flags().StringVarP(&dirFlag, "dir", "d", ".", "Directory to watch")
	rootCmd.Flags().IntVarP(&delayFlag, "delay", "t", 500, "Delay (ms) before re-running command")

	if err := rootCmd.Execute(); err != nil {
		fmt.Println("❌ Error:", err)
		os.Exit(1)
	}
}
