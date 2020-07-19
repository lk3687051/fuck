package fuck
import (
	"time"
	"encoding/json"
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
	if IsSameDay(s.UpdateAt, time.Now()) {
		return
	}
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
	s.DailyQuote.CalcMA()
	s.UpdateAt = time.Now()
	s.Save()
}

func SetupStocks()  {
	ss := GetStocks()
	for _,s := range ss {
		stockStr := LoadResource(StockResource, s.TsCode)
		if len(stockStr) != 0 {
			json.Unmarshal(stockStr, &s)
		}
		s.UpdateQuote()
		stockMap[s.TsCode] = &s
	}
}
