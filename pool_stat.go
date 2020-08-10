package fuck
import (
  // log "github.com/sirupsen/logrus"
)

type Stat struct {
  Nums     int `json:"nums"`
	LimitUp    int  `json:"limit_up"`
  LimitDown    int  `json:"limit_down"`
  LimitUpHistory    int  `json:"limit_up_history"`
  LimitDownHistory    int  `json:"limit_down_history"`
  UpNum    int  `json:"up_nums"`
  DownNum    int  `json:"down_nums"`
}

func (p *Pool)Statistics()  {
  p.Stat.LimitUp = 0
  p.Stat.LimitDown = 0
  p.Stat.UpNum = 0
  p.Stat.DownNum = 0

  for _, tscode := range p.Stocks {
    s, _ := GlobalManager.Stocks[tscode]
    length := len(s.DailyQuote.UP)
    if length < 1 {
      continue
    }
    if s.DailyQuote.UP[length - 1] {
      p.Stat.UpNum = p.Stat.UpNum + 1
    }

    if s.DailyQuote.IslimitUp[length - 1] {
      p.Stat.LimitUp = p.Stat.LimitUp + 1
    }

    if s.DailyQuote.IslimitDown[length - 1] {
      p.Stat.LimitDown = p.Stat.LimitDown + 1
    }
  }
}
