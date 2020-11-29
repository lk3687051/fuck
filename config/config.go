package config
type Configuration struct {
  ElasticHost        string    `envconfig:"ELASTIC_HOST" default:"http://127.0.0.1:9200/"`
  TushareToken       string    `envconfig:"TUSHARE_TOKEN" default:"http://127.0.0.1:9200/"`
  DailyLength        int       `envconfig:"DAILAY_LENGTH" default:"365"`
}

var Config Configuration
func init() {
  Config = Configuration {
      ElasticHost: "http://127.0.0.1:9200",
      TushareToken: "33454fe9ccab969249773ab8a05a5a7dad45c23a15cc71e17e6a33b6",
      DailyLength: 365,
    }
}
