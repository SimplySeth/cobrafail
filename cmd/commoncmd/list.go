package commoncmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// ListCmd represents the list command
var ListCmd = &cobra.Command{
	Use:   "list",
	Short: "List all requests or list all tasks",
	Long:  `Get a list of all tasks or a get a list of all requests that belongs to our group.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("list called")
	},
}

var (
	AccessListCmd    = *ListCmd
	ApprovalsListCmd = *ListCmd
)

func init() {}
