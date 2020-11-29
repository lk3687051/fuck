package tushare
import (
  "time"
  "errors"
  "fuck/types"
  "go.uber.org/ratelimit"
)
const (
  dailyTushareLimit = 8
)
var dailyLimit = ratelimit.New(dailyTushareLimit)
func (spider *Spider) Daily(ts_code string, lenth int) ([]types.Daily, error) {
  dailyLimit.Take()
  end_date := time.Now()
  start_date := end_date.AddDate(0, 0, -lenth)
  var fields = []string{"ts_code", "trade_date", "open",
                        "high", "low", "close", "pre_close",
                        "change", "pct_chg", "vol", "amount"}
  params := make(map[string]string)
  params["ts_code"] = ts_code
  // params["adj"] = "hfq"
  params["start_date"] = start_date.Format("20060102")
  params["end_date"] = end_date.Format("20060102")
	body := RequestData{
		API_NAME: "daily",
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

  dailys := make([]types.Daily, len(r.Data.Items))
  for i, stock := range r.Data.Items {
    dailys[i].TsCode = stock[0].(string)
    dailys[i].TradeDate,_ = time.Parse("20060102", stock[1].(string))
    dailys[i].Open = stock[2].(float64)
    dailys[i].High = stock[3].(float64)
    dailys[i].Low = stock[4].(float64)
    dailys[i].Close = stock[5].(float64)
    dailys[i].PreClose = stock[6].(float64)
    dailys[i].Change = stock[7].(float64)
    dailys[i].PctChg = stock[8].(float64)
    dailys[i].Vol = stock[9].(float64)
    dailys[i].Amount = stock[10].(float64)
  }
	return dailys, err
}
