package fuck
import (
	// "fmt"
	"time"
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
	MA5       []float64
	MA10       []float64
	MA20       []float64
	MA62       []float64
	UP         []bool
	IslimitUp  []bool
	IslimitDown  []bool
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
