package cmd

import (
	"os"

	"github.com/ZiplEix/utilitaire/concat"
	"github.com/spf13/cobra"
)

// concatCmd represents the concat command
var concatCmd = &cobra.Command{
	Use:   "concat",
	Short: "Concat all the files corresponding the the pattern on one output file",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		globPattern := args
		err := concat.Concat(globPattern)
		if err != nil {
			cmd.PrintErrf("Error: %v\n", err)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(concatCmd)

	concatCmd.Flags().StringVarP(&concat.OutputFile, "output", "o", "concat.txt", "Output file name")
}
