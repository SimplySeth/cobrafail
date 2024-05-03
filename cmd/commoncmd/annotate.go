package commoncmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// annotateCmd represents the annotate command
var AnnotateCmd = &cobra.Command{
	Use:   "annotate",
	Short: "Annotate (add a comment) to a request or task",
	Long:  `Comments are useful for tracking the history and progress of a request or task. This commmand allows you to add a comment to a request or task.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("annotate command called")
		fmt.Println(cmd.Parent().Name())
	},
}

var (
	AccessAnnotateCmd    = *AnnotateCmd
	ApprovalsAnnotateCmd = *AnnotateCmd
)
