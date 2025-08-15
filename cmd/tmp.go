package cmd

import (
	"fmt"
	"os"

	"github.com/ZiplEix/utilitaire/tmp"
	"github.com/spf13/cobra"
)

// tmpCmd represents the tmp command
var tmpCmd = &cobra.Command{
	Use:   "tmp [directoryPath] [expiration]",
	Short: "A command to create a directory with an expiration",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		directoryPath := ""
		expiration := "1d"

		if len(args) == 1 {
			directoryPath = args[0]
		} else if len(args) == 2 {
			directoryPath = args[0]
			expiration = args[1]
		}

		err := tmp.TmpDir(directoryPath, expiration)
		if err != nil {
			fmt.Println("Error creating temporary directory:", err)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(tmpCmd)

	tmpCmd.Flags().BoolVarP(&tmp.Verbose, "verbose", "v", false, "Enable verbose output")
}
