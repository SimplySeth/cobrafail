package commoncmd

import (
	"github.com/spf13/cobra"
)

// AssignCmd represents the assign command
var AssignCmd = &cobra.Command{
	Use:   "assign",
	Short: "Assigns the ticket or request",
	Long:  `Assigning a ticket gives you credit for the work done. This also assures the requester that it is being worked on.`,
	Run: func(cmd *cobra.Command, args []string) {
		ticket, _ := cmd.Flags().GetString("ticket")
		if ticket == "" {
			missingTicket()
			cmd.Help()
		}
	},
}

var (
	AccessAssignCmd    = *AssignCmd
	ApprovalsAssignCmd = *AssignCmd
)

func init() {}
