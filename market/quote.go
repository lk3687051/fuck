package market
import (
	// "fmt"
	"time"
	// "path"
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
}

func (q *Quote)Update()  {

}

func (q *Quote)Load()  {
	// pathname := path.Join(DataDir, "quote", q.Symbol+".csv")
	// fileExists()
}

func (q *Quote)Save()  {
	// pathname := path.Join(DataDir, "quote", q.Symbol+".csv")
	// fileExists()
}

func NewQuote(ts_code string) Quote {
	params := make(map[string]string)
  params["ts_code"] = ts_code
  params["start_date"] = "20200710"
  // params["end_date"] = "20180718"
  var fields = []string {}
  // 根据api 请求对应的接口
  data, _ := c.Daily(params, fields)
	bars := len(data.Data.Items)
	// fmt.Printf("%+v", data)
  q := Quote{
    TsCode:     ts_code,
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
	for index, item := range data.Data.Items {
		q.Date[index],_ = time.Parse("20060102", item[1].(string))
		q.Open[index] = item[2].(float64)
		q.High[index] = item[3].(float64)
		q.Low[index] = item[4].(float64)
		q.Close[index] = item[5].(float64)
		q.Volume[index] = item[9].(float64)
		q.PreClose[index] = item[6].(float64)
		q.Change[index] = item[7].(float64)
		q.PctChg[index] = item[8].(float64)
		q.Amount[index] = item[10].(float64)
  }
	return q
}
