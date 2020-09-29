package main
import (
  "fuck/fetch"
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
  fetch.DownLoadAllStock()
}
