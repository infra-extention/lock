package lock

type LockMgr interface {
	Lock(key, value string, ttl int) error
	Unlock() error
	ExpireLongTimeLock() error
}
