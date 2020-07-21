package fuck
import(
	// "fmt"
  	"github.com/markcheno/go-talib"
)

func (q *Quote) PreCalc()  {
	if len(q.Date) <= 63 {
		return
	}
	q.MA5 = talib.Ma(q.Close, 5, talib.SMA)
	q.MA10 = talib.Ma(q.Close, 10, talib.SMA)
	q.MA20 = talib.Ma(q.Close, 20, talib.SMA)
	q.MA62 = talib.Ma(q.Close, 62, talib.SMA)

  // 设置是否上涨
  q.UP = make([]bool, len(q.Date))
  for index, _ := range q.Date {
    if q.Close[index] > q.PreClose[index] {
      q.UP[index] = true
    } else {
      q.UP[index] = false
    }
  }
}
