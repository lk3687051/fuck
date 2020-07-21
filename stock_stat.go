package fuck
import (
  log "github.com/sirupsen/logrus"
)
func (s *Stock) Statistics()  {
  log.Debugf("Statistics Stock  %s", s.Name)
  s.DailyQuote.PreCalc()
}
