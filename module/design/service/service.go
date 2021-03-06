package service

import (
	"github.com/foxiswho/shop-go/module/db"
	"text/template"
	"os"
	"strings"
	"github.com/foxiswho/shop-go/module/conf"
	"github.com/foxiswho/shop-go/module/log"
	"github.com/foxiswho/shop-go/util"
)

func MakeService(template_file, service_path string) error {

	sql := "show tables"
	result, err := db.Db().Engine.QueryString(sql)
	if err != nil {
		return util.NewError("数据库错误:" + err.Error())
	} else {
		if len(service_path) < 1 {
			service_path = "./models/crud"
		}
		field := "Tables_in_" + conf.Conf.DB.Name
		for _, val := range result {
			tmpl, err := template.ParseFiles(template_file)
			data := make(map[string]interface{})
			data["tables"] = val[field]
			//data["tables_first"] = gonicCasedName(val[field])
			data["tables_Camel_Case"] = LintGonicMapper.Table2Obj(val[field])
			err = os.MkdirAll(service_path, os.ModePerm)
			if err != nil {
				log.Debugf("Create Directory ERROR! : %v", err)
			}
			log.Debugf("Create Directory OK! : %v", service_path)
			service_file := service_path + "/" + val[field] + ".go"
			file, err := os.OpenFile(service_file, os.O_CREATE|os.O_WRONLY, os.ModePerm)
			if err != nil {
				return util.NewError(service_file + " 目录中 不能创建文件或不能创建目录 error:" + err.Error())
			} else {
				err = tmpl.Execute(file, data)
				if err != nil {
					log.Debugf(" Execute: %v", err)
				}
			}
			//break
		}
		return nil
	}
}

func gonicCasedName(name string) string {
	newstr := make([]rune, 0, len(name)+3)
	for idx, chr := range name {
		if isASCIIUpper(chr) && idx > 0 {
			if !isASCIIUpper(newstr[len(newstr)-1]) {
				newstr = append(newstr, '_')
			}
		}

		if !isASCIIUpper(chr) && idx > 1 {
			l := len(newstr)
			if isASCIIUpper(newstr[l-1]) && isASCIIUpper(newstr[l-2]) {
				newstr = append(newstr, newstr[l-1])
				newstr[l-1] = '_'
			}
		}

		newstr = append(newstr, chr)
	}
	return strings.ToLower(string(newstr))
}
func isASCIIUpper(r rune) bool {
	return 'A' <= r && r <= 'Z'
}

func toASCIIUpper(r rune) rune {
	if 'a' <= r && r <= 'z' {
		r -= ('a' - 'A')
	}
	return r
}

// GonicMapper implements IMapper. It will consider initialisms when mapping names.
// E.g. id -> ID, user -> Admin and to table names: UserID -> user_id, MyUID -> my_uid
type GonicMapper map[string]bool

func (mapper GonicMapper) Obj2Table(name string) string {
	return gonicCasedName(name)
}

func (mapper GonicMapper) Table2Obj(name string) string {
	newstr := make([]rune, 0)

	name = strings.ToLower(name)
	parts := strings.Split(name, "_")

	for _, p := range parts {
		_, isInitialism := mapper[strings.ToUpper(p)]
		for i, r := range p {
			if i == 0 || isInitialism {
				r = toASCIIUpper(r)
			}
			newstr = append(newstr, r)
		}
	}

	return string(newstr)
}

var LintGonicMapper = GonicMapper{
	"API":   true,
	"ASCII": true,
	"CPU":   true,
	"CSS":   true,
	"DNS":   true,
	"EOF":   true,
	"GUID":  true,
	"HTML":  true,
	"HTTP":  true,
	"HTTPS": true,
	"ID":    true,
	"IP":    true,
	"JSON":  true,
	"LHS":   true,
	"QPS":   true,
	"RAM":   true,
	"RHS":   true,
	"RPC":   true,
	"SLA":   true,
	"SMTP":  true,
	"SSH":   true,
	"TLS":   true,
	"TTL":   true,
	"UI":    true,
	"UID":   true,
	"UUID":  true,
	"URI":   true,
	"URL":   true,
	"UTF8":  true,
	"VM":    true,
	"XML":   true,
	"XSRF":  true,
	"XSS":   true,
}