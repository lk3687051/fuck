package stock
import (
  "fmt"
  "encoding/json"
)

type Statistics struct {
  Daily  time.Time  `json:"date"`
  UpPencent  float  `json:"upPencent"`
  DownPencent  float  `json:"downPencent"`
}

type Pool struct {
    Category  string    `json:"category"`
    Name      string    `json:"Name"`
    Stats      map[string]Statistics      `json:"stat"`
  	Stocks    []string
}

func (p *Pool)Save()  {
  key := fmt.Sprintf("pool:%s:%s:info", p.Category, p.Name)
  data, _ := json.Marshal(p)
  err := rdb.Set(key, string(data), 0).Err()
  if err != nil {
      panic(err)
  }
}

func NewPool(category string, name string) *Pool{
  p := new(Pool)
  p.Category = category
  p.Name = name
  key := fmt.Sprintf("pool:%s:%s:stocks", category, name)
  p.Stocks, _ = rdb.SMembers(key).Result()
  return p
}

func AddStockToPool(category string, name string, code string) {
  rdb.SAdd("pools:"+category, name)
  key := fmt.Sprintf("pool:%s:%s:stocks", category, name)
  rdb.SAdd(key, code)
}

func GetPoolByCategory(category string) []*Pool {
  p_list := []*Pool{}
  key := fmt.Sprintf("pools:%s", category)
  pools, _ := rdb.SMembers(key).Result()
  for _, pool := range pools {
    p := NewPool(category, pool)
    p_list = append(p_list, p)
  }
  return p_list
}
