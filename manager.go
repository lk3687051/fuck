package fuck
import (
  log "github.com/sirupsen/logrus"
)

type Manager struct {
  Stocks []string
  StockInfo map[string]*Stock
  Pool map[string]Pool
}

func NewManager() Manager {
  m := Manager{}
  m.StockInfo = map[string]*Stock{}
  return m
}

func (m *Manager) Load()  {
  log.Info("Start Loading data")
	m.Stocks = GetAllTsCodes()
  for _, ts_code := range m.Stocks {
    s := NewStock(ts_code)
    m.StockInfo[ts_code] = s
    // log.Infof("%v", s)
  }
  log.Info("End Loading data")
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
