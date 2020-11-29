package types
import (
  "time"
)
type Daily struct {
  TsCode string `json:"ts_code"`
  TradeDate time.Time `json:"trade_date"`
  Close float64  `json:"close"`
  Open float64 `json:"open"`
  High float64 `json:"high"`
  Low float64 `json:"low"`
  PreClose float64 `json:"pre_close"`
  Change float64 `json:"change"`
  PctChg float64 `json:"pct_chg"`
  Vol float64 `json:"vol"`
  Amount float64 `json:"amount"`
  Area string  `json:"area"`
  Industry string `json:"industry"`
}

type Stock struct {
  TsCode string `json:"ts_code"`
  Symbol string `json:"symbol"`
  Name string `json:"name"`
  Area string  `json:"area"`
  Industry string `json:"industry"`
  ListDate time.Time `json:"list_date"`
  Concepts []string `json:"concepts"`
}
