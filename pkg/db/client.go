package db

import (
	"fmt"
	"time"

	"github.com/qingyunjun/tool-kit/pkg/db/define"

	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Client struct {
	cfg Config
	db  *gorm.DB
}

func NewClient(conf Config) *Client {
	return &Client{
		cfg: conf,
	}
}

func (c *Client) ConnectDB() (db *gorm.DB, err error) {
	newLogger := logger.Default.LogMode(logger.Warn)
	if c.cfg.LogMode {
		newLogger = newLogger.LogMode(logger.Info)
	}

	if c.cfg.DBType == "postgres" {
		dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Shanghai",
			c.cfg.Host, c.cfg.UserName, c.cfg.PassWord, c.cfg.DBName, c.cfg.Port)
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
			Logger: newLogger,
		})
	} else if c.cfg.DBType == "mysql" {
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			c.cfg.UserName, c.cfg.PassWord, c.cfg.Host, c.cfg.Port, c.cfg.DBName)
		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
			Logger: newLogger,
		})
	} else {
		return nil, gorm.ErrUnsupportedDriver
	}

	if err != nil {
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	if err := sqlDB.Ping(); err != nil {
		return nil, err
	}

	if c.cfg.MaxIdleConns > 0 {
		sqlDB.SetMaxIdleConns(c.cfg.MaxIdleConns)
	}

	if c.cfg.MaxOpenConns > 0 {
		sqlDB.SetMaxOpenConns(c.cfg.MaxOpenConns)
	}

	if c.cfg.ConnMaxLifetime > 0 {
		sqlDB.SetConnMaxLifetime(time.Duration(c.cfg.ConnMaxLifetime) * time.Second)
	}
	c.db = db

	return db, nil
}

// GetTableColumnInfo 查询 information_schema.COLUMNS 表，获取指定表的列信息
func (c *Client) GetTableColumnInfo(dbName, tableName string) ([]*define.TableColumn, error) {
	// 初始化返回值
	tableColumns := make([]*define.TableColumn, 0)

	sql := "select COLUMN_NAME,DATA_TYPE,COLUMN_KEY,IS_NULLABLE,COLUMN_TYPE, COLUMN_COMMENT from COLUMNS where TABLE_SCHEMA = ? and TABLE_NAME = ?"
	if err := c.db.Raw(sql, dbName, tableName).Scan(&tableColumns).Error; err != nil {
		return nil, err
	}

	return tableColumns, nil
}

func (c *Client) name() {

}

// GetTables 查询存在的所有表
func (c *Client) GetTables() ([]*define.Table, error) {
	tables := make([]*define.Table, 0)
	if err := c.db.Raw("select relname as table_name,(select description from pg_description where " +
		"objoid=oid and objsubid=0) as table_comment from pg_class where relkind ='r' and relname" +
		" NOT LIKE 'pg%' AND relname NOT LIKE 'sql_%' order by table_name;").Scan(&tables).Error; err != nil {
		return nil, err
	}
	return tables, nil
}

// GetTableDDL 获取表结构
func (c *Client) GetTableDDL(tableName string) ([]*define.TableDDL, error) {
	tbDDL := make([]*define.TableDDL, 0)
	sql := fmt.Sprintf("select a.attname as field_name, b.typname as field_type, "+
		"(case when atttypmod-4>0 then atttypmod-4 else 0 end) as field_len, (case when (select count(*)"+
		" from pg_constraint where conrelid = a.attrelid and conkey[1]=attnum and contype='p')>0 then "+
		"'PRI' when (select count(*) from pg_constraint where conrelid = a.attrelid and conkey[1]=attnum"+
		" and contype='u')>0 then 'UNI' when (select count(*) from pg_constraint where conrelid = a.attrelid"+
		" and conkey[1]=attnum and contype='f')>0 then 'FRI' else '' end) as field_idx, (case when a.attnotnull=true "+
		"then 'NO' else 'YES' end) as field_empty, col_description(a.attrelid,a.attnum) as field_desc from pg_attribute a,pg_type b "+
		"where attstattarget=-1 and a.atttypid=b.oid and attrelid = (select oid from pg_class where relname ='%s');", tableName)
	if err := c.db.Raw(sql).Scan(&tbDDL).Error; err != nil {
		return nil, err
	}
	return tbDDL, nil
}
