package cmd

import (
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(runCommand)
}

var runCommand = &cobra.Command{
	Use:   "run",
	Short: "运行服务",
	Long:  `数据库模型重新生成，与models定义的数据模型同步`,
	Run: func(cmd *cobra.Command, args []string) {
	
		run()
	},
}

func run() {


}
