// @Description: 数据库表字段转换为结构体
// @Author: Arvin
// @Date: 2021/3/9 11:46 下午
package struct_cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// TODO 需要实现
var DbToStructCmd = &cobra.Command{
	Use:   "dts",
	Short: "database table convert struct-cmd",
	Long:  "database table convert struct-cmd",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println(111)
		return nil
	},
}
