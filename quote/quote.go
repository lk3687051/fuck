package market
import (
	// "fmt"
	"time"
	"path/filepath"
	"encoding/json"
	"io/ioutil"
	"github.com/markcheno/go-talib"
)
const (
	// Min1 - 1 Minute time string
	Min1 string = "60"
	// Min3 - 3 Minute time string
	Min3 string = "3m"
	// Min5 - 5 Minute time string
	Min5 string = "300"
	// Min15 - 15 Minute time string
	Min15 string = "900"
	// Min30 - 30 Minute time string
	Min30 string = "1800"
	// Min60 - 60 Minute time string
	Min60 string = "3600"
	// Hour2 - 2 hour time string
	Hour2 string = "2h"
	// Hour4 - 4 hour time string
	Hour4 string = "4h"
	// Hour6 - 6 hour time string
	Hour6 string = "6h"
	// Hour8 - 8 hour time string
	Hour8 string = "8h"
	// Hour12 - 12 hour time string
	Hour12 string = "12h"
	// Daily time string
	Daily string = "d"
	// Day3 - 3 day time string
	Day3 string = "3d"
	// Weekly time string
	Weekly string = "w"
	// Monthly time string
	Monthly string = "m"
)

type Quote struct {
	TsCode    string      `json:"symbol"`
	Precision int64       `json:"-"`
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

func (q *Quote) Save()  {
	file, _ := json.MarshalIndent(q, "", " ")
	_ = ioutil.WriteFile(filepath.Join(DataDir, "Daily", q.TsCode + ".json"), file, 0644)
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

func (q *Quote) Update()  {
	params := make(map[string]string)
	params["ts_code"] = q.TsCode
	if len(q.Date) == 0 {
		params["start_date"] = "20150101"
	} else {
		lastDate := q.Date[len(q.Date) - 1]
		if time.Now().Format("20060102") ==  lastDate.Format("20060102") {
			return
		}
		params["start_date"] = lastDate.Format("20060102")
	}
	var fields = []string {}
	data, _ := c.Daily(params, fields)
	bars := len(data.Data.Items)
	_q := Quote{
		TsCode:     q.TsCode,
		Date:       make([]time.Time, bars),
		Open:       make([]float64, bars),
		High:       make([]float64, bars),
		Low:        make([]float64, bars),
		Close:      make([]float64, bars),
		Volume:     make([]float64, bars),
		PreClose:   make([]float64, bars),
		Change:     make([]float64, bars),
		PctChg:     make([]float64, bars),
		Amount:     make([]float64, bars),
	}

	l := len(data.Data.Items)
	if l == 0 {
		return
	}

	for _index, item := range data.Data.Items {
		index := l - _index - 1
		_q.Date[index],_ = time.Parse("20060102", item[1].(string))
		_q.Open[index] = item[2].(float64)
		_q.High[index] = item[3].(float64)
		_q.Low[index] = item[4].(float64)
		_q.Close[index] = item[5].(float64)
		_q.Volume[index] = item[9].(float64)
		_q.PreClose[index] = item[6].(float64)
		_q.Change[index] = item[7].(float64)
		_q.PctChg[index] = item[8].(float64)
		_q.Amount[index] = item[10].(float64)
  }

	q.Date = append(q.Date, _q.Date...)
	q.Open = append(q.Open, _q.Open...)
	q.High = append(q.High, _q.High...)
	q.Low = append(q.Low, _q.Low...)
	q.Close = append(q.Close, _q.Close...)
	q.Volume = append(q.Volume, _q.Volume...)
	q.PreClose = append(q.PreClose, _q.PreClose...)
	q.Change = append(q.Change, _q.Change...)
	q.PctChg = append(q.PctChg, _q.PctChg...)
	q.Amount = append(q.Amount, _q.Amount...)
	q.CalcMA()
	q.Save()
}

func NewQuote(ts_code string) Quote {
	q := Quote{
		TsCode:     ts_code,
		Date:       make([]time.Time, 0),
		Open:       make([]float64, 0),
		High:       make([]float64, 0),
		Low:        make([]float64, 0),
		Close:      make([]float64, 0),
		Volume:     make([]float64, 0),
		PreClose:   make([]float64, 0),
		Change:     make([]float64, 0),
		PctChg:     make([]float64, 0),
		Amount:     make([]float64, 0),
	}
	path := filepath.Join(DataDir, "Daily", ts_code + ".json")
	if fileExists(path) {
		file, _ := ioutil.ReadFile(path)
		_ = json.Unmarshal([]byte(file), &q)
	}
	q.Update()
	return q
}
