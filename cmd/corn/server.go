package corn

import (
	"fmt"
	"github.com/spf13/cobra"
)

var StartCmd = &cobra.Command{
	Use:   "fish",
	Short: "fish的第一次尝试融合了",
	Long:  `fish的第一次尝试融合了`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("我是崔可鱼")
	},
}
