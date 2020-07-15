package market
import()
var DataDir = "D:\\stock"
type Market struct {
	Quotes    []Quote
	TsCodes   []string
}

func (m *Market) GetBasic()  {
	params := make(map[string]string)
  // 字段
  var fields = []string {"ts_code","symbol","name","area","industry","list_date"}
  // 根据api 请求对应的接口
  data, _ := c.StockBasic(params, fields)

	l := len(data.Data.Items)
	m.TsCodes = make([]string, l)
	for i, item := range data.Data.Items {
		m.TsCodes[i] = item[0].(string)
	}
}
func (m *Market) Init()  {
	for _, ts_code := range m.TsCodes {
		q := NewQuote(ts_code)
		m.Quotes = append(m.Quotes, q)
	}
}

func NewMarket()  {
	m := Market{}
	m.GetBasic()
	m.Init()
}
