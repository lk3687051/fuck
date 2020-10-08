package main

import (
  "time"
  "context"
  "fmt"
  log "github.com/sirupsen/logrus"
)

/*下面是简单的CURD*/
//创建
func Create(index string, id string, data map[string]interface{}) {
    //使用结构体
    _, err := client.Index().
        Index(index).
        Type("doc").
        Id(id).
        BodyJson(data).
        Do(context.Background())
    if err != nil {
        log.Error(err)
    }
}

func GetStockDailyByTushare(ts_code string) {
  params := make(map[string]string)
  params["ts_code"] = ts_code
  params["start_date"] = "20150101"

  log.Debugf("Get stock %s quote from tushare %s", ts_code, params["start_date"])
  var fields = []string {}
  data, _ := C.Daily(params, fields)
  fmt.Println(data.Data.Fields)
  bulkRequest := client.Bulk()
  for _, item := range data.Data.Items {
    s := map[string]interface{}{}
    for i, field := range data.Data.Fields {
      s[field] =  item[i]
    }
    indexReq := elastic.NewBulkIndexRequest().Index("daily").Id(s["ts_code"].(string) + "." + s["trade_date"].(string)).Doc(s)
    bulkRequest.Add(indexReq)
  }
  bulkResponse, err := bulkRequest.Do(context.TODO())
	if err != nil {
		log.Error(err)
	} else {
    log.Infof("%+v", bulkResponse)
  }
}

func main()  {
  startT := time.Now()
  // Stocks := make([]string, 0)
	params := make(map[string]string)
  // 字段
  var fields = []string {"ts_code","symbol","name","area","industry","list_date", "exchange"}
  // 根据api 请求对应的接口
  data, _ := C.StockBasic(params, fields)
  // return data.Data.Items
  for index, item := range data.Data.Items {
    if index % 100 == 0 {
      log.Infof("Now we are process %d stock", index)
    }
    // fmt.Println(_s)
    s := map[string]interface{}{}
    for i, field := range data.Data.Fields {
      s[field] =  item[i]
    }
    Create("basic", s["ts_code"].(string), s)
    GetStockDailyByTushare(s["ts_code"].(string))
  }

  tc := time.Since(startT)	//计算耗时
  fmt.Printf("time cost = %v\n", tc)
  // return Stocks
}
