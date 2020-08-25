package lock

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"testing"
)

func TestRedisLockMgr(t *testing.T) {
	conn, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println(err)
		return
	}
	lockMgr := NewRedisLockMgr(conn)
	err = lockMgr.Lock("tk", "tv", 30)
	if err != nil {
		fmt.Println("lock err ", err)
		return
	}
}
