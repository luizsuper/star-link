var StartCmd = &cobra.Command{
	Use:   "fish",
	Short: "fish�ĵ�һ�γ���",
	Long: `fish�ĵ�һ�γ���fish�ĵ�һ�γ���fish�ĵ�һ�γ���fish�ĵ�һ�γ���`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("���Ǵ޿���")
	},
}