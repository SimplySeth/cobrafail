package commoncmd

import (
	"github.com/spf13/cobra"
)

// RejectCmd represents the reject command
var RejectCmd = &cobra.Command{
	Use:   "reject",
	Short: "Reject a request",
	Long:  `Reject a request by providing the request number and a reason for rejection.`,
	Run: func(cmd *cobra.Command, args []string) {
		ticket, _ := cmd.Flags().GetString("ticket")
		if ticket == "" {
			missingTicket()
			cmd.Help()
		}
	},
}

var ApprovalsRejectCmd = *RejectCmd

func init() {
}
