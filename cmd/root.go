package cmd

import (
	"os"

	"github.com/urfave/cli/v2"
)

// 创建一个新的空命令作为根
var rootCmd = cli.NewApp()

// Execute 命令执行
func Execute() error {
	rootCmd.Usage = "tool 使用方式"
	return rootCmd.Run(os.Args)
}
