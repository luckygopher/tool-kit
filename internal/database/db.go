// @Description: 数据库操作
// @Author: Arvin
// @Date: 2021/3/10 2:11 下午
package database

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
)

// 数据库连接的核心对象
type DBModel struct {
	DBEngine *sql.DB
	DBConfig *DBConfig
}

// 数据库配置信息
type DBConfig struct {
	// 驱动类型（如：mysql）
	DBType string
	// 主机地址
	Host string
	// 数据库名
	DBName string
	// 用户名
	UserName string
	// 密码
	PassWord string
	// 字符集
	CharSet string
}

// 存储 information_schema.COLUMNS 表中我们需要的一些字段（单列信息）
type TableColumn struct {
	// 列的名称
	ColumnName string
	// 列的数据类型 varchar
	DataType string
	// 列是否允许为NULL，值为 YES or NO
	IsNullable string
	// 列是否被索引
	ColumnKey string
	// 列的数据类型信息 varchar(512)
	ColumnType string
	// 列的注释信息
	ColumnComment string
}

// 数据库中数据类型映射到go中的数据类型
var DBTypeToStructType = map[string]string{
	"int":        "int32",
	"tinyint":    "int8",
	"smallint":   "int",
	"mediumint":  "int64",
	"bigint":     "int64",
	"bit":        "int",
	"bool":       "bool",
	"enum":       "string",
	"set":        "string",
	"varchar":    "string",
	"char":       "string",
	"tinytext":   "string",
	"mediumtext": "string",
	"text":       "string",
	"longtext":   "string",
	"blob":       "string",
	"tinyblob":   "string",
	"mediumblob": "string",
	"longblob":   "string",
	"date":       "time.Time",
	"datetime":   "time.Time",
	"timestamp":  "time.Time",
	"time":       "time.Time",
	"float":      "float64",
	"double":     "float64",
}

// 初始化连接对象
func NewDBModel(config *DBConfig) *DBModel {
	return &DBModel{
		DBConfig: config,
	}
}

// 创建数据库连接
// Open initialize a new db connection, need to import driver first,e.g:
// import _ "github.com/go-sql-driver/mysql"
// sql.Open("mysql", "user:pwd@tcp(host)/dbname?charset=utf8mb4&parseTime=True&loc=Local")
func (db *DBModel) Connect() error {
	var err error
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s&parseTime=True&loc=Local",
		db.DBConfig.UserName,
		db.DBConfig.PassWord,
		db.DBConfig.Host,
		db.DBConfig.DBName,
		db.DBConfig.CharSet,
	)

	if db.DBEngine, err = sql.Open(db.DBConfig.DBType, dsn); err != nil {
		log.Printf("创建数据库连接失败：%s", err)
		return err
	}
	return nil
}

// 查询 information_schema.COLUMNS 表，获取指定表的列信息
func (db *DBModel) GetTableColumnInfo(dbName, tableName string) ([]*TableColumn, error) {
	// 初始化返回值
	var (
		tableColumn  *TableColumn
		tableColumns = make([]*TableColumn, 0)
	)

	sql := "select * from COLUMNS where TABLE_SCHEMA = ? and TABLE_NAME = ?"
	rows, err := db.DBEngine.Query(sql, dbName, tableName)
	if err != nil {
		return nil, err
	}
	if rows == nil {
		return nil, errors.New("无数据")
	}
	// 释放资源
	defer rows.Close()
	// 遍历查询结果
	for rows.Next() {
		if err := rows.Scan(&tableColumn); err != nil {
			log.Printf("遍历查询结果失败:%s", err)
			return nil, err
		}
		tableColumns = append(tableColumns, tableColumn)
	}
	return tableColumns, nil
}
