package fuck
import (
	// "fmt"
	"time"
	"github.com/markcheno/go-talib"
)

type QuoteType int
const (
    DailyType QuoteType = 1 + iota
)

type Quote struct {
	Date      []time.Time `json:"date"`
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
	MA10       []float64     `json:"ma10"`
	MA20       []float64     `json:"ma20"`
	MA62       []float64     `json:"ma62"`
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

func (q *Quote) CalcMA()  {
	if len(q.Date) <= 63 {
		return
	}
	q.MA5 = talib.Ma(q.Close, 5, talib.SMA)
	q.MA10 = talib.Ma(q.Close, 10, talib.SMA)
	q.MA20 = talib.Ma(q.Close, 20, talib.SMA)
	q.MA62 = talib.Ma(q.Close, 62, talib.SMA)
}
