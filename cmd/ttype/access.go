package ttype

import (
	"fmt"

	"github.com/spf13/cobra"

	"ihsnow/cmd/commoncmd"
)

// accessCmd represents the access command
var AccessCmd = &cobra.Command{
	Use:   "access",
	Short: "Manage tickets for adding to access groups",
	Long:  `For managing tickets that add users to access groups for granting of access privileges.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("access called")
		fmt.Println(cmd.Parent().Name())
		// client := snow.Connect()
		// fmt.Println(client)
	},
	ValidArgs: []string{"annotate", "assign", "get", "close", "list"},
}

var (
	GcpAccessCmd = *AccessCmd
	SshAccessCmd = *AccessCmd
)

func init() {
	AccessCmd.AddCommand(&commoncmd.AccessAnnotateCmd)
	AccessCmd.AddCommand(&commoncmd.AccessAssignCmd)
	AccessCmd.AddCommand(&commoncmd.AccessGetCmd)
	AccessCmd.AddCommand(&commoncmd.AccessCloseCmd)
	AccessCmd.AddCommand(&commoncmd.AccessListCmd)
}
