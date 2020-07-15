package main
import (
  "fmt"
  "fuck/market"
)

func main()  {
  q := market.NewQuote("000001.SZ")
  fmt.Print(q)
}
