package define

type Table struct {
	Name    string `gorm:"column:table_name"`
	Comment string `gorm:"column:table_Comment"`
}

type TableDDL struct {
	FieldName  string `gorm:"column:field_name"`
	FieldType  string `gorm:"column:field_type"`
	FieldLen   int    `gorm:"column:field_len"`
	FieldIdx   string `gorm:"column:field_idx"`
	FieldEmpty string `gorm:"column:field_empty"`
	FieldDesc  string `gorm:"column:field_desc"`
}

// TableColumn 存储 information_schema.COLUMNS 表中我们需要的一些字段（单列信息）
type TableColumn struct {
	// 列的名称
	ColumnName string `gorm:"column:COLUMN_NAME"`
	// 列的数据类型 varchar
	DataType string `gorm:"column:DATA_TYPE"`
	// 列是否允许为NULL，值为 YES or NO
	IsNullable string `gorm:"column:IS_NULLABLE"`
	// 列是否被索引
	ColumnKey string `gorm:"column:COLUMN_KEY"`
	// 列的数据类型信息 varchar(512)
	ColumnType string `gorm:"column:COLUMN_TYPE"`
	// 列的注释信息
	ColumnComment string `gorm:"column:COLUMN_COMMENT"`
}

// TypeMapping 数据库的数据类型与Go类型的映射关系
var TypeMapping = map[string]string{
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
