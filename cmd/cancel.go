package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/ZiplEix/utilitaire/tmp"
	"github.com/spf13/cobra"
)

// cancelCmd represents the cancel command
var cancelCmd = &cobra.Command{
	Use:   "cancel <directoryPath>",
	Short: "Cancel the deletion of a temporary directory",
	Args:  cobra.ExactArgs(1),
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		abs, err := filepath.Abs(args[0])
		if err != nil {
			fmt.Println("Invalid path:", err)
			os.Exit(1)
		}
		if err := tmp.Cancel(abs); err != nil {
			fmt.Println("Cancel failed:", err)
			os.Exit(1)
		}
		fmt.Println("Cancellation done for", abs)
	},
}

func init() {
	tmpCmd.AddCommand(cancelCmd)
}
