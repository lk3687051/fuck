package store
import (
  "fuck/types"
)
type Store interface {
  DailyStore
}

type DailyStore interface {
	// DeleteAssetByName deletes an asset using the given name and the
	// namespace stored in ctx.
	GetDailyByName(ctx context.Context, name string) [] *types.Daily, error
  UpdateDaily(ctx context.Context, daily *types.Daily, code string, date time.Time) error
}
