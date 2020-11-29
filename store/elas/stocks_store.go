package elas
import (
	"context"
	"fuck/types"
  "encoding/json"
	"github.com/olivere/elastic/v7"
)
const (
	stockIndexName = "stock"
)

func (s *Store) SaveStocks(stocks []types.Stock) (error) {
	bulkRequest := s.client.Bulk().Index(stockIndexName)
	for _, stock := range stocks{
		r := elastic.NewBulkIndexRequest().Index(stockIndexName).Id(stock.TsCode).Doc(stock)
		bulkRequest.Add(r)
	}
	bulkRequest.Do(context.TODO())
	return nil
}

func (s *Store) GetStocks() ([]types.Stock, error) {
  searchResult, err := s.client.Search().
      Index(stockIndexName).
      Size(1000). // 默认拿10000个结果肯定够的
      Do(context.TODO())
  if err != nil {
    logger.WithError(err).Error("Can not get all stocks")
    return nil, err
  }

  stocks := make([]types.Stock, len(searchResult.Hits.Hits))
  for i, hit := range searchResult.Hits.Hits {
		err := json.Unmarshal(hit.Source, &stocks[i])
		if err != nil {
			logger.WithError(err).Error("Decode %s", hit.Source)
      return nil, err
		}
	}
  return stocks, nil
}
