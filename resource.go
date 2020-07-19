package fuck
import (
  // "path"
  "os"
  "path/filepath"
  "io/ioutil"
)

type ResourceName string
const (
    StockResource ResourceName = "stock"
    PoolResource ResourceName = "pool"
    StockQuoteDaily  ResourceName = "quotedaily"
    MarketResource    ResourceName = "market"
)

func fileExists(filename string) bool {
    info, err := os.Stat(filename)
    if os.IsNotExist(err) {
        return false
    }
    return !info.IsDir()
}

func SaveResource(rn ResourceName, index string, data []byte)  {
  pathname := ""
  if index == "" {
    pathname = filepath.Join(DataDir, string(rn) + ".json")
  } else {
    pathname = filepath.Join(DataDir, string(rn), index + ".json")
  }
	ioutil.WriteFile(pathname, data, 0644)
}

func LoadResource(rn ResourceName, index string) []byte{
  pathname := ""
  if index == "" {
    pathname = filepath.Join(DataDir, string(rn) + ".json")
  } else {
    pathname = filepath.Join(DataDir, string(rn), index + ".json")
  }
  file, _ := ioutil.ReadFile(pathname)
  return file
}
