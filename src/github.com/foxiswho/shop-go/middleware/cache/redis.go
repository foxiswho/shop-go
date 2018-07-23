package cache

import (
	"github.com/gomodule/redigo/redis"
	"time"
	"fmt"
)

// Wraps the Redis client to meet the Cache interface.
type RedisStore struct {
	pool              *redis.Pool
	defaultExpiration time.Duration
}

// until redigo supports sharding/clustering, only one host will be in hostList
func NewRedisCache(host string, password string, defaultExpiration time.Duration) *RedisStore {
	var pool = &redis.Pool{
		MaxIdle:     10,
		MaxActive:   512,
		Wait:        false,
		IdleTimeout: 3 * time.Second,
		Dial: func() (redis.Conn, error) {
			// the redis protocol should probably be made sett-able
			c, err := redis.Dial("tcp", host)
			if err != nil {
				return nil, err
			}
			if len(password) > 0 {
				if _, err := c.Do("AUTH", password); err != nil {
					c.Close()
					return nil, err
				}
				//if _, err := c.Do("SELECT",1); err != nil {
				// c.Close()
				// return nil, err
				//}
			} else {
				// check with PING
				if _, err := c.Do("PING"); err != nil {
					c.Close()
					return nil, err
				}
			}
			return c, err
		},
		// custom connection test method
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			if _, err := c.Do("PING"); err != nil {
				return err
			}
			return nil
		},
	}
	return &RedisStore{pool, defaultExpiration}
}

func (c *RedisStore) Set(key string, value interface{}, expires time.Duration) error {
	return c.invoke(c.pool.Get().Do, key, value, expires)
}

func (c *RedisStore) Add(key string, value interface{}, expires time.Duration) error {
	conn := c.pool.Get()
	if exists(conn, key) {
		return ErrNotStored
	}
	return c.invoke(conn.Do, key, value, expires)
}

func (c *RedisStore) Replace(key string, value interface{}, expires time.Duration) error {
	conn := c.pool.Get()
	if !exists(conn, key) {
		return ErrNotStored
	}
	err := c.invoke(conn.Do, key, value, expires)
	if value == nil {
		return ErrNotStored
	} else {
		return err
	}
}

func (c *RedisStore) Get(key string, ptrValue interface{}) error {
	conn := c.pool.Get()
	defer conn.Close()
	raw, err := conn.Do("GET", key)
	if raw == nil {
		return ErrCacheMiss
	}
	item, err := redis.Bytes(raw, err)
	if err != nil {
		return err
	}
	return deserialize(item, ptrValue)
}

func exists(conn redis.Conn, key string) bool {
	retval, _ := redis.Bool(conn.Do("EXISTS", key))
	return retval
}

func (c *RedisStore) Delete(key string) error {
	conn := c.pool.Get()
	defer conn.Close()
	if !exists(conn, key) {
		return ErrCacheMiss
	}
	_, err := conn.Do("DEL", key)
	return err
}

func (c *RedisStore) Increment(key string, delta uint64) (uint64, error) {
	conn := c.pool.Get()
	defer conn.Close()
	// Check for existance *before* increment as per the cache contract.
	// redis will auto create the key, and we don't want that. Since we need to do increment
	// ourselves instead of natively via INCRBY (redis doesn't support wrapping), we get the value
	// and do the exists check this way to minimize calls to Redis
	val, err := conn.Do("GET", key)
	if val == nil {
		return 0, ErrCacheMiss
	}
	if err == nil {
		currentVal, err := redis.Int64(val, nil)
		if err != nil {
			return 0, err
		}
		var sum int64 = currentVal + int64(delta)
		_, err = conn.Do("SET", key, sum)
		if err != nil {
			return 0, err
		}
		return uint64(sum), nil
	} else {
		return 0, err
	}
}

func (c *RedisStore) Decrement(key string, delta uint64) (newValue uint64, err error) {
	conn := c.pool.Get()
	defer conn.Close()
	// Check for existance *before* increment as per the cache contract.
	// redis will auto create the key, and we don't want that, hence the exists call
	if !exists(conn, key) {
		return 0, ErrCacheMiss
	}
	// Decrement contract says you can only go to 0
	// so we go fetch the value and if the delta is greater than the amount,
	// 0 out the value
	currentVal, err := redis.Int64(conn.Do("GET", key))
	if err == nil && delta > uint64(currentVal) {
		tempint, err := redis.Int64(conn.Do("DECRBY", key, currentVal))
		return uint64(tempint), err
	}
	tempint, err := redis.Int64(conn.Do("DECRBY", key, delta))
	return uint64(tempint), err
}

