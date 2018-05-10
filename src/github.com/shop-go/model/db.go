package model

import (
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/hb-go/echo-web/middleware/cache"
	"github.com/hb-go/gorm"

	. "github.com/hb-go/echo-web/conf"
	"github.com/hb-go/echo-web/model/orm"
	"github.com/hb-go/echo-web/module/log"
)

var db *gorm.DB
var dbCacheStore cache.CacheStore

func DB() *gorm.DB {
	if db == nil {
		log.Debugf("Model NewDB")

		newDb, err := newDB()
		if err != nil {
			panic(err)
		}
		newDb.DB().SetMaxIdleConns(10)
		newDb.DB().SetMaxOpenConns(100)

		newDb.SetLogger(orm.Logger{})
		newDb.LogMode(true)
		db = newDb
	}

	return db
}

func newDB() (*gorm.DB, error) {
	sqlConnection := Conf.DB.UserName + ":" + Conf.DB.Pwd + "@tcp(" + Conf.DB.Host + ":" + Conf.DB.Port + ")/" + Conf.DB.Name + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open("mysql", sqlConnection)
	if err != nil {
		return nil, err
	}
	return db, nil
}

func CacheStore() cache.CacheStore {
	if dbCacheStore == nil {
		switch Conf.CacheStore {
		case MEMCACHED:
			dbCacheStore = cache.NewMemcachedStore([]string{Conf.Memcached.Server}, time.Hour)
		case REDIS:
			dbCacheStore = cache.NewRedisCache(Conf.Redis.Server, Conf.Redis.Pwd, time.Hour)
		default:
			dbCacheStore = cache.NewInMemoryStore(time.Hour)
		}
	}

	return dbCacheStore
}

func Cache(db *gorm.DB) *orm.CacheDB {
	return orm.NewCacheDB(db, CacheStore(), orm.CacheConf{
		Expire: time.Second * 10,
	})
}
