// Package database @Description: 数据库操作
// @Author: Arvin
// @Date: 2021/3/10 2:11 下午
package db

import (
	"database/sql"
	"errors"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

// DBModel 数据库连接的核心对象
type DBModel struct {
	DBEngine *sql.DB
	DBConfig *DBConfig
}

// TableColumn 存储 information_schema.COLUMNS 表中我们需要的一些字段（单列信息）
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

// DBTypeToStructType 数据库中数据类型映射到go中的数据类型
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

// NewDBModel 初始化连接对象
func NewDBModel(conf DBConfig, dbName string) *DBModel {
	return &DBModel{
		DBConfig: &DBConfig{
			DBType:   conf.DBType,
			Host:     conf.Host,
			Port:     conf.Port,
			DBName:   dbName,
			UserName: conf.UserName,
			PassWord: conf.PassWord,
			CharSet:  conf.CharSet,
		},
	}
}

// Connect 创建数据库连接
// Open initialize a new db connection, need to import driver first,e.g:
// import _ "github.com/go-sql-driver/mysql"
// sql.Open("mysql", "user:pwd@tcp(host)/dbname?charset=utf8mb4&parseTime=True&loc=Local")
func (db *DBModel) Connect() error {
	var err error
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=Local",
		db.DBConfig.UserName,
		db.DBConfig.PassWord,
		db.DBConfig.Host,
		db.DBConfig.Port,
		db.DBConfig.DBName,
		db.DBConfig.CharSet,
	)

	if db.DBEngine, err = sql.Open(db.DBConfig.DBType, dsn); err != nil {
		log.Printf("创建数据库连接失败：%s", err)
		return err
	}
	return nil
}

// GetTableColumnInfo 查询 information_schema.COLUMNS 表，获取指定表的列信息
func (db *DBModel) GetTableColumnInfo(dbName, tableName string) ([]*TableColumn, error) {
	// 初始化返回值
	tableColumns := make([]*TableColumn, 0)

	sql := "select COLUMN_NAME,DATA_TYPE,COLUMN_KEY,IS_NULLABLE,COLUMN_TYPE, COLUMN_COMMENT from COLUMNS where TABLE_SCHEMA = ? and TABLE_NAME = ?"
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
		var tableColumn TableColumn
		// 此处绑定数据
		if err := rows.Scan(
			&tableColumn.ColumnName,
			&tableColumn.DataType,
			&tableColumn.ColumnKey,
			&tableColumn.IsNullable,
			&tableColumn.ColumnType,
			&tableColumn.ColumnComment,
		); err != nil {
			log.Printf("遍历查询结果失败:%s", err)
			return nil, err
		}
		tableColumns = append(tableColumns, &tableColumn)
	}
	return tableColumns, nil
}
