package main
import (
  // "fmt"
  "time"
  "github.com/jinzhu/gorm"
  _ "github.com/go-sql-driver/mysql"
  "github.com/markcheno/go-quote"
	// "github.com/markcheno/go-talib"
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

func main()  {
  db, err := gorm.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/test")
  if err != nil {
    panic("连接数据库失败")
  }
  defer db.Close()

  var dailys []StockDaily
  db.Where("ts_code = ?", "000001.SZ").Find(&dailys)
  length := len(dailys)
  bars := quote.Quote{
		Symbol: "000001.SZ",
		Date:   make([]time.Time, length),
		Open:   make([]float64, length),
		High:   make([]float64, length),
		Low:    make([]float64, length),
		Close:  make([]float64, length),
		Volume: make([]float64, length),
	}

  for i, daily := range dailys {
    date, _ := time.Parse("20060102", daily.TradeDate)
    bars.Date[i] = date
    bars.Open[i] = daily.Open
    bars.High[i] = daily.High
    bars.Low[i] = daily.Low
    bars.Close[i] = daily.Close
    bars.Volume[i] = daily.Vol
  }
  bars.WriteHighstock("")
}
