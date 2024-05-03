package ttype

import (
	"fmt"

	"github.com/spf13/cobra"

	"ihsnow/cmd/commoncmd"
)

// ApprovalsCmd represents the approvals command
var ApprovalsCmd = &cobra.Command{
	Use:   "approvals",
	Short: "Manage approval requests by granting or revoking access",
	Long:  `For managing approval tickets by granting or revoking.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("approvals called")
		fmt.Println(cmd.Parent().Name())
	},
}

var (
	GcpApprovalsCmd = *ApprovalsCmd
	SshApprovalsCmd = *ApprovalsCmd
)

func init() {
	ApprovalsCmd.AddCommand(&commoncmd.ApprovalsApproveCmd)
	ApprovalsCmd.AddCommand(&commoncmd.ApprovalsGetCmd)
	ApprovalsCmd.AddCommand(&commoncmd.ApprovalsListCmd)
	ApprovalsCmd.AddCommand(&commoncmd.ApprovalsRejectCmd)
}
