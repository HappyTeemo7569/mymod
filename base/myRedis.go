package base

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"time"
)

//Redis连接池
var RedisPool redis.Pool

//获取RPC Redis连接
func GetRedisRpc() redis.Conn {
	rc := RedisPool.Get()
	rc.Do("SELECT", configBase.Redis.Rpc)
	return rc
}

//获取API Redis连接
func GetRedisApi() redis.Conn {
	rc := RedisPool.Get()
	rc.Do("SELECT", configBase.Redis.Api)
	return rc
}

func InitRedis() {
	//连接池
	RedisPool = redis.Pool{
		Dial: func() (redis.Conn, error) {
			connstr := fmt.Sprintf("%s:%d", configBase.Redis.Addr, configBase.Redis.Port)
			con, err := redis.Dial("tcp", connstr)
			if err != nil {
				//panic(err)
				return nil, err
			}
			if len(configBase.Redis.Pwd) > 0 {
				err = con.Send("auth", configBase.Redis.Pwd)
				if err != nil {
					//panic(err)
					return nil, err
				}
			}
			if _, err := con.Do("SELECT", configBase.Redis.Rpc); err != nil {
				//panic(err)
				return nil, err
			}
			return con, nil
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			if time.Since(t) < time.Minute {
				return nil
			}
			_, err := c.Do("PING")
			return err
		},
		MaxIdle:         3000,
		MaxActive:       5000,
		IdleTimeout:     240 * time.Second,
		Wait:            false,
		MaxConnLifetime: 0,
	}
}
