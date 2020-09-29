package api
import (
  "fuck/stock"
)

func returnAllStocks(w http.ResponseWriter, r *http.Request) {
  stocks := stock.GetAllStock()
  json.NewEncoder(w).Encode(stocks)
}

func returnAllPools(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    key := vars["category"]
    pools := fuck.GetPoolByCategory(key)
    json.NewEncoder(w).Encode(pools)
}

func returnSingleArticle(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    key := vars["code"]

    for _, stock := range stocks {
        if stock.TsCode == key {
            json.NewEncoder(w).Encode(stock)
            return
        }
    }
}
