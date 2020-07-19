package fuck
import (
  "time"
  // "encoding/json"
  // log "github.com/sirupsen/logrus"
)

type Market struct {
    UpdateAt        time.Time       `json:"updateAt"`
  	// Stocks          []string        `json:"stocks"`
    // Pools           []string        `json:"pools"`
    AreaPools       map[string] string `json:"areaPools"`
    IndustryPools   map[string] string `json:"industryPools"`
    ExchangePools   map[string] string  `json:"exchangePools"`
    ConceptPools    map[string] string  `json:"conceptPools"`
}


func Start()  {
  SetupPools()
  SetupStocks()
  for _,p := range PoolMap {
    p.Statistics()
  }
}
