package cmd

import (
	"fmt"
	"os"

	"github.com/avayayu/quant_data/internal/utils"
	"github.com/spf13/cobra"
)

var (
	configFilePath string
	release        bool
)

var rootCmd = &cobra.Command{
	Use:   "BRIS",
	Short: "cmd for BRIS.",
	Long: `BRIS CMD 命令行，主要包括以下功能:
	1.migrate 数据库迁移 ,
    2.run 运行服务
	`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	//	Run: func(cmd *cobra.Command, args []string) { },
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	// cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	utils.EnsureDir("~/.bris")

	rootCmd.PersistentFlags().BoolVar(&release, "release", false, "后端工作环境，调试环境或者生产环境")
	rootCmd.PersistentFlags().StringVar(&configFilePath, "config", "~/.bris/config.yaml", "配置文件目录")
	// Cobra also supports local flags, which will only run
	// when this action is called directly.

}
