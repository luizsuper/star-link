package api

import (
	"fmt"
	"github.com/spf13/cobra"
)

var StartCmd = &cobra.Command{
	Use:   "server",
	Short: "Git is a distributed version control system.",
	Long: `Git is a free and open source distributed version control system
designed to handle everything from small to very large projects 
with speed and efficiency.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("hello world")
	},
}
