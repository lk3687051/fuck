package fuck
import (
  "time"
)

type Manager struct {

}

func  (m *Manager) NewManager()  {
  return Manager{}
}

func SetupData()  {
  stocks := GetStocksByTushare()
	for _,s := range stocks {
		rdb.SAdd("stocks", s.TsCode)
		s.Save()
	}
  SetupPools()
  SetupStocks()
  SetupQuotes()
}
