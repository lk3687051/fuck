package main
import (
  "fmt"
  "time"
  "fuck/stock"
)

func main()  {
  startT := time.Now()
  stock.GetAllStock()
  tc := time.Since(startT)	//计算耗时
  fmt.Printf("time cost = %v\n", tc)
}
