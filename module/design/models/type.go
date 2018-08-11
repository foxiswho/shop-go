package models

import (
	"time"
	"reflect"
	"strings"
	"github.com/go-xorm/core"
)

var c_TIME_DEFAULT time.Time
// default sql type change to go types
func SQLType2Type(st core.SQLType) reflect.Type {
	name := strings.ToUpper(st.Name)
	switch name {
	case core.Bit, core.TinyInt, core.SmallInt, core.MediumInt, core.Int, core.Integer, core.Serial:
		return reflect.TypeOf(1)
	case core.BigInt, core.BigSerial:
		return reflect.TypeOf(int64(1))
	case core.Float, core.Real:
		return reflect.TypeOf(float32(1))
	case core.Double:
		return reflect.TypeOf(float64(1))
	case core.Char, core.Varchar, core.NVarchar, core.TinyText, core.Text, core.NText, core.MediumText, core.LongText, core.Enum, core.Set, core.Uuid, core.Clob, core.SysName:
		return reflect.TypeOf("")
	case core.TinyBlob, core.Blob, core.LongBlob, core.Bytea, core.Binary, core.MediumBlob, core.VarBinary, core.UniqueIdentifier:
		return reflect.TypeOf([]byte{})
	case core.Bool:
		return reflect.TypeOf(true)
	case core.DateTime, core.Date, core.Time, core.TimeStamp, core.TimeStampz:
		return reflect.TypeOf(c_TIME_DEFAULT)
	case core.Decimal, core.Numeric:
		return reflect.TypeOf(float64(1))
	default:
		return reflect.TypeOf("")
	}
}
