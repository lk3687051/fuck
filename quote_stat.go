package fuck
import(
	// "fmt"
  "math"
  "github.com/markcheno/go-talib"
)

func (q *Quote) PreCalc()  {
  // 计算涨停跌停
  q.IslimitUp = make([]bool, len(q.Date))
  q.IslimitDown = make([]bool, len(q.Date))
  for index, _ := range q.Date {
    limitUp := math.Floor(q.PreClose[index]*1.1*100 + 0.5)/100
    limitDown := math.Floor(q.PreClose[index]*0.9*100 + 0.5)/100
    q.IslimitUp[index] = limitUp == q.Close[index]
    q.IslimitDown[index] = limitDown == q.Close[index]
  }

  // 设置是否上涨
  q.UP = make([]bool, len(q.Date))
  for index, _ := range q.Date {
    if q.Close[index] > q.PreClose[index] {
      q.UP[index] = true
    } else {
      q.UP[index] = false
    }
  }

  // 计算MA值
	if len(q.Date) <= 63 {
		return
	}
	q.MA5 = talib.Ma(q.Close, 5, talib.SMA)
	q.MA10 = talib.Ma(q.Close, 10, talib.SMA)
	q.MA20 = talib.Ma(q.Close, 20, talib.SMA)
	q.MA62 = talib.Ma(q.Close, 62, talib.SMA)
}
