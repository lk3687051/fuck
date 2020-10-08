package config
import (
  "github.com/ShawnRong/tushare-go"
  "github.com/go-redis/redis"
  "github.com/olivere/elastic/v7"
)
var TsClient = tushare.New("6d8cb43818f3724d89d561797bd6b37a9a3555c109160aa51ed7428e")
var EsClient *elastic.Client
func init()  {
  var err error
  EsClient, err = elastic.NewClient(elastic.SetURL("http://127.0.0.1:9200/"))
  if err != nil {
      log.Fatal(err)
  }
}
