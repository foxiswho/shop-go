// Copyright 2017 The Xorm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package models

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"strings"
	"text/template"

	"github.com/go-xorm/core"
	"github.com/lunny/log"

	_ "github.com/denisenkom/go-mssqldb"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	_ "github.com/ziutek/mymysql/godrv"
	"github.com/go-xorm/xorm"
	"github.com/foxiswho/shop-go/module/conf"
)

var (
	genJson bool = true
)

func printReversePrompt(flag string) {
}

type Tmpl struct {
	Tables  []*core.Table
	Imports map[string]string
	Models  string
}

func dirExists(dir string) bool {
	d, e := os.Stat(dir)
	switch {
	case e != nil:
		return false
	case !d.IsDir():
		return false
	}

	return true
}

//数据库xiangg dsn 格式化
func dsn() string {
	db_user := conf.Conf.DB.UserName
	db_pass := conf.Conf.DB.Pwd
	db_host := conf.Conf.DB.Host
	db_port := conf.Conf.DB.Port
	db_name := conf.Conf.DB.Name
	dsn := db_user + ":" + db_pass + "@tcp(" + db_host + ":" + db_port + ")/" + db_name + "?charset=utf8mb4&parseTime=True&loc=Asia%2FShanghai"
	return dsn
}

func MakeModels(dir, model string) {
	var isMultiFile bool = true
	supportComment = true

	curPath, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Getwd >>> ", curPath)
	var genDir string
	var filterPat *regexp.Regexp
	genDir = path.Join(curPath, model)

	if len(dir) <= 0 {
		log.Errorf("%v", dir)
		return
	}

	_, err = os.Stat(dir)
	if err != nil {
		log.Errorf("Template %v file is not exist", dir)
		return
	}

	var langTmpl LangTmpl
	var ok bool
	var lang string = "go"
	var prefix string = "" //[SWH|+]

	if langTmpl, ok = langTmpls[lang]; !ok {
		fmt.Println("Unsupported programing language", lang)
		return
	}

	os.MkdirAll(genDir, os.ModePerm)
	Orm, err := xorm.NewEngine("mysql", dsn())
	if err != nil {
		log.Errorf("%v", err)
		return
	}

	tables, err := Orm.DBMetas()
	if err != nil {
		log.Errorf("%v", err)
		return
	}
	if filterPat != nil && len(tables) > 0 {
		size := 0
		for _, t := range tables {
			if filterPat.MatchString(t.Name) {
				tables[size] = t
				size++
			}
		}
		tables = tables[:size]
	}

	filepath.Walk(dir, func(f string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}

		if info.Name() == "config" {
			return nil
		}

		bs, err := ioutil.ReadFile(f)
		if err != nil {
			log.Errorf("%v", err)
			return err
		}

		t := template.New(f)
		t.Funcs(langTmpl.Funcs)

		tmpl, err := t.Parse(string(bs))
		if err != nil {
			log.Errorf("%v", err)
			return err
		}

		var w *os.File
		fileName := info.Name()
		newFileName := fileName[:len(fileName)-4]
		ext := path.Ext(newFileName)
		fmt.Println("table make start ", isMultiFile)
		if !isMultiFile {
			fmt.Println("isMultiFile=", isMultiFile)
			w, err = os.Create(path.Join(genDir, newFileName))
			if err != nil {
				log.Errorf("%v", err)
				return err
			}

			imports := langTmpl.GenImports(tables)

			tbls := make([]*core.Table, 0)
			for _, table := range tables {
				//[SWH|+]
				if prefix != "" {
					table.Name = strings.TrimPrefix(table.Name, prefix)
				}
				tbls = append(tbls, table)
			}

			newbytes := bytes.NewBufferString("")

			t := &Tmpl{Tables: tbls, Imports: imports, Models: model}
			err = tmpl.Execute(newbytes, t)
			if err != nil {
				log.Errorf("%v", err)
				return err
			}

			tplcontent, err := ioutil.ReadAll(newbytes)
			if err != nil {
				log.Errorf("%v", err)
				return err
			}
			var source string
			if langTmpl.Formater != nil {
				source, err = langTmpl.Formater(string(tplcontent))
				if err != nil {
					log.Errorf("%v", err)
					return err
				}
			} else {
				source = string(tplcontent)
			}

			w.WriteString(source)
			w.Close()

		} else {
			log.Debugf("table make start")
			fmt.Println("table make start")
			for _, table := range tables {
				//[SWH|+]
				if prefix != "" {
					table.Name = strings.TrimPrefix(table.Name, prefix)
				}
				// imports
				tbs := []*core.Table{table}
				imports := langTmpl.GenImports(tbs)

				w, err := os.Create(path.Join(genDir, table.Name+ext))
				if err != nil {
					log.Errorf("%v", err)
					return err
				}
				defer w.Close()

				newbytes := bytes.NewBufferString("")

				t := &Tmpl{Tables: tbs, Imports: imports, Models: model}
				err = tmpl.Execute(newbytes, t)
				if err != nil {
					log.Errorf("%v", err)
					return err
				}

				tplcontent, err := ioutil.ReadAll(newbytes)
				if err != nil {
					log.Errorf("%v", err)
					return err
				}
				var source string
				if langTmpl.Formater != nil {
					source, err = langTmpl.Formater(string(tplcontent))
					if err != nil {
						log.Errorf("%v-%v", err, string(tplcontent))
						return err
					}
				} else {
					source = string(tplcontent)
				}

				w.WriteString(source)
				w.Close()
				log.Debugf("table %v make success", table.Name)
			}
		}

		return nil
	})

}
