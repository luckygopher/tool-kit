package cmd

import (
	"github.com/pkg/errors"
	"github.com/qingyunjun/tool-kit/config"
	"github.com/qingyunjun/tool-kit/pkg/db"
	"github.com/urfave/cli/v2"
)

func dbToStructCmd() *cli.Command {
	// InformationDB 参数定义
	const InformationDB = "information_schema"

	var (
		tableName  string // 需要转换的表名
		filePath   string // 结果输出文件路径
		configPath string // 配置文件
	)
	tableNameFlag := &cli.StringFlag{
		Name: "tableName", Usage: "需要转换的表名", Aliases: []string{"t"}, Destination: &tableName,
	}
	filePathFlag := &cli.StringFlag{
		Name: "filePath", Usage: "结果输出文件路径", Aliases: []string{"f"}, Destination: &filePath,
	}
	configPathFlag := &cli.StringFlag{
		Name: "configPath", Usage: "配置文件路径", Aliases: []string{"c"}, Destination: &configPath,
	}

	cmd := &cli.Command{
		Name:  "dts",
		Usage: "转换表结构为gorm的模型结构",
		Flags: []cli.Flag{
			tableNameFlag,
			filePathFlag,
			configPathFlag,
		},
		Action: func(ctx *cli.Context) error {
			config.ParseConfig(configPath)
			// 获取数据库对象
			config.C.Database.DBName = InformationDB
			dbObj := db.NewClient(config.C.Database)
			// 连接数据库
			if _, err := dbObj.ConnectDB(); err != nil {
				return errors.Errorf("数据库连接失败:%s", err)
			}
			// 获取表中列的信息
			tableColumns, err := dbObj.GetTableColumnInfo(config.C.Database.DBName, tableName)
			if err != nil {
				return errors.Errorf("获取表中列的信息失败:%s", err)
			}
			// 创建结构体模版对象
			tmp := db.NewStructTemplate()
			// 设置输出文件路径
			tmp.FilePath = filePath
			// 将数据库查询结果转为结构体
			tmpColumns := tmp.AssemblyColumns(tableColumns)
			// 将转换之后的结构体用模版解析渲染，并输出
			if err := tmp.Generate(tableName, tmpColumns); err != nil {
				return errors.Errorf("渲染结构体模版错误:%s", err)
			}
			return nil
		},
	}

	return cmd
}

func init() {
	rootCmd.Commands = append(rootCmd.Commands, dbToStructCmd())
}
