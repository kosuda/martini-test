package db

import (
	"github.com/garyburd/redigo/redis"
	"net/url"
	"os"
)

var conn redis.Conn

const (
	address = "127.0.0.1:6379"
)

func init() {

	if s := os.Getenv("REDISTOGO_URL"); s != "" {
		u, _ := url.Parse(s)
		conn, _ = redis.Dial("tcp", u.Host)
		p, _ := u.User.Password()
		conn.Do("AUTH", p)

	} else if s := os.Getenv("WERCKER_REDIS_HOST"); s != "" {
		conn, _ = redis.Dial("tcp", s+":6379")

	} else {
		conn, _ = redis.Dial("tcp", address)

	}

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
