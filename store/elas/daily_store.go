package elas
import (
	"time"
	"context"
	"fuck/types"
	"encoding/json"
	"github.com/olivere/elastic/v7"
)
const (
	dailyIndexName = "daily"
)

func (s *Store) SaveDailys(dailys []types.Daily) (error) {
	bulkRequest := s.client.Bulk().Index(dailyIndexName)
	for _, daily := range dailys{
		r := elastic.NewBulkIndexRequest().Index(dailyIndexName).Id(daily.TsCode + "-" + daily.TradeDate.Format("20060102")).Doc(daily)
		bulkRequest.Add(r)
	}
	bulkRequest.Do(context.TODO())
	return nil
}

func (s *Store) GetDailyByTsCode(ts_code string) ([]types.Daily, error) {
	termQuery := elastic.NewTermQuery("ts_code.keyword", ts_code)
  searchResult, err := s.client.Search().
      Index(dailyIndexName).
			Query(termQuery).
			Sort("trade_date", false).
      Size(1000).
      Do(context.TODO())
  if err != nil {
    logger.WithError(err).Error("Can not get all stocks")
    return nil, err
  }

  dailys := make([]types.Daily, len(searchResult.Hits.Hits))
  for i, hit := range searchResult.Hits.Hits {
		err := json.Unmarshal(hit.Source, &dailys[i])
		if err != nil {
			logger.WithError(err).Error("Decode %s", hit.Source)
      return nil, err
		}
	}
  return dailys, nil
}

func (s *Store) GetDailyByTradeDate(trade_date time.Time) ([]types.Daily, error) {
	termQuery := elastic.NewTermQuery("trade_date", trade_date)
  searchResult, err := s.client.Search().
      Index(dailyIndexName).
			Query(termQuery).
			Sort("trade_date", false).
      Size(1000).
      Do(context.TODO())
  if err != nil {
    logger.WithError(err).Error("Can not get all stocks")
    return nil, err
  }

  dailys := make([]types.Daily, len(searchResult.Hits.Hits))
  for i, hit := range searchResult.Hits.Hits {
		err := json.Unmarshal(hit.Source, &dailys[i])
		if err != nil {
			logger.WithError(err).Error("Decode %s", hit.Source)
      return nil, err
		}
	}
  return dailys, nil
}
