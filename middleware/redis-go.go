package middleware

import (
	red "github.com/gomodule/redigo/redis"
	"strings"
	"time"
)

type Redis struct {
	pool *red.Pool
}

var redis *Redis

func init() {
	initRedis()
}

// 初始化连接
func initRedis() {
	redis = new(Redis)
	redis.pool = &red.Pool{
		MaxIdle:     31, // 最大空闲数
		MaxActive:   32,
		IdleTimeout: time.Duration(10) * time.Millisecond,
		Dial: func() (red.Conn, error) {
			return red.Dial(
				"tcp",
				"127.0.0.1:6379",
				red.DialReadTimeout(time.Duration(1000)*time.Millisecond),    // 读取单个命令应答的超时时间。
				red.DialWriteTimeout(time.Duration(1000)*time.Millisecond),   // 指定写单个命令的超时时间
				red.DialConnectTimeout(time.Duration(1000)*time.Millisecond), //连接Redis服务器的超时时间，如果没有指定默认30s
				red.DialDatabase(0), // DialDatabase拨号连接时选择的数据库。
				//red.DialPassword(""),
			)
		},
	}
}

func RedisExec(cmd string, key interface{}, args ...interface{}) (string, error) {
	con := redis.pool.Get()
	if err := con.Err(); err != nil {
		Log.Error(err)
		return "", err
	}
	defer con.Close()
	parmas := make([]interface{}, 0)
	parmas = append(parmas, key)
	if len(args) > 0 {
		for _, v := range args {
			parmas = append(parmas, v)
		}
	}
	if strings.EqualFold(cmd, "Set") {
		// 写入redis数据命令
		result, err := con.Do(cmd, parmas...)
		// 解析数据
		str, err := red.String(result, err)
		return str, err
	} else if strings.EqualFold(cmd, "Get") {
		//获取redis数据命令
		result, err := con.Do(cmd, parmas...)
		// 解析数据
		str, err := red.String(result, err)
		if err != nil {
			if strings.Contains(err.Error(), "redigo: nil returned") {
				return "key对应数据为空!", err
			}
		}

		return str, err
	}
	return "没有匹配到命令!", nil
}

func RedisExpire(key interface{}, day int64) (interface{}, error) {
	con := redis.pool.Get()
	if err := con.Err(); err != nil {
		return nil, err
	}
	defer con.Close()
	parmas := make([]interface{}, 0)
	parmas = append(parmas, key)
	if day > 30 {
		day = 30 * 60 * 60 * 24
	} else {
		day = 30 * 60 * 60 * 24
	}
	parmas = append(parmas, day)
	result, err := con.Do("EXPIRE", parmas...)
	// 解析数据
	str, err := red.String(result, err)
	return str, err
}
