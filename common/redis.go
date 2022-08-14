package common

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"time"
)

var rdb *redis.Client
var ctx context.Context = context.Background()

func init() {
	rdb = redis.NewClient(&redis.Options{
        Addr:     "127.0.0.1:6379",
        Password: "",
        DB:       0,
        PoolSize: 20,
		MinIdleConns: 10,
		MaxRetries:      3,
    })
	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		fmt.Printf("connect redis success....%v\n", err)
		panic("connect redis error")
	}
	fmt.Println("connect redis success....")
}

func Get(key string) (string,error) {
    val, err := rdb.Get(ctx, key).Result()
    if err != nil {
        fmt.Printf("Get key:%v get  err:%v", key, err)
        return "", err
    }
	fmt.Printf("val: %v\n", val)
	return val, nil
}

func Decr(key string) (int64,error) {
	val, err := rdb.Decr(ctx, key).Result()
	if err != nil {
		fmt.Printf("Decr key:%v get  err:%v", key, err)
		return 0, err
	}
	fmt.Printf("the latest val: %v\n", val)
	return val, nil
}

func Incr(key string) (int64,error) {
	val, err := rdb.Incr(ctx, key).Result()
	if err != nil {
		fmt.Printf("Incr key:%v get  err:%v", key, err)
		return 0, err
	}
	fmt.Printf("Incr the latest val: %v\n", val)
	return val, nil
}

func SetInt32(key string, val int32) error {
	err := rdb.Set(ctx, key, val, 0).Err()
	if err != nil {
		fmt.Printf("Set key:%v get  err:%v", key, err)
		return err
	}
	return nil
}

func Pop(key string) (string, error) {
	val, err := rdb.RPop(ctx, key).Result()
	if err != nil {
		fmt.Printf("Pop key:%v get  err:%v", key, err)
		return "", err
	}
	fmt.Printf("the pop list val: %v\n", val)
	return val, nil
}

func LPush(key string, val string) error {
	listLen, err := rdb.LPush(ctx, key, val).Result()

	if err != nil {
		fmt.Printf("LPush key:%v get  err:%v", key, err)
		return err
	}
	fmt.Printf("the pop list len: %v\n", listLen)
	return nil
}

func LLen(key string) (int64,error) {
	val, err := rdb.LLen(ctx, key).Result()
	if err != nil {
		fmt.Printf("LLen key:%v get  err:%v", key, err)
		return 0, err
	}
	fmt.Printf("LLen val: %v\n", val)
	return val, nil
}

func SetNX(key string, val string, expire time.Duration) (bool,error) {
	setRes, err := rdb.SetNX(ctx, key, val, expire).Result()
	if err != nil {
		fmt.Printf("SetNX key:%v get  err:%v", key, err)
		return setRes, err
	}
	return setRes, nil
}

func Del(key string) error {
	err := rdb.Del(ctx, key).Err()
	if err != nil {
		fmt.Printf("Del key:%v get  err:%v", key, err)
		return err
	}
	return nil
}

func Exec(script string, key []string, arg []string) error {
	_, err := rdb.Eval(ctx, script, key, arg).Result()
	if err != nil {
		fmt.Printf("exec key:%v get  err:%v", key, err)
		return err
	}
	return nil
}

