package fuck
import (
  "time"
  "github.com/ShawnRong/tushare-go"
  "context"
  "github.com/go-redis/redis"
)
var c = tushare.New("6d8cb43818f3724d89d561797bd6b37a9a3555c109160aa51ed7428e")
var dailyStart = time.Date(2015, 7, 1, 0, 0, 0, 0, time.UTC)
var ctx = context.Background()
var rdb = redis.NewClient(&redis.Options{
        Addr:     "localhost:6379",
        Password: "", // no password set
        DB:       0,  // use default DB
    })
