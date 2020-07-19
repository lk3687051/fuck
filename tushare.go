package fuck
import (
  "time"
  log "github.com/sirupsen/logrus"
)

func GetStocks() []Stock {
  Stocks := make([]Stock, 0)
	params := make(map[string]string)
  // 字段
  var fields = []string {"ts_code","symbol","name","area","industry","list_date", "exchange"}
  // 根据api 请求对应的接口
  data, _ := c.StockBasic(params, fields)
  // return data.Data.Items
  for _, _s := range data.Data.Items {
    if _s[3] == nil {
      continue
    }
    ld,_ := time.Parse("20060102", _s[6].(string))
    log.Debugf("Get data %+v\n", _s)
    s := Stock{
  		TsCode: _s[0].(string),
  		Name: _s[2].(string),
  		Area: _s[3].(string),
  		Industry: _s[4].(string),
      Exchange: _s[5].(string),
  		ListDate: ld,
  	}
    Stocks = append(Stocks, s)
  }
  return Stocks
}

func GetDaily(ts_code string, date time.Time) Quote {
  params := make(map[string]string)
  params["ts_code"] = ts_code
  if date.After(dailyStart) {
    params["start_date"] = date.Format("20060102")
  } else {
    params["start_date"] = dailyStart.Format("20060102")
  }
  log.Debugf("Get stock %s uote from tushare %s", ts_code, params["start_date"])
  var fields = []string {}
  data, _ := c.Daily(params, fields)
  bars := len(data.Data.Items)
  q := Quote{
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

  for _index, item := range data.Data.Items {
    index := bars - _index - 1
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
