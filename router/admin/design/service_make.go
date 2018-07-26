package design

import (
	"github.com/foxiswho/shop-go/module/db"
	"fmt"
	"text/template"
	"os"
	"strings"
	"github.com/foxiswho/shop-go/module/conf"
	"net/http"
	"github.com/foxiswho/shop-go/module/context"
	"github.com/labstack/echo"
)

func ServiceMakeHandler(c *context.BaseContext) error {

	sql := "show tables"
	result, err := db.Db().Engine.QueryString(sql)
	fmt.Println("err", err)
	fmt.Println("result", result)
	if err != nil {
		return c.JSON(http.StatusOK, echo.Map{
			"code":    http.StatusBadRequest,
			"message": "数据库错误:" + err.Error(),
		})
	} else {
		template_file := "./template/design/make/service.go.tpl"
		service_path := "./models/crud"
		field := "Tables_in_" + conf.Conf.DB.Name
		for i, val := range result {
			fmt.Println("result index=>", i)
			fmt.Println("result val=>", val)
			fmt.Println("result val=>", val[field])
			tmpl, err := template.ParseFiles(template_file)
			fmt.Println("template err", err)
			fmt.Println("template err", tmpl)
			//val[field] = "admin_menu"
			data := make(map[string]interface{})
			data["tables"] = val[field]
			//data["tables_first"] = gonicCasedName(val[field])
			data["tables_Camel_Case"] = LintGonicMapper.Table2Obj(val[field])
			fmt.Println("data=>", data)

			//
			//
			//
			err = os.MkdirAll(service_path, os.ModePerm)
			if err != nil {
				fmt.Println("%s", err)
			} else {
				fmt.Println("Create Directory OK! ", service_path)
			}
			service_file := service_path + "/" + val[field] + ".go"
			fmt.Println("Create file :", service_file)
			fmt.Println("Create file :", service_file)
			fmt.Println("Create file :", service_file)
			file, err := os.OpenFile(service_file, os.O_CREATE|os.O_WRONLY, os.ModePerm)
			if err != nil {
				fmt.Println("open failed err:", err)
				return c.JSON(http.StatusOK, echo.Map{
					"code":    http.StatusBadRequest,
					"message": service_file+" 目录中 不能创建文件或不能创建目录 error:" + err.Error(),
				})
			} else {
				err = tmpl.Execute(file, data)
				fmt.Println("tmpl.Execute=>", err)
			}
			//break
		}
		return c.JSON(http.StatusOK, echo.Map{
			"code":    http.StatusOK,
			"message": "生成成功",
		})
	}
	return nil
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

/**
 * 字符串首字母转化为大写 ios_bbbbbbbb -> iosBbbbbbbbb
 */
func strFirstToUpper(str string) string {
	temp := strings.Split(str, "_")
	var upperStr string
	for y := 0; y < len(temp); y++ {
		vv := []rune(temp[y])
		if y != 0 {
			for i := 0; i < len(vv); i++ {
				if i == 0 {
					vv[i] -= 32
					upperStr += string(vv[i]) // + string(vv[i+1])
				} else {
					upperStr += string(vv[i])
				}
			}
		}
	}
	return temp[0] + upperStr
}
