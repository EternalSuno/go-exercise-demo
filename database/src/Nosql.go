package src

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var (
	Pool *redis.Pool
)

func init() {
	redisHost := ":6379"
	Pool = newPool(redisHost)
	close()
}

func newPool(server string) *redis.Pool {
	return &redis.Pool{
		MaxIdle:     3,
		IdleTimeout: 240 * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", server)
			if err != nil {
				return nil, err
			}
			return c, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}
}

func close() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, syscall.SIGTERM)
	signal.Notify(c, syscall.SIGKILL)
	go func() {
		<-c
		Pool.Close()
		os.Exit(0)
	}()
}

func Get(key string) ([]byte, error) {
	conn := Pool.Get()
	defer conn.Close()

	var data []byte
	data, err := redis.Bytes(conn.Do("GET", key))
	if err != nil {
		return data, fmt.Errorf("error get key %s: %v", key, err)
	}
	return data, err
}

func nosql() {
	test, err := Get("test")
	fmt.Println(test, err)
	//非关系型数据库
	//目前流行的 NOSQL 主要有 redis、mongoDB、Cassandra 和 Membase 等

	//redis
	//redis 是一个 key-value 存储系统。
	//和 Memcached 类似，它支持存储的 value 类型相对更多，
	//包括 string (字符串)、list (链表)、set (集合 ) 和 zset (有序集合)。
	//

	//Go 目前支持 redis 的驱动有如下
	//
	//github.com/gomodule/redigo (推荐)
	//github.com/go-redis/redis
	//github.com/hoisie/redis
	//github.com/alphazero/Go-Redis
	//github.com/simonz05/godis
	//

	//
	//package main
	//
	//import (
	//	"fmt"
	//
	//"github.com/astaxie/goredis"
	//)
	//
	//func main() {
	//	var client goredis.Client
	//	设置端口为 redis 默认端口
	//client.Addr = "127.0.0.1:6379"
	//
	//字符串操作
	//client.Set("a", []byte("hello"))
	//val, _ := client.Get("a")
	//fmt.Println(string(val))
	//client.Del("a")

	// list 操作
	//vals := []string{"a", "b", "c", "d", "e"}
	//for _, v := range vals {
	//	client.Rpush("l", []byte(v))
	//}
	//dbvals,_ := client.Lrange("l", 0, 4)
	//for i, v := range dbvals {
	//	println(i,":",string(v))
	//}
	//client.Del("l")
	//}

}
