package fuck
import (
	// "fmt"
	"time"
	"math"
	"encoding/json"
	"github.com/markcheno/go-talib"
	log "github.com/sirupsen/logrus"
)

type QuoteType int
const (
    DailyType QuoteType = 1 + iota
)

type Quote struct {
	Open      []float64   `json:"open"`
	High      []float64   `json:"high"`
	Low       []float64   `json:"low"`
	Close     []float64   `json:"close"`
	Volume    []float64   `json:"volume"`
	PreClose  []float64   `json:"pre_close"`
	Change    []float64     `json:"change"`
	PctChg    []float64     `json:"pct_chg"`
	Amount    []float64     `json:"amount"`
	MA5       []float64     `json:"ma5"`
	MA10       []float64    `json:"ma10"`
	MA20       []float64    `json:"ma20"`
	MA62       []float64    `json:"ma62"`
	UP         []bool       `json:"up"`
	IslimitUp  []bool       `json:"islimitUp"`
	IslimitDown  []bool     `json:"islimitDown"`
}

func (q *Quote) Save()  {
	key := "stock:" + q.TsCode + ":dailyquote"
	data, _ := json.Marshal(q)
	err := rdb.Set(key, string(data), 0).Err()
	if err != nil {
			panic(err)
	}
}

func NewQuote(ts_code string) *Quote {
	key := "stock:" + ts_code + ":dailyquote"
	_q := Quote{}
	val, err := rdb.Get(key).Result()
	if err != nil {
		return nil
	}
	err = json.Unmarshal([]byte(val) , &_q)
	if err != nil {
		log.Info("error:", err)
	}
	return &_q
}

func NewDailyQuoteFromWeb(ts_code string) {
	q := GetStockDailyByTushare(ts_code, time.Now())
	q.PreCalc()
	q.Save()
}

func (q *Quote) GetIndex(d string) (int, bool) {
	date,_ := time.Parse("20060102", d)
	for index, _date := range q.Date {
		if date == _date {
			return index, true
		}
	}
	return 0,false
}

func (q *Quote) PreCalc()  {
  // 计算涨停跌停
  q.IslimitUp = make([]bool, len(q.Date))
  q.IslimitDown = make([]bool, len(q.Date))
  for index, _ := range q.Date {
    limitUp := math.Floor(q.PreClose[index]*1.1*100 + 0.5)/100
    limitDown := math.Floor(q.PreClose[index]*0.9*100 + 0.5)/100
    q.IslimitUp[index] = limitUp == q.Close[index]
    q.IslimitDown[index] = limitDown == q.Close[index]
  }

  // 设置是否上涨
  q.UP = make([]bool, len(q.Date))
  for index, _ := range q.Date {
    if q.Close[index] > q.PreClose[index] {
      q.UP[index] = true
    } else {
      q.UP[index] = false
    }
  }

  // 计算MA值
	if len(q.Date) <= 63 {
		return
	}
	q.MA5 = talib.Ma(q.Close, 5, talib.SMA)
	q.MA10 = talib.Ma(q.Close, 10, talib.SMA)
	q.MA20 = talib.Ma(q.Close, 20, talib.SMA)
	q.MA62 = talib.Ma(q.Close, 62, talib.SMA)
}