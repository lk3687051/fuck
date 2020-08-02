package fuck
import (
  "fmt"
  // "time"
  // "encoding/json"
  // log "github.com/sirupsen/logrus"
)

type Stat struct {
    mean  float64
}

var PoolMap = Pools{
  Pools: make(map[string]Pool, 0),
}
type Pools struct {
    Pools map[string]Pool  `json:"Pools"`
}

type Pool struct {
    Category  string    `json:"category"`
    Name      string    `json:"Name"`
    NUms      int       `json:"Nums"`
    Stat      Stat      `json:"stat"`
  	Stocks    []string  `json:"stocks"`
}

func NewPool(category string, name string) Pool{
  p :=  Pool{
    Category: category,
    Name: name,
    Stocks: make([]string,0),
  }
  return p
}

func AddStockToPool(category string, name string, code string) {
  rdb.SAdd("pools:"+category, name)
  key := fmt.Sprintf("pool:%s:%s", category, name)
  rdb.SAdd(key, code)
}

func SetupPools()  {
  stocks := GetStocks()
  for _, s := range stocks{
    AddStockToPool("Area", s.Area, s.TsCode)
    AddStockToPool("Industry", s.Industry, s.TsCode)
    AddStockToPool("Exchange", s.Exchange, s.TsCode)
  }
}
