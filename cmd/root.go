package cmd

import (
	"fmt"
	"os"
	"spacer/formatter"

	"github.com/spf13/cobra"
)

var (
	path string
)

func init() {
	rootCmd.PersistentFlags().StringVarP(&path, "file", "f", "./README.md", "指定需要格式化的文件或目录路径。")
	// todo 指定配置文件、输出目录
}

var rootCmd = &cobra.Command{
	Use:   "spacer",
	Short: "格式化 markdown 文件。",
	Long:  `spacer 用于为你格式化 markdown 文件，spacer 非常灵活，每一条格式化规则都是可插拔的，并且您可以根据自己的情况，轻松对其进行扩展。`,
	Run: func(cmd *cobra.Command, args []string) {
		//_ = cmd.Usage()
		formatter.Start(path)
	},
	// todo 子命令，规则，查看最终规则列表，全部规则列表，规则说明
	// todo 子命令，处理生成锚点
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
