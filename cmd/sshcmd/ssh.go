package sshcmd

import (
	"ihsnow/cmd/ttype"

	"github.com/spf13/cobra"
)

// SshCmd represents the ssh command
var SshCmd = &cobra.Command{
	Use:   "ssh",
	Short: "Manage SSH related requests or tasks",
	Long:  `This command is for the handling of SSH related requests or tasks`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cmd.Help()
		}
	},
	ValidArgs: []string{"access", "approvals"},
}

func init() {
	SshCmd.AddCommand(&ttype.SshAccessCmd)
	SshCmd.AddCommand(&ttype.SshApprovalsCmd)
}
