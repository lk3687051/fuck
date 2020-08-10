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
	DailyQuote  *Quote
}

func NewStock(ts_code string) *Stock {
	key := "stock:" + ts_code + ":info"
	_s := Stock{}
	val, err := rdb.Get(key).Result()
	if err != nil {
		return nil
	}
	err = json.Unmarshal([]byte(val) , &_s)
	if err != nil {
		fmt.Println("error:", err)
	}
	_s.DailyQuote =  NewQuote(ts_code)
	return &_s
}

func (s *Stock) Save()  {
	key := "stock:" + s.TsCode + ":info"
	data, _ := json.Marshal(s)
	err := rdb.Set(key, string(data), 0).Err()
	if err != nil {
			panic(err)
	}
	if s.DailyQuote != nil {
		s.DailyQuote.Save()
	}
}

func GetStockByCodes(ts_codes []string) []Stock {
	stocks := make([]Stock, len(ts_codes))
	for i, ts_code := range ts_codes {
		stocks[i] = *NewStock(ts_code)
	}
	log.Infof("%v\n", stocks)
	return stocks
}

func GetAllTsCodes() []string {
	ts_codes, _ := rdb.SMembers("stocks").Result()
	return ts_codes
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
