package cmd

import (
	"fmt"
	"strings"

	"github.com/ZiplEix/utilitaire/gitign"
	gitignParams "github.com/ZiplEix/utilitaire/gitign/params"
	"github.com/spf13/cobra"
)

// gitignCmd represents the gitign command
var gitignCmd = &cobra.Command{
	Use:   "gitign",
	Short: "Generate a gitignore file base on project files and settings",
	Long: `Generate a .gitignore file for your project based on detected or specified languages,
with options to ignore folders, append to an existing file, or optimize the generated rules.`,
	Run: func(cmd *cobra.Command, args []string) {
		if gitignParams.Params.Optimize && len(args) <= 1 {
			fmt.Println("Optimizing .gitignore")
			gitign.OptimizeGitignore()
			return
		}

		if len(args) == 0 {
			gitign.DetectLanguages(gitignParams.Params)
		} else {
			fmt.Println("Generating gitignore")
			gitign.GenerateGitignoreFromExtensions(args, gitignParams.Params)
		}
	},
}

func init() {
	rootCmd.AddCommand(gitignCmd)

	var ignore string
	gitignCmd.Flags().StringVarP(&ignore, "ignore", "i", "", "Comma-separated list of things to ignore. Example: --ignore node_modules,build,.go")
	gitignCmd.Flags().BoolVarP(&gitignParams.Params.Append, "append", "a", false, "Append the generated rules to an existing .gitignore file")
	gitignCmd.Flags().BoolVarP(&gitignParams.Params.Optimize, "optimize", "o", false, "Optimize the generated rules by removing duplicates and unnecessary entries")

	cobra.OnInitialize(func() {
		if ignore != "" {
			gitignParams.Params.Ignore = strings.Split(ignore, ",")
		}
	})
}
