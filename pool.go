package fuck
import (
  // "time"
  "encoding/json"
  log "github.com/sirupsen/logrus"
)

type Stat struct {
    mean  float64
}

var PoolMap = make(map[string]*Pool)
type Pool struct {
    Category  string    `json:"category"`
    Name      string    `json:"Name"`
    NUms      int       `json:"Nums"`
    Stat      Stat      `json:"stat"`
  	Stocks    []string  `json:"stocks"`
}

func NewPool(category string, name string) *Pool{
  p :=  new(Pool)
  p.Category = category
  p.Name = name
  p.Stocks = make([]string,0)
  return p
}

func (p *Pool)AddStock(ts_code string)  {
  log.Debugf("Add Stock %s to Pool %s:%s\n", ts_code, p.Category, p.Name)
  p.Stocks = append(p.Stocks, ts_code)
}

func (p *Pool)Debug()  {
  log.Debugf("The Pools %s:%s have stocks %d stocks\n", p.Category, p.Name, len(p.Stocks))
}

func (p *Pool)Statistics()  {
  // today := time.Now().Format("20060102")
  // sum := 0.0
  // num := 0
  // for _, q := range p.Quotes {
  //   i, ok := q.GetIndex(today)
  //   if ok {
  //     sum += q.PctChg[i]
  //     num ++
  //   }
  // }
  // p.Stat.mean = sum / float64(num)
}

func GetPool(category string, name string)  *Pool {
  key := category + ":" + name
  p, ok := PoolMap[key]
  if !ok {
    _p := NewPool(category, name)
    PoolMap[key] = _p
    return _p
  } else {
    return p
  }
}

func SetupPools()  {
  PoolStr := LoadResource(PoolResource, "")
  if len(PoolStr) != 0 {
    json.Unmarshal(PoolStr, &PoolMap)
    return
  }
  stocks := GetStocks()
  for _, s := range stocks{
    p := GetPool("Area", s.Area)
    p.Stocks = append(p.Stocks, s.TsCode)
    p = GetPool("Industry", s.Industry)
    p.Stocks = append(p.Stocks, s.TsCode)
    p = GetPool("Exchange", s.Exchange)
    p.AddStock(s.TsCode)
  }
  data, _ := json.MarshalIndent(PoolMap, "", " ")
  SaveResource(PoolResource, "", data)
}
