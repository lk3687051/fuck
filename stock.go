package fuck
import (
	"time"
	"encoding/json"
	log "github.com/sirupsen/logrus"
)

var stockMap = make(map[string]*Stock)
type Stock struct {
	TsCode    string  `json:"ts_code"`
	Name      string  `json:"name"`
	Area      string  `json:"area"`
	Industry   string  `json:"industry"`
	ListDate    time.Time  `json:"list_date"`
	Exchange    string     `json:"exchange"`
	UpdateAt  time.Time  `json:"updateat"`
  DailyQuote   Quote `json:"dailyquote"`
}

func (s *Stock) Save()  {
	data, _ := json.MarshalIndent(s, "", " ")
	SaveResource(StockResource, s.TsCode, data)
}

func (s *Stock)UpdateQuote()  {
	// if IsSameDay(s.UpdateAt, time.Now()) {
	// 	return
	// }
	_q := GetDaily(s.TsCode, s.UpdateAt.AddDate(0, 0, 1))
	s.DailyQuote.Date = append(s.DailyQuote.Date, _q.Date...)
	s.DailyQuote.Open = append(s.DailyQuote.Open, _q.Open...)
	s.DailyQuote.High = append(s.DailyQuote.High, _q.High...)
	s.DailyQuote.Low = append(s.DailyQuote.Low, _q.Low...)
	s.DailyQuote.Close = append(s.DailyQuote.Close, _q.Close...)
	s.DailyQuote.Volume = append(s.DailyQuote.Volume, _q.Volume...)
	s.DailyQuote.PreClose = append(s.DailyQuote.PreClose, _q.PreClose...)
	s.DailyQuote.Change = append(s.DailyQuote.Change, _q.Change...)
	s.DailyQuote.PctChg = append(s.DailyQuote.PctChg, _q.PctChg...)
	s.DailyQuote.Amount = append(s.DailyQuote.Amount, _q.Amount...)
	s.Save()
}

func GetStockByCodes(ts_codes []string) []*Stock {
	stocks := make([]*Stock, len(ts_codes))
	for i, ts_code := range ts_codes {
		s,_ := stockMap[ts_code]
		stocks[i] = s
	}
	return stocks
}

func LoadStocks()  {
	ss := GetStocks()
	for _,s := range ss {
		_s := Stock{}
		log.Debugf("Load StockResource %s", s.TsCode)
		stockStr := LoadResource(StockResource, s.TsCode)
		if len(stockStr) != 0 {
			json.Unmarshal(stockStr, &_s)
		}
		stockMap[s.TsCode] = &_s
	}
}

func SetupStocks()  {
	ss := GetStocks()
	for _,s := range ss {
		s.UpdateQuote()
		stockMap[s.TsCode] = &s
	}
}
