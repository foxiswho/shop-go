package orm

import (
	"reflect"
	"time"

	"github.com/hb-go/echo-web/middleware/cache"
	"github.com/hb-go/gorm"

	"github.com/hb-go/echo-web/module/log"
	"github.com/hb-go/echo-web/util/conv"
	"github.com/hb-go/echo-web/util/crypt"
	"github.com/hb-go/echo-web/util/sql"
)

const (
	CacheExpireDefault = time.Minute
	CacheKeyFormat     = "SQL:%s SQLVars:%v"
)

type CacheDB struct {
	*gorm.DB
	store  cache.CacheStore
	Expire time.Duration
}

type CacheConf struct {
	Expire time.Duration
}

func NewCacheDB(db *gorm.DB, store cache.CacheStore, conf CacheConf) *CacheDB {
	switch conf.Expire {
	case time.Duration(0):
		conf.Expire = CacheExpireDefault
	}

	newDB := CacheDB{
		DB:     db,
		store:  store,
		Expire: conf.Expire,
	}
	return &newDB
}

func (c *CacheDB) First(out interface{}, where ...interface{}) *CacheDB {
	sql := gorm.SQL{}
	key := ""
	c.DB = c.FirstSQL(&sql, out, where...)
	if err := c.DB.Error; err != nil {
		return c
	} else {
		key = cacheKey(sql)
	}

	if err := c.store.Get(key, out); err != nil {
		log.Debugf("first no cache data")
		c.DB = c.DB.First(out, where...)
		if err := c.DB.Error; err == nil {
			c.store.Set(key, out, c.Expire)
		}
	}
	return c
}

func (c *CacheDB) Last(out interface{}, where ...interface{}) *CacheDB {
	sql := gorm.SQL{}
	key := ""
	c.DB = c.DB.LastSQL(&sql, out, where...)
	if err := c.DB.Error; err != nil {
		return c
	} else {
		key = cacheKey(sql)
	}

	if err := c.store.Get(key, out); err != nil {
		log.Debugf("last no cache data")
		c.DB = c.DB.Last(out, where...)
		if err := c.DB.Error; err == nil {
			c.store.Set(key, out, c.Expire)
		}
	}
	return c
}

func (c *CacheDB) Find(out interface{}, where ...interface{}) *CacheDB {
	sql := gorm.SQL{}
	key := ""
	c.DB = c.DB.FindSQL(&sql, out, where...)
	if err := c.DB.Error; err != nil {
		return c
	} else {
		key = cacheKey(sql)
	}

	if err := c.store.Get(key, out); err != nil {
		log.Debugf("find no cache data")
		c.DB = c.DB.Find(out, where...)
		if err := c.DB.Error; err == nil {
			c.store.Set(key, out, c.Expire)
		}
	}
	return c
}

func (c *CacheDB) Count(out interface{}) *CacheDB {
	sql := gorm.SQL{}
	key := ""
	c.DB = c.DB.CountSQL(&sql, out)
	if err := c.DB.Error; err != nil {
		return c
	} else {
		key = cacheKey(sql)
	}

	if err := c.store.Get(key, out); err != nil {
		log.Debugf("count no cache data, err:%s", err)
		c.DB = c.DB.Count(out)
		if err := c.DB.Error; err == nil {
			var value interface{}
			if v := reflect.ValueOf(out); v.Kind() == reflect.Ptr {
				p := v.Elem()
				switch p.Kind() {
				case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
					value = conv.IntPtrTo64(out)
				case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
					value = conv.UintPtrTo64(out)
				}
			}
			if err := c.store.Set(key, value, c.Expire); err != nil {
				c.DB.AddError(err)
			}
		}
	}

	return c
}

func cacheKey(gSql gorm.SQL) string {
	//sqlStr := fmt.Sprintf(CacheKeyFormat, sql.SQL, sql.SQLVars)
	sqlStr := sql.SqlParse(gSql.SQL, gSql.SQLVars)
	return crypt.MD5([]byte(sqlStr))
}
