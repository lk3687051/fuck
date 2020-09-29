
// main.go
package main

import (
    "encoding/json"
    "log"
    "net/http"
    "github.com/gorilla/mux"
    "github.com/gorilla/handlers"
    "fuck/api"
)

func handleRequests() {
    myRouter := mux.NewRouter().StrictSlash(true)
    // myRouter.HandleFunc("/", homePage)
    myRouter.HandleFunc("/stocks", api.AllStocks)
    myRouter.HandleFunc("/stock/{code}", api.SingleArticle)

    myRouter.HandleFunc("/pool/{category}", api.AllPools)

    log.Fatal(http.ListenAndServe("127.0.0.1:8888",handlers.CORS()(myRouter)))
}

func InitData()  {
  stocks = fuck.GetStockList()
}
func main() {
  InitData()
  handleRequests()
}
