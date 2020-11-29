package main
import (
	"fmt"
	"time"
	"fuck/store/elas"
	"fuck/config"
)

func main()  {
	store  := elas.NewStore(config.Config.ElasticHost)
	// stocks,_ := store.GetStocks()
	// fmt.Println(stocks)
	// dailys,_ := store.GetDailyByTsCode("000001.SZ")
	// fmt.Println(dailys)

	trade_date,_ := time.Parse("20060102", "20201117")
	dailys,_ := store.GetDailyByTradeDate(trade_date)
	fmt.Println(dailys)
}
