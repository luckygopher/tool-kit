package document

import (
	"fmt"
	"strconv"

	"github.com/qingyunjun/tool-kit/pkg/db/define"

	"baliance.com/gooxml/color"
	"baliance.com/gooxml/document"
	"baliance.com/gooxml/measurement"
	"baliance.com/gooxml/schema/soo/wml"
)

type Writer struct {
	doc *document.Document
}

func NewWriter() *Writer {
	doc := document.New()
	return &Writer{doc: doc}
}

func (w *Writer) WriterTable(tableName, tableComment string, tableInfo []*define.TableDDL) {
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
