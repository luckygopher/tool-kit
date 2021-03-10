// @Description:
// @Author: Arvin
// @Date: 2021/3/9 11:18 下午
package root

import (
	dbCmd "github.com/qingyunjun/tool-kit/cmd/database"
	"github.com/spf13/cobra"
)

// 创建一个新的空命令作为根
var rootCmd = &cobra.Command{}

// 初始化方法用来注册子命令
func init() {
	rootCmd.AddCommand(dbCmd.DbCmd)
}

// 命令执行
func Execute() error {
	return rootCmd.Execute()
}
