package spider
type Spider interface {
  TushareSpider
}

type TushareSpider interface {
	// DeleteAssetByName deletes an asset using the given name and the
	// namespace stored in ctx.
	GetDailyByTscode(ctx context.Context, name string) [] *types.Daily, error
  GetStockBasic(ctx context.Context, daily *types.Daily, code string, date time.Time) error
}
