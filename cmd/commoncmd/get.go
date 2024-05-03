package commoncmd

import (
	"github.com/spf13/cobra"
)

// GetCmd represents the get command
var GetCmd = &cobra.Command{
	Use:   "get",
	Short: "Get allows you to view details about a request or task.",
	Long:  `Get allows you to view details about a request or task, which allows you to determine what to do next.`,
	Run: func(cmd *cobra.Command, args []string) {
		ticket, _ := cmd.Flags().GetString("ticket")
		if ticket == "" {
			missingTicket()
			cmd.Help()
		}
	},
}

var (
	AccessGetCmd    = *GetCmd
	ApprovalsGetCmd = *GetCmd
)

func init() {
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// GetCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// GetCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
