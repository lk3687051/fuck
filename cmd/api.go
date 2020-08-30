
// main.go
package main

import (
    "encoding/json"
    "log"
    "net/http"
    "github.com/gorilla/mux"
    "github.com/gorilla/handlers"
    "fuck"
)

var stocks []*fuck.Stock
func returnAllStocks(w http.ResponseWriter, r *http.Request) {
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
        }
    }
}

func handleRequests() {
    myRouter := mux.NewRouter().StrictSlash(true)
    // myRouter.HandleFunc("/", homePage)
    myRouter.HandleFunc("/stocks", returnAllStocks)
    myRouter.HandleFunc("/stock/{code}", returnSingleArticle)

    myRouter.HandleFunc("/pool/{category}", returnAllPools)

    log.Fatal(http.ListenAndServe("127.0.0.1:8888",handlers.CORS()(myRouter)))
}

func InitData()  {
  stocks = fuck.GetStockList()
}
func main() {
  InitData()
  handleRequests()
}
