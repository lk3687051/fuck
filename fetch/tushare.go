package fetch
import (
  "time"
  // "fmt"
  "encoding/json"
  "fuck/stock"
  "fuck/config"
)

func DownLoadAllStock()  {
	params := make(map[string]string)
	// 字段
	var fields = []string {}
	// 根据api 请求对应的接口
	data, _ := config.C.StockBasic(params, fields)
	// return data.Data.Items
	for _, _s := range data.Data.Items {
    ld,_ := time.Parse("20060102", _s[6].(string))
    s := stock.Stock{
  		TsCode: _s[0].(string),
  		Name: _s[2].(string),
  		Area: _s[3].(string),
  		Industry: _s[4].(string),
      Exchange: _s[5].(string),
  		ListDate: ld,
  	}
    data, _ := json.Marshal(s)
    config.Rdb.HSet("stocks", s.TsCode, data)
    GetDaily(s.TsCode)
	}
}

func GetDaily(ts_code string)  {
	params := make(map[string]string)
  params["ts_code"] = ts_code
  params["start_date"] = time.Date(2015, 7, 1, 0, 0, 0, 0, time.UTC).Format("20060102")
  var fields = []string {}
  data, _ := config.C.Daily(params, fields)
	key := "stock:" + ts_code + ":daily"
  for _, item := range data.Data.Items {
    q := stock.DailyQuota{}
    q.Open = item[2].(float64)
    q.High = item[3].(float64)
    q.Low = item[4].(float64)
    q.Close = item[5].(float64)
    q.Volume = item[9].(float64)
    q.PreClose = item[6].(float64)
    q.Change = item[7].(float64)
    q.PctChg = item[8].(float64)
    q.Amount = item[10].(float64)
    data, _ := json.Marshal(q)
    config.Rdb.HSet(key, item[1].(string), data)
  }
}
