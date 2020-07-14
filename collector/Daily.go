package collector
import (
  "fmt"
  "github.com/jinzhu/gorm"
  _ "github.com/go-sql-driver/mysql"
)

type StockDaily struct {
  TsCode string `gorm:"index:idx_name_code"`
  TradeDate string `gorm:"index:idx_name_code"`
  Open   float64
  High    float64
  Low  float64
  Close  float64
  PreClose  float64
  Change  float64
  PctChg  float64
  Vol  float64
  Amount  float64
}

func UpdateStockDaily(ts_code string)  {
  db, err := gorm.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/test")
  if err != nil {
    panic("连接数据库失败")
  }
  defer db.Close()
  db.AutoMigrate(&StockDaily{})

  params := make(map[string]string)
  params["ts_code"] = ts_code
  params["start_date"] = "20150101"
  // params["end_date"] = "20180718"
  var fields = []string {}
  // 根据api 请求对应的接口
  data, _ := c.Daily(params, fields)

  for _, item := range data.Data.Items {
    if item[3] == nil {
      continue
    }
    fmt.Printf("%v\n", item)
    db.Create(&StockDaily{TsCode: item[0].(string), TradeDate: item[1].(string), Open: item[2].(float64),
        High: item[3].(float64), Low: item[4].(float64), Close: item[5].(float64),
        PreClose: item[3].(float64), Change: item[4].(float64), PctChg: item[5].(float64),
        Vol: item[4].(float64), Amount: item[5].(float64)})
  }
}
