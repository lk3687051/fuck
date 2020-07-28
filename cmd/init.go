package main
import (
  "fuck"
  log "github.com/sirupsen/logrus"
)

func main()  {
  log.SetFormatter(&log.TextFormatter{
		DisableColors: true,
		FullTimestamp: true,
	})
  log.SetLevel(log.DebugLevel)
  log.Info("Now Start fuck program\n")
  // fuck.SetupStocks()
  fuck.SetupData()
}
