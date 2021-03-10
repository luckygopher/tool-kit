// @Description: 数据库操作根命令
// @Author: Arvin
// @Date: 2021/3/10 11:30 上午
package database

import "github.com/spf13/cobra"

var DbCmd = &cobra.Command{
	Use:   "db",
	Short: "数据库相关操作",
	Long:  "数据库相关操作",
	RunE: func(cmd *cobra.Command, args []string) error {
		// todo 实现列出已有的db配置缓存
		// todo 清除配置缓存
		return nil
	},
}

func init() {
	DbCmd.AddCommand(dbToStructCmd)
}
