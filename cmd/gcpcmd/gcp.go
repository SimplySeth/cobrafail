package gcpcmd

import (
	"ihsnow/cmd/ttype"

	"github.com/spf13/cobra"
)

// GcpCmd represents the gcp command
var GcpCmd = &cobra.Command{
	Use:   "gcp",
	Short: "Manage GCP related requests or tasks",
	Long:  `This command is for handling of GCP related  requests or tasks.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cmd.Help()
		}
	},
	ValidArgs: []string{"access", "approvals"},
}

func init() {
	GcpCmd.AddCommand(&ttype.GcpAccessCmd)
	GcpCmd.AddCommand(&ttype.GcpApprovalsCmd)
}
