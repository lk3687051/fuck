package fuck
import (
  log "github.com/sirupsen/logrus"
)

type Manager struct {
  Stocklist []string
  Stocks map[string]*Stock
  Pools []*Pool
}

var GlobalManager = new(Manager)

func (m *Manager) Init()  {
  log.Info("Start Loading stock data")
	m.Stocklist = GetAllTsCodes()
  m.Stocks = map[string]*Stock{}
  m.Pools = []*Pool{}
  for _, ts_code := range m.Stocklist {
    s := NewStock(ts_code)
    m.Stocks[ts_code] = s
  }
  log.Info("Start Loading pool data")
  p_types := []string {"Area", "Industry", "Exchange"}
  for _, p_type := range p_types {
    p_names ,_ := rdb.SMembers("pools:" + p_type).Result()
    log.Infof("%+v", p_names)
    for _, p_name := range p_names {
      p := NewPool(p_type, p_name)
      m.Pools = append(m.Pools, p)
    }
  }
  log.Info("End Loading data")
}

func (m *Manager) Worker()  {
  for _, p := range m.Pools {
    p.Statistics()
    p.Save()
  }
}

func SetupData()  {
	ss := GetStocksByTushare()
	for _,s := range ss {
		rdb.SAdd("stocks", s.TsCode)
		s.Save()
    AddStockToPool("Area", s.Area, s.TsCode)
    AddStockToPool("Industry", s.Industry, s.TsCode)
    AddStockToPool("Exchange", s.Exchange, s.TsCode)
    NewDailyQuoteFromWeb(s.TsCode)
	}
}
