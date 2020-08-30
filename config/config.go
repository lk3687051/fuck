package config
import (
  "github.com/ShawnRong/tushare-go"
  "github.com/go-redis/redis"
)
var C = tushare.New("6d8cb43818f3724d89d561797bd6b37a9a3555c109160aa51ed7428e")

var Rdb = redis.NewClient(&redis.Options{
        Addr:     "localhost:6379",
        Password: "", // no password set
        DB:       0,  // use default DB
    })
