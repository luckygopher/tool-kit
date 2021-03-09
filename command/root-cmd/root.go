// @Description:
// @Author: Arvin
// @Date: 2021/3/9 11:18 下午
package root_cmd

import (
	convertStructCmd "github.com/qingyunjun/tool-kit/command/convert-struct-cmd"
	"github.com/spf13/cobra"
)

// 创建一个新的空命令
var root = &cobra.Command{}

// 初始化方法用来注册子命令
func init() {
	root.AddCommand(convertStructCmd.Cmd)
}

// 命令执行
func Execute() error {
	return root.Execute()
}
