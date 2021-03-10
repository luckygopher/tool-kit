// @Description: 数据库表字段转换为结构体
// @Author: Arvin
// @Date: 2021/3/9 11:46 下午
package database

import (
	"fmt"

	"github.com/pkg/errors"
	"github.com/qingyunjun/tool-kit/internal/database"
	"github.com/spf13/cobra"
)

// 参数定义
const InformationDB = "information_schema"

var (
	DBType    string // 数据库驱动类型
	Host      string // 主机地址
	DBName    string // 需要转换的表所在数据库名
	TableName string // 需要转换的表名
	UserName  string // 用户名
	PassWord  string // 密码
	CharSet   string // 字符集
)

// 数据库操作子命令 dts 定义
var dbToStructCmd = &cobra.Command{
	Use:   "dts",
	Short: "转换表结构为gorm的模型结构",
	Long:  "转换表结构为gorm的模型结构",
	RunE: func(cmd *cobra.Command, args []string) error {
		// 获取数据库对象
		db := database.NewDBModel(&database.DBConfig{
			DBType:   DBType,
			Host:     Host,
			DBName:   InformationDB,
			UserName: UserName,
			PassWord: PassWord,
			CharSet:  CharSet,
		})
		// 连接数据库
		if err := db.Connect(); err != nil {
			return errors.Errorf("数据库连接失败:%s", err)
		}
		// 获取表中列的信息
		tableColumns, err := db.GetTableColumnInfo(DBName, TableName)
		if err != nil {
			return errors.Errorf("获取表中列的信息失败:%s", err)
		}
		fmt.Println(tableColumns)
		// todo 创建结构体模版对象
		// todo 将数据库查询结果转为结构体
		// todo 将转换之后的结构体用模版解析渲染，并输出到标准控制台
		return nil
	},
}

func init() {
	// 定义命令行参数时需要注意-h和--help已经被占用
	dbToStructCmd.Flags().StringVarP(&DBType, "type", "t", "mysql", "数据库驱动类型")
	dbToStructCmd.Flags().StringVarP(&Host, "host", "d", "127.0.0.1", "主机地址")
	dbToStructCmd.Flags().StringVarP(&DBName, "dbname", "n", InformationDB, "数据库名")
	dbToStructCmd.Flags().StringVarP(&TableName, "table", "b", "", "表名")
	dbToStructCmd.Flags().StringVarP(&UserName, "user", "u", "", "用户名")
	dbToStructCmd.Flags().StringVarP(&PassWord, "password", "p", "", "密码")
	dbToStructCmd.Flags().StringVarP(&CharSet, "charset", "c", "utf8mb4", "字符集")
}
