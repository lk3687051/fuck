package tushare
import (
	"github.com/go-resty/resty/v2"
)

// Endpoint URL
const Endpoint = "http://api.tushare.pro"
type Spider struct {
	token  string
	client *resty.Client
}

type RequestData struct {
  API_NAME string  `json:"api_name"`
  Token string  `json:"token"`
  PARAMS map[string]string  `json:"params"`
  Fields []string  `json:"fields"`
}

type ResponseData struct {
	RequestID string      `json:"request_id"`
	Code      int         `json:"code"`
	Msg       interface{} `json:"msg"`
	Data      struct {
		Fields []string        `json:"fields"`
		Items  [][]interface{} `json:"items"`
	} `json:"data"`
}

func NewSpider(token string) *Spider {
  client := resty.New()
  client.SetHostURL(Endpoint)
  client.SetHeader("Accept", "application/json")
	// client.SetDebug(true)
  spider := &Spider{
    token: token,
    client: client,
  }
	return spider
}
