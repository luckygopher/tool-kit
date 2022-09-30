package db

import (
	"fmt"
	"os"
	"strings"
	"text/template"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"

	"github.com/luckygopher/tool-kit/pkg/db/define"
)

const structTpl = `type {{.TableName | ToCamelCase}} struct {
	{{range .Columns}}{{$length := len .Comment}}{{if gt $length 0}}
	// {{.Comment}}{{else}}// {{.Name}}{{end}}
	{{$typeLen := len .Type}}{{if gt $typeLen 0}}{{.Name | ToCamelCase}} {{.Type}} {{.Tag}}{{else}}{{.Name}}{{end}}{{end}}
}
func (m *{{.TableName | ToCamelCase}}) TableName() string {
	return "{{.TableName}}"
}
`

type StructTemplate struct {
	structTpl string // 模版内容
	FilePath  string // 输出文件路径
}

// StructColumn 存储转换后的 Go 结构体中的所有字段信息
type StructColumn struct {
	Name    string
	Type    string
	Tag     string
	Comment string
}

// StructTemplateDB 存储最终用于渲染的模版对象信息
type StructTemplateDB struct {
	TableName string
	Columns   []*StructColumn
}

func NewStructTemplate() *StructTemplate {
	return &StructTemplate{structTpl: structTpl}
}

// AssemblyColumns 对通过查询 COLUMNS 表所组装得到的 tbColumns 进行进一步的分解和转换
func (t *StructTemplate) AssemblyColumns(tbColumns []*define.TableColumn) []*StructColumn {
	tplColumns := make([]*StructColumn, 0, len(tbColumns))
	for _, column := range tbColumns {
		// 取得字段数据类型
		dataType := define.TypeMapping[column.DataType]
		// 判断是否有无符号
		if strings.Contains(column.ColumnType, "unsigned") {
			dataType = fmt.Sprintf("u%s", dataType)
		}
		// 组装最终数据
		tplColumns = append(tplColumns, &StructColumn{
			Name:    column.ColumnName,
			Type:    dataType,
			Tag:     fmt.Sprintf("`gorm:\"column:"+"%s"+"\"`", column.ColumnName),
			Comment: column.ColumnComment,
		})
	}
	return tplColumns
}

// Generate 用转换之后的结构体去渲染模版
// template.Must 方法判断返回的 *Template 是否有错误，引发panic，导致程序崩溃(如果模版解析错误，则直接让程序挂掉)
func (t *StructTemplate) Generate(tableName string, tmpColumns []*StructColumn) error {
	var (
		writer *os.File = os.Stdout // 默认输出到终端
		err    error
	)
	tpl := template.Must(template.New("sql-struct").Funcs(template.FuncMap{
		"ToCamelCase": func(s string) string {
			// 蛇形转大写驼峰
			s = strings.Replace(s, "_", " ", -1)
			// Title 方法将按空格分割的字符串首字母转为大写
			title := cases.Title(language.English).String(s)
			return strings.Replace(title, " ", "", -1)
		},
	}).Parse(t.structTpl))
	// 传入模版的数据
	tplDB := StructTemplateDB{
		TableName: tableName,
		Columns:   tmpColumns,
	}
	// 如果未指定输出文件路径
	if t.FilePath != "" {
		if writer, err = os.OpenFile(t.FilePath, os.O_WRONLY|os.O_CREATE, 0666); err != nil {
			return err
		}
	}
	// 渲染模版
	if err = tpl.Execute(writer, tplDB); err != nil {
		return err
	}
	return nil
}
