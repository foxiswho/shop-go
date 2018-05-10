package db

import (
	"time"

	_ "github.com/go-sql-driver/mysql"

	"fmt"
	"github.com/xormplus/xorm"
	. "github.com/foxiswho/shop-go/conf"
	"github.com/foxiswho/shop-go/module/log"
	"strconv"
	"reflect"
)

//结构体
type Db struct {
	Engine        *xorm.Engine
	FilterSession *xorm.Session
}

//全局变量
var db *Db

//数据库xiangg dsn 格式化
func dsn() string {
	db_user := Conf.DB.UserName
	db_pass := Conf.DB.Pwd
	db_host := Conf.DB.Host
	db_port := Conf.DB.Port
	db_name := Conf.DB.Name
	dsn := db_user + ":" + db_pass + "@tcp(" + db_host + ":" + db_port + ")/" + db_name + "?charset=utf8mb4&parseTime=True&loc=Asia%2FShanghai"
	return dsn
}

//创建数据库连接
func newDB() (*Db, error) {
	var err error
	db = new(Db)
	db.Engine, err = xorm.NewEngine("mysql", dsn())
	if err != nil {
		fmt.Println("NewEngine", err)
		panic(err.Error())
	}
	db.Engine.ShowSQL(true)
	locat, _ := time.LoadLocation("Asia/Shanghai")
	db.Engine.TZLocation = locat
	return db, nil
}

//快捷调研
func DB() *Db {
	if db == nil {
		log.Debugf("Model NewDB")
		newDb, err := newDB()
		if err != nil {
			panic(err)
		}
		newDb.Engine.SetMaxIdleConns(10)
		newDb.Engine.SetMaxIdleConns(100)
		//newDb.Engine.SetLogger(orm.Logger{})
		db = newDb
	}
	return db
}

//sql map XML文件
func SqlMapFile(directory, extension string) (*xorm.XmlSqlMap) {
	return xorm.Xml(directory, extension)
}


//初始化
//func Init() {
//	DB()
//}


type QuerySession struct {
	Session *xorm.Session
}

var Query *QuerySession

func Filter(where map[int]QueryCondition) *xorm.Session {
	db := DB()
	Query = new(QuerySession)
	if len(where) > 0 {
		i := 1
		for _, qc := range where {
			condition :=qc.Condition
			key:=qc.Field+qc.Operation
			//fmt.Println(k, condition, reflect.TypeOf(condition))
			//fmt.Println("?号个数为", strings.Count(k, "?"))
			isEmpty := false
			isMap := false
			arrCount := 0
			str := ""
			var arr []string
			switch condition.(type) {
			case string:
				//是字符时做的事情
				isEmpty = condition == ""
			case int:
				//是整数时做的事情
			case []string:
				isMap = true
				arr = condition.([]string)
				arrCount = len(arr)
				isEmpty = arrCount == 0
				for j, val := range arr {
					if j > 0 {
						str += ","
					}
					str += val
				}
			case []int:
				isMap = true
				arrInt := condition.([]int)
				arrCount = len(arrInt)
				isEmpty = arrCount == 0
				for j, val := range arrInt {
					if j > 0 {
						str += ","
					}
					str += strconv.Itoa(val)
				}
			}
			if isEmpty {
				FilterWhereAnd(db, i, key, "")
			} else if !isEmpty {
				//是数组
				if isMap {
					FilterWhereAnd(db, i,key, str)
				} else {
					//不是数组
					FilterWhereAnd(db, i, key, condition)
				}
			} else {
				fmt.Println("其他还没有收录")
			}
			i++
		}
	} else {
		//初始化
		Query.Session = db.Engine.Limit(20, 0)
	}

	return Query.Session
}
func FilterWhereAnd(db *Db, i int, key string, value ...interface{}) {
	fmt.Println("key", key)
	fmt.Println("value", value)
	fmt.Println("TypeOf", reflect.TypeOf(value))
	if i == 1 {
		Query.Session = DB().Engine.Where(key, value...)
	} else {
		Query.Session = Query.Session.And(key, value...)
	}
}
