
package {{.Models}}

{{$ilen := len .Imports}}
{{if gt $ilen 0}}
import (
	{{range .Imports}}"{{.}}"{{end}}
)
{{end}}

{{range .Tables}}
type {{Mapper .Name}} struct {
{{$table := .}}
{{range .ColumnsSeq}}{{$col := $table.GetColumn .}}	{{Mapper $col.Name}}	{{Type $col}} {{Tag $table $col}}
{{end}}

    //
	ExtData      interface{} `json:"ExtData" xorm:"- <- ->"`
}

//初始化
func New{{Mapper .Name}}() *{{Mapper .Name}}{
	return new({{Mapper .Name}})
}

{{end}}