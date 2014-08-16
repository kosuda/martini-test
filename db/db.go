package db

import (
	"github.com/garyburd/redigo/redis"
)

var conn redis.Conn

const (
	address = "127.0.0.1:6379"
)

func init() {
	conn, _ = redis.Dial("tcp", address)
}

// Read function
func Read(key string, data interface{}) error {
	res, _ := redis.Values(conn.Do("HGETALL", key))

	if err := redis.ScanStruct(res, data); err != nil {
		return err
	}

	return nil
}

// Write fucntion
func Write(key string, data interface{}) error {

	if _, err := conn.Do("HMSET", redis.Args{}.Add(key).AddFlat(data)...); err != nil {
		return err
	}

	return nil
}
