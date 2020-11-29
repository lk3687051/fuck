package elas
import (
	// "github.com/sirupsen/logrus"
	"github.com/olivere/elastic/v7"
)
// Store is an implementation of the sensu-go/backend/store.Store iface.
type Store struct {
	client         *elastic.Client
}

// NewStore creates a new Store.
func NewStore(host string) *Store {
	client, err := elastic.NewClient(
				elastic.SetURL(host),
				// elastic.SetTraceLog(logger),
	)

  if err != nil {
      logger.Fatal(err)
  }
	store := &Store{
		client:         client,
	}
	return store
}
