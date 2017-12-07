package redisClient

import "fmt"

type RedisManager struct {
	RedisClientMain   *RedisClient
	RedisClientBackup *RedisClient
}

func (rm *RedisManager) SetRedisServer(address, password string, isMainServer bool) {

	var c RedisClient
	err := c.ConnectTo(address, password)
	if err != nil {
		fmt.Println("connect failed:", address, err)
	}
	if isMainServer {
		rm.RedisClientMain = &c
		c.EnableAutoReconnect()
	} else {
		rm.RedisClientBackup = &c
	}

}

func (rm *RedisManager) Close() {
	if rm.RedisClientBackup != nil {
		rm.RedisClientBackup.Close()
	}
	if rm.RedisClientMain != nil {
		rm.RedisClientMain.Close()
	}
}

func (rm *RedisManager) GetString(key string) (string, error) {

	if rm.RedisClientMain != nil && rm.RedisClientMain.IsAvaliable {
		value, err := rm.RedisClientMain.GetString(key)
		if err != nil && rm.RedisClientBackup != nil && rm.RedisClientBackup.IsAvaliable {
			return rm.RedisClientBackup.GetString(key)
		} else {
			return value, nil
		}
	} else if rm.RedisClientBackup != nil && rm.RedisClientBackup.IsAvaliable {
		return rm.RedisClientBackup.GetString(key)
	} else {
		return "", fmt.Errorf("no redis avaliable")
	}
}
