package lock

import (
	"github.com/garyburd/redigo/redis"
)

type RedisLockMgr struct {
	conn redis.Conn
}

func NewRedisLockMgr(redisConn redis.Conn) LockMgr {
	return &RedisLockMgr{
		conn: redisConn,
	}
}

func (p *RedisLockMgr) Lock(key, value string, ttl int) error {
	//res, err := redis.String(p.conn.Do("SET", key, value, "EX", ttl, "NX"))
	_, err := p.conn.Do("SET", key, value, "EX", ttl, "NX")
	if err != nil {
		return err
	}
	return nil
}

func (p *RedisLockMgr) Unlock() error {
	return nil
}

func (p *RedisLockMgr) ExpireLongTimeLock() error {
	return nil
}
