package cmd

import (
	"fmt"
	"strconv"

	"baliance.com/gooxml/color"
	"baliance.com/gooxml/document"
	"baliance.com/gooxml/measurement"
	"baliance.com/gooxml/schema/soo/wml"

	"github.com/urfave/cli/v2"

	"github.com/pkg/errors"
	"github.com/qingyunjun/tool-kit/config"
	"github.com/qingyunjun/tool-kit/pkg/db"
)

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

func dbToDocxCmd() *cli.Command {
	var (
		filePath   string // 结果输出文件路径
		configPath string // 配置文件
	)
	filePathFlag := &cli.StringFlag{
		Name: "filePath", Usage: "结果输出文件路径", Aliases: []string{"f"}, Destination: &filePath,
	}
	configPathFlag := &cli.StringFlag{
		Name: "configPath", Usage: "配置文件路径", Aliases: []string{"c"}, Destination: &configPath,
	}

	cmd := &cli.Command{
		Name:  "docx",
		Usage: "导出表结构为docx",
		Flags: []cli.Flag{
			filePathFlag,
			configPathFlag,
		},
		Action: func(ctx *cli.Context) error {
			config.ParseConfig(configPath)
			writer := NewWriter()
			// 获取数据库对象
			dbObj, err := db.ConnectDB(config.C.Database)
			if err != nil {
				return errors.Errorf("数据库连接失败:%s", err)
			}
			// 查询所有表
			tables := make([]*Table, 0)
			dbObj.Raw("select relname as table_name,(select description from pg_description where " +
				"objoid=oid and objsubid=0) as table_comment from pg_class where relkind ='r' and relname" +
				" NOT LIKE 'pg%' AND relname NOT LIKE 'sql_%' order by table_name;").Scan(&tables)
			for _, tb := range tables {
				tbDDL := make([]*TableDDL, 0)
				// 获取表结构
				sql := fmt.Sprintf("select a.attname as field_name, b.typname as field_type, "+
					"(case when atttypmod-4>0 then atttypmod-4 else 0 end) as field_len, (case when (select count(*)"+
					" from pg_constraint where conrelid = a.attrelid and conkey[1]=attnum and contype='p')>0 then "+
					"'PRI' when (select count(*) from pg_constraint where conrelid = a.attrelid and conkey[1]=attnum"+
					" and contype='u')>0 then 'UNI' when (select count(*) from pg_constraint where conrelid = a.attrelid"+
					" and conkey[1]=attnum and contype='f')>0 then 'FRI' else '' end) as field_idx, (case when a.attnotnull=true "+
					"then 'NO' else 'YES' end) as field_empty, col_description(a.attrelid,a.attnum) as field_desc from pg_attribute a,pg_type b "+
					"where attstattarget=-1 and a.atttypid=b.oid and attrelid = (select oid from pg_class where relname ='%s');", tb.Name)
				dbObj.Raw(sql).Scan(&tbDDL)
				writer.WriterTable(tb.Name, tb.Comment, tbDDL)
			}
			if err := writer.Save(filePath); err != nil {
				fmt.Printf("写入失败 %s", err)
				return err
			}
			return nil
		},
	}

	return cmd
}

type Writer struct {
	doc *document.Document
}

func NewWriter() *Writer {
	doc := document.New()
	return &Writer{doc: doc}
}

func (w *Writer) WriterTable(tableName, tableComment string, tableInfo []*TableDDL) {
	// 写入表名
	w.doc.AddParagraph().AddRun().AddText(fmt.Sprintf("%s:%s", tableName, tableComment))
	// 添加一个表格
	table := w.doc.AddTable()
	table.Properties().SetWidthPercent(100)

	borders := table.Properties().Borders()
	borders.SetAll(wml.ST_BorderSingle, color.Auto, measurement.Zero)

	row := table.AddRow()
	row.AddCell().AddParagraph().AddRun().AddText("序号")
	row.AddCell().AddParagraph().AddRun().AddText("字段名称")
	row.AddCell().AddParagraph().AddRun().AddText("字段类型")
	row.AddCell().AddParagraph().AddRun().AddText("字段长度")
	row.AddCell().AddParagraph().AddRun().AddText("默认值")
	row.AddCell().AddParagraph().AddRun().AddText("允许为空")
	row.AddCell().AddParagraph().AddRun().AddText("备注")

	for idx, val := range tableInfo {
		row = table.AddRow()
		row.AddCell().AddParagraph().AddRun().AddText(strconv.Itoa(idx + 1))
		row.AddCell().AddParagraph().AddRun().AddText(val.FieldName)
		row.AddCell().AddParagraph().AddRun().AddText(val.FieldType)
		row.AddCell().AddParagraph().AddRun().AddText(strconv.Itoa(val.FieldLen))
		row.AddCell().AddParagraph().AddRun().AddText("")
		row.AddCell().AddParagraph().AddRun().AddText(val.FieldEmpty)
		row.AddCell().AddParagraph().AddRun().AddText(val.FieldDesc)
	}
	w.doc.AddParagraph()
}

func (w *Writer) Save(fileName string) error {
	return w.doc.SaveToFile(fileName)
}

func init() {
	rootCmd.Commands = append(rootCmd.Commands, dbToDocxCmd())
}
