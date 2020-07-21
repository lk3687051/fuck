package fuck
import (
  // "time"
  // "fmt"
  "encoding/json"
  log "github.com/sirupsen/logrus"
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

func (p *Pool)AddStock(ts_code string)  {
  log.Debugf("Add Stock %s to Pool %s:%s\n", ts_code, p.Category, p.Name)
  p.Stocks = append(p.Stocks, ts_code)
}

func (p *Pool)Debug()  {
  log.Debugf("The Pools %s:%s have stocks %d stocks\n", p.Category, p.Name, len(p.Stocks))
}

func AddStock(category string, name string, code string) {
  key := category + ":" + name
  p, ok := PoolMap.Pools[key]
  if !ok {
    _p := Pool{
      Category: category,
      Name: name,
      Stocks: make([]string,0),
    }
    _p.Stocks = append(_p.Stocks, code)
    PoolMap.Pools[key] = _p

  } else {
    p.Stocks = append(p.Stocks, code)
    PoolMap.Pools[key] = p
  }
}

func LoadPools()  {
  PoolStr := LoadResource(PoolResource, "")
  if len(PoolStr) != 0 {
    json.Unmarshal(PoolStr, &PoolMap)
    return
  }
}
func SetupPools()  {
  stocks := GetStocks()
  for _, s := range stocks{
    AddStock("Area", s.Area, s.TsCode)
    AddStock("Industry", s.Industry, s.TsCode)
    AddStock("Exchange", s.Exchange, s.TsCode)
  }
  data, _ := json.MarshalIndent(PoolMap, "", " ")

  SaveResource(PoolResource, "", data)
}
