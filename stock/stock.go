package stock
import (
	"fmt"
	"time"
	"fuck/config"
	"encoding/json"
	// log "github.com/sirupsen/logrus"
)

type DailyQuota struct {
	Open      float64   `json:"open"`
	High      float64   `json:"high"`
	Low       float64   `json:"low"`
	Close     float64   `json:"close"`
	Volume    float64   `json:"volume"`
	PreClose  float64   `json:"pre_close"`
	Change    float64     `json:"change"`
	PctChg    float64     `json:"pct_chg"`
	Amount    float64     `json:"amount"`
}

type Stock struct {
	TsCode string  `json:"ts_code"`
	Name      string  `json:"name"`
	Area      string  `json:"area"`
	Industry   string  `json:"industry"`
  ListDate    time.Time  `json:"list_date"`
  Exchange    string     `json:"exchange"`
	DailyQuote  map[string]DailyQuota
}

func (s *Stock) Update()  {
	data, _ := json.Marshal(s)
	config.Rdb.HSet("stocks", s.TsCode, data)
}

func NewStock() Stock {
	s := Stock{}
	return s
}

func GetAllStock() {
	stocks := config.Rdb.HGetAll("stocks")
	fmt.Printf("%+v", stocks)
}
