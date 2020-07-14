package collector
import (
  "github.com/jinzhu/gorm"
  _ "github.com/go-sql-driver/mysql"
)

bars := quote.Quote{
  Symbol: "000001.SZ",
  Date:   make([]time.Time, length),
  Open:   make([]float64, length),
  High:   make([]float64, length),
  Low:    make([]float64, length),
  Close:  make([]float64, length),
  Volume: make([]float64, length),
}

func UpdateStockBasic()  {
  db, err := gorm.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/test")
  if err != nil {
    panic("连接数据库失败")
  }
  defer db.Close()

  db.AutoMigrate(&StockBasic{})

  // 参数
  params := make(map[string]string)
  // 字段
  var fields = []string {"ts_code","symbol","name","area","industry","list_date"}
  // 根据api 请求对应的接口
  data, _ := c.StockBasic(params, fields)
  for _, item := range data.Data.Items {
    if item[3] == nil {
      continue
    }
    UpdateStockDaily(item[0].(string))
    db.Create(&StockBasic{TsCode: item[0].(string), Symbol: item[1].(string), Name: item[2].(string), Area: item[3].(string), Industry: item[4].(string), ListDate: item[5].(string)})
  }
}
