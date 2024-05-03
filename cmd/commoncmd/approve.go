package commoncmd

import (
	"github.com/spf13/cobra"
)

// approveCmd represents the approve command
var ApproveCmd = &cobra.Command{
	Use:   "approve",
	Short: "Approve a request or task",
	Long:  `Approving a request indicates that you have reviewed the request and wish to promote it to the next stage of the workflow. This command allows you to approve a request or task.`,
	Run: func(cmd *cobra.Command, args []string) {
		ticket, _ := cmd.Flags().GetString("ticket")
		if ticket == "" {
			missingTicket()
			cmd.Help()
		}
	},
}

var ApprovalsApproveCmd = *ApproveCmd

func init() {
}
