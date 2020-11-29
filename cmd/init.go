package main
import (
	tusharespider "fuck/spider/tushare"
	"fuck/store/elas"
	"fuck/config"
  "github.com/sirupsen/logrus"
)

var logger = logrus.WithFields(logrus.Fields{
	"component": "init",
})

func main()  {
  logger.Info("Start init all the data")
	spider := tusharespider.NewSpider(config.Config.TushareToken)
	store  := elas.NewStore(config.Config.ElasticHost)
	stocks,err := spider.StockBasic()
	if err != nil {
    logger.WithError(err).Error("Can not get stocks")
    return
  }
	store.SaveStocks(stocks)
	// 建立Cache
	// stockMap = map[string]&types.Stock
	// for _, stock := range stocks {
	// 	stockMap[stock.TsCode] = stock
	// }

  logger.Info("Start get daily data")
	for _, stock := range stocks {
		dailys,_ := spider.Daily(stock.TsCode, config.Config.DailyLength)
		for i, _ := range dailys {
			dailys[i].Area = stock.Area
			dailys[i].Industry = stock.Industry
		}
		store.SaveDailys(dailys)
	}
  logger.Info("End get daily data")
}
