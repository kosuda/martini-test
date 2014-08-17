package db

import (
	"fmt"
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
		var err error
		fmt.Println(s)
		u, _ := url.Parse(s)
		conn, err = redis.Dial("tcp", u.Host)
		if err != nil {
			fmt.Println(err.Error())
		}
		p, _ := u.User.Password()
		fmt.Println(p)
		_, err = conn.Do("AUTH", p)
		if err != nil {
			fmt.Println(err.Error())
		}
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
