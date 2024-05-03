package commoncmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// CloseCmd represents the close command
var CloseCmd = &cobra.Command{
	Use:   "close",
	Short: "Closes a request or task",
	Long:  `Closing means marking a task or request as done. This indicates to the requester, that they may proceed to use the service(s) or feature(s) requested.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("close called")
		ticket, _ := cmd.Flags().GetString("ticket")
		if ticket == "" {
			missingTicket()
			cmd.Help()
		}
	},
}

var (
	AccessCloseCmd    = *CloseCmd
	ApprovalsCloseCmd = *CloseCmd
)

func init() {}
