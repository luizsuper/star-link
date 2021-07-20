var StartCmd = &cobra.Command{
	Use:   "fish",
	Short: "fish的第一次尝试",
	Long: `fish的第一次尝试fish的第一次尝试fish的第一次尝试fish的第一次尝试`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("我是崔可鱼")
	},
}