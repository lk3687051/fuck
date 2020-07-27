package fuck
import (
  log "github.com/sirupsen/logrus"
)
func (p *Pool)Statistics()  {
  Up := 0
  Sum := 0
  stocks := GetStockByCodes(p.Stocks)
  for _,s :=  range stocks {
    length := len(s.DailyQuote.UP)
    if length < 3 {
      continue
    }
    Sum = Sum + 3
    if s.DailyQuote.UP[length - 1] {
      Up = Up + 1
    }
    if s.DailyQuote.UP[length - 2] {
      Up = Up + 1
    }
    if s.DailyQuote.UP[length - 3] {
      Up = Up + 1
    }
  }
  if float64(Up*100)/float64(Sum) > 65.0 {
    log.Infof("%s:%s %d stock %d up   %2f\n", p.Category, p.Name, len(p.Stocks), Up, float64(Up*100)/float64(Sum))
  }

  // 计算涨跌停个数
  limitUp := 0
  limitDown := 0
  for _,s :=  range stocks {
    length := len(s.DailyQuote.Date)
    if length == 0 {
      continue
    }
    if s.DailyQuote.IslimitUp[length - 1] {
      limitUp = limitUp + 1
    }

    if s.DailyQuote.IslimitDown[length - 1] {
      limitDown = limitDown + 1
    }
  }
  log.Infof("%s:%s All:%d limitUp:%d limitDown:%d\n", p.Category, p.Name, len(p.Stocks), limitUp, limitDown)
}
