package main

import (
	"fmt"
	"net/http"
	"fuck"
	"html/template"
	log "github.com/sirupsen/logrus"
)
type StockListPage struct {
    Stocks     []*fuck.Stock
}

type PoolListPage struct {
    Pools     []*fuck.Pool
}

func stock_list(w http.ResponseWriter, req *http.Request) {
	page := StockListPage{}
	tmpl := template.Must(template.ParseFiles("templates/stock_table.tmpl"))
	page.Stocks = fuck.GetStockList()
	tmpl.Execute(w, page)
}

func main() {
	fs := http.FileServer(http.Dir("./static"))
  http.Handle("/", fs)
	http.HandleFunc("/stock_tables", stock_list)
  log.Println("Listening...")
	err := http.ListenAndServe("127.0.0.1:8083", nil)
	if err != nil {
				fmt.Println(err)
		}
}