func (c *RedisStore) Flush() error {
	conn := c.pool.Get()
	defer conn.Close()
	_, err := conn.Do("FLUSHALL")
	return err
}

func (c *RedisStore) invoke(f func(string, ...interface{}) (interface{}, error),
	key string, value interface{}, expires time.Duration) error {

	switch expires {
	case DEFAULT:
		expires = c.defaultExpiration
	case FOREVER:
		expires = time.Duration(0)
	}

	b, err := serialize(value)
	if err != nil {
		return err
	}
	conn := c.pool.Get()
	defer conn.Close()
	if expires > 0 {
		_, err := f("SETEX", key, int32(expires/time.Second), b)
		return err
	} else {
		_, err := f("SET", key, b)
		return err
	}
}

func (c *RedisStore) invokeHash(f func(string, ...interface{}) (interface{}, error),
	key string, field string, value interface{}, expires time.Duration) error {

	switch expires {
	case DEFAULT:
		expires = c.defaultExpiration
	case FOREVER:
		expires = time.Duration(0)
	}

	b, err := serialize(value)
	if err != nil {
		return err
	}
	conn := c.pool.Get()
	defer conn.Close()
	_, err = f("hset", key, field, b)
	if err != nil {
		return err
	}
	if expires > 0 {
		// 设置过期时间为24小时
		_, err = f("EXPIRE", key, expires)
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *RedisStore) HGet(key string, field string, ptrValue interface{}) error {
	conn := c.pool.Get()
	defer conn.Close()
	res, err := conn.Do("hget", key, field)
	if err != nil {
		return err
	}
	item, err := redis.Bytes(res, err)
	return deserialize(item, ptrValue)
}

func (c *RedisStore) HSet(key string, field string, value interface{}, expires time.Duration) error {
	return c.invokeHash(c.pool.Get().Do, key, field, value, expires)
}

func (c *RedisStore) HExists(key string, field string) bool {
	conn := c.pool.Get()
	defer conn.Close()
	isExist, err := redis.Bool(conn.Do("hexists", key, field))
	if err != nil {
		return false
	}
	return isExist
}

func (c *RedisStore) HLen(key string) int {
	conn := c.pool.Get()
	defer conn.Close()
	ilen, err := redis.Int(conn.Do("hlen", key))
	if err != nil {
		return 0
	}
	return ilen
}

func (c *RedisStore) HKeys(key string) ([]string, error) {
	arr := []string{}
	conn := c.pool.Get()
	defer conn.Close()
	res, err := conn.Do("hkeys", key)
	if err != nil {
		return arr, err
	}
	resKeys, err := redis.Strings(res, err)
	if err != nil {
		return arr, err
	}
	return resKeys, nil
}
func (c *RedisStore) HVals(key string) ([]interface{}, error) {
	arr := []interface{}{}
	conn := c.pool.Get()
	defer conn.Close()
	res, err := redis.Values(conn.Do("hvals", key))
	if err != nil {
		return arr, err
	}
	return res, nil
}

func (c *RedisStore) HDel(key string, field string) bool {
	conn := c.pool.Get()
	defer conn.Close()
	isExist, err := redis.Bool(conn.Do("HDEL", key, field))
	if err != nil {
		return false
	}
	return isExist
}

//获取多个值
func (c *RedisStore) HGetAll(key string) ([]interface{}, error) {
	conn := c.pool.Get()
	defer conn.Close()
	result, err := redis.Values(conn.Do("hgetall", key))
	if err != nil {
		return nil, err
	} else {
		arr := []interface{}{}
		fmt.Println("hgetall",result)
		for _, v := range result {
			fmt.Printf("hgetall %s ", v.([]byte))
			arr = append(arr, v)
		}
		return arr, nil
	}
}

//设置多个值
func (c *RedisStore) HMSet(fvs map[string]interface{}) (bool, error) {
	conn := c.pool.Get()
	defer conn.Close()
	args := []interface{}{}
	for key, value := range fvs {
		args = append(args, key, value)
	}
	_, err := conn.Do("HMSET", args)
	if err != nil {
		return false, err
	} else {
		return true, nil
	}
}

//获取多个值
func (c *RedisStore) HMGet(keys []string) ([]interface{}, error) {
	conn := c.pool.Get()
	defer conn.Close()
	result, err := redis.Values(conn.Do("hmget", keys))
	if err != nil {
		return nil, err
	} else {
		arr := []interface{}{}
		for _, v := range result {
			arr = append(arr, v)
		}
		return arr, nil
	}
}
