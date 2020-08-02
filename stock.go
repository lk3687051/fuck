package fuck
import (
	"fmt"
	"time"
	"encoding/json"
	log "github.com/sirupsen/logrus"
)

type Stock struct {
	TsCode    string  `json:"ts_code"`
	Name      string  `json:"name"`
	Area      string  `json:"area"`
	Industry   string  `json:"industry"`
	ListDate    time.Time  `json:"list_date"`
	Exchange    string     `json:"exchange"`
	UpdateAt  time.Time  `json:"updateat"`
  // DailyQuote   Quote
}

func NewStock(ts_code string) *Stock {
	key := "stock:" + ts_code
	_s := Stock{}
	val, err := rdb.Get(key).Result()
	if err != nil {
		return &_s
	}
	err = json.Unmarshal([]byte(val) , &_s)
	if err != nil {
		fmt.Println("error:", err)
	}
	return &_s
}

func (s *Stock) Save()  {
	data, _ := json.Marshal(s)
	err := rdb.Set("stock:" + s.TsCode, string(data), 0).Err()
	if err != nil {
			panic(err)
	}
}

func (s *Stock)UpdateQuote()  {
	// _q := GetDaily(s.TsCode, s.UpdateAt.AddDate(0, 0, 1))
	// s.DailyQuote.Date = append(s.DailyQuote.Date, _q.Date...)
	// s.DailyQuote.Open = append(s.DailyQuote.Open, _q.Open...)
	// s.DailyQuote.High = append(s.DailyQuote.High, _q.High...)
	// s.DailyQuote.Low = append(s.DailyQuote.Low, _q.Low...)
	// s.DailyQuote.Close = append(s.DailyQuote.Close, _q.Close...)
	// s.DailyQuote.Volume = append(s.DailyQuote.Volume, _q.Volume...)
	// s.DailyQuote.PreClose = append(s.DailyQuote.PreClose, _q.PreClose...)
	// s.DailyQuote.Change = append(s.DailyQuote.Change, _q.Change...)
	// s.DailyQuote.PctChg = append(s.DailyQuote.PctChg, _q.PctChg...)
	// s.DailyQuote.Amount = append(s.DailyQuote.Amount, _q.Amount...)
	// s.Save()
}

func GetStockByCodes(ts_codes []string) []Stock {
	stocks := make([]Stock, len(ts_codes))
	for i, ts_code := range ts_codes {
		stocks[i] = *NewStock(ts_code)
	}
	log.Infof("%v\n", stocks)
	return stocks
}

func GetStockList()  []*Stock{
	ts_codes, _ := rdb.SMembers("stocks").Result()
	stocks := make([]*Stock, len(ts_codes))
	for i,ts_code := range ts_codes {
		stocks[i] = NewStock(ts_code)
		log.Infof("%+v\n", stocks[i])
	}

	return stocks
}
func SetupStocks()  {
	ss := GetStocks()
	for _,s := range ss {
		rdb.SAdd("stocks", s.TsCode)
		s.Save()
	}
}
