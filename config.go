package fuck
import (
  "time"
  "github.com/ShawnRong/tushare-go"
)
var DataDir = "D:\\fuck"
var c = tushare.New("6d8cb43818f3724d89d561797bd6b37a9a3555c109160aa51ed7428e")
var dailyStart = time.Date(2015, 1, 1, 0, 0, 0, 0, time.UTC)
