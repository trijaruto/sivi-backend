package connection

import (
	"github.com/go-redis/redis"
)

func NewConnectionRedis(result *redis.Client, options *redis.Options) (message string, err error) {
	*result = *redis.NewClient(options)
	pong, err := result.Ping().Result()
	if err != nil {
		return "", err
	}
	return pong, nil
}

type Redis struct {
	ListRedis map[string] *redis.Client
}

func (r *Redis)NewRedisMultipleConnection(name []string,options ...*redis.Options) (error){
	m := make(map[string]*redis.Client)
	for i, opt := range options{
		result :=redis.NewClient(opt)
		_, err :=result.Ping().Result()
		if err != nil {
			panic(err)
		}
		m[name[i]]=result
	}
	r.ListRedis =m
	return nil
}
