package fuck
import (
  "time"
)
type TradePeriod int
const (
    FreePeriod TradePeriod = 1 + iota
    CallPeriod
    ContinuousPeriod
)

func IsTradeDate(t time.Time) bool {
  return true
}

func GetTradePeriod(t time.Time) TradePeriod {
  return FreePeriod
}

func IsSameDay(t1 time.Time, t2 time.Time) bool {
  if t1.Year() == t2.Year() &&
    t1.Month() == t2.Month() && 
    t1.Day() == t2.Day()  {
      return true
    }

  return false
}
