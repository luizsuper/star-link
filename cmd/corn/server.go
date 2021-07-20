package corn

import (
	"fmt"
	"github.com/spf13/cobra"
)

var StartCmd = &cobra.Command{
	Use:   "fish",
	Short: "fish sss",
	Long:  `fish`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("fish sss")
	},
}
