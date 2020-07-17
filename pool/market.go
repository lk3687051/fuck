package market
import(
	"fmt"
)
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

func (m *Market) CalcMA()  {
	up5 := 0
	down5 := 0
	up10 := 0
	down10 := 0
	up20 := 0
	down20 := 0
	up62 := 0
	down62 := 0
	for _, q := range m.Quotes{
		l := len(q.Date)
		if l == 0 || q.MA5 == nil || q.MA5[l-1] == 0{
			continue
		}

		fmt.Printf("Update %s\n", q.TsCode)
		if q.Close[l -1] > q.MA5[l -1] {
			up5++
		} else{
			down5++
		}
		if q.Close[l -1] > q.MA10[l -1] {
			up10++
		} else{
			down10++
		}
		if q.Close[l -1] > q.MA20[l -1] {
			up20++
		} else{
			down20++
		}
		if q.Close[l -1] > q.MA62[l -1] {
			up62++
		} else{
			down62++
		}
	}
	fmt.Printf("%d up than ma5, %d down than MA5\n", up5, down5)
	fmt.Printf("%d up than ma10, %d down than ma10\n", up10, down10)
	fmt.Printf("%d up than ma20, %d down than ma20\n", up20, down20)
	fmt.Printf("%d up than ma62, %d down than ma62\n", up62, down62)
}

func NewMarket()  Market{
	m := Market{}
	m.GetBasic()
	m.Init()
	return m
}
