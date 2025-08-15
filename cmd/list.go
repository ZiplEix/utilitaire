package cmd

import (
	"fmt"
	"time"

	"github.com/ZiplEix/utilitaire/tmp"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Liste les suppressions programmées",
	RunE: func(cmd *cobra.Command, args []string) error {
		st, err := tmp.List()
		if err != nil {
			return err
		}
		for _, r := range st {
			switch r.Scheduler {
			case tmp.SchedSystemd:
				fmt.Printf("%s  | systemd unit=%s  | expires=%s\n", r.Path, r.Unit, r.Expiration.Format(time.RFC3339))
			case tmp.SchedAt:
				fmt.Printf("%s  | at job=%d      | expires=%s\n", r.Path, r.AtJob, r.Expiration.Format(time.RFC3339))
			}
		}
		return nil
	},
}

func init() {
	tmpCmd.AddCommand(listCmd)
}
