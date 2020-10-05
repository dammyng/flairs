package redis

import (
	"log"
	"github.com/gomodule/redigo/redis"
)

//NewPool creates a new connection to redis server
func NewPool(port string) *redis.Pool {
	return &redis.Pool{
		MaxIdle:80,
		MaxActive:100,
		Dial:func () (redis.Conn, error)  {
			c, err := redis.Dial("tcp", port, redis.DialPassword("adminadmin"))
			if err != nil {
				log.Fatalf("%v", err.Error())
			}
			return c, err
		},
	}
}