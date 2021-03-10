// @Description: 数据库表字段转换为结构体
// @Author: Arvin
// @Date: 2021/3/9 11:46 下午
package database

import (
	"github.com/spf13/cobra"
)

// 参数定义

// TODO 需要实现
var dbToStructCmd = &cobra.Command{
	Use:   "dts",
	Short: "转换表结构为gorm的模型结构",
	Long:  "转换表结构为gorm的模型结构",
	RunE: func(cmd *cobra.Command, args []string) error {
		// 获取数据库对象
		// 连接数据库
		// 获取表中列的信息
		// 创建结构体模版对象
		// 将数据库查询结果转为结构体
		// 将转换之后的结构体用模版解析渲染，并输出到标准控制台
		return nil
	},
}
