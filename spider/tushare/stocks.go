package tushare
import (
  // "fmt"
  "time"
  "errors"
  "fuck/types"
)

func (spider *Spider) StockBasic() ([]types.Stock, error) {
  // stocks := []types.Stock{}
  var fields = []string{"ts_code","symbol","name","area","industry","list_date"}
  params := make(map[string]string)
  // params["is_hs"] = "N"
  params["list_status"] = "L"
  params["exchange"] = ""
	body := RequestData{
		API_NAME: "stock_basic",
		Token: spider.token,
		PARAMS: params,
		Fields: fields,
	}
	resp, err := spider.client.R().
				SetBody(body).
				SetResult(ResponseData{}).
				Post("/")

  if err != nil {
    logger.WithError(err).Error("Can not get response")
    return nil, err
  }
  r := resp.Result().(*ResponseData)
  if r.Code != 0 {
    logger.Errorf("Response error %f", r.Msg.(string))
    return nil, errors.New(r.Msg.(string))
  }

  stocks := make([]types.Stock, len(r.Data.Items))
  for i, stock := range r.Data.Items {
    stocks[i].TsCode = stock[0].(string)
    stocks[i].Symbol = stock[1].(string)
    stocks[i].Name = stock[2].(string)
    stocks[i].Area = stock[3].(string)
    stocks[i].Industry = stock[4].(string)
    stocks[i].ListDate,_ = time.Parse("20060102", stock[5].(string))
  }
  // fmt.Println(stocks)
	return stocks, err
}
