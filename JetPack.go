package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/p4tin/jetpack/model/Order"
	"github.com/p4tin/jetpack/model/PriceInfo"
	"github.com/p4tin/jetpack/model/PricingAdjustment"
	"fmt"
)

func main() {
	mux := mux.NewRouter()
	mux.HandleFunc("/pricing", Pricing)
	log.Println("Starting to listen on port 4101")
	log.Fatal(http.ListenAndServe(":4101", mux))
}

func Pricing(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		if r.Body == nil {
			http.Error(w, "Please send a request body", 400)
			return
		}
		var order order.Order
		err := json.NewDecoder(r.Body).Decode(&order)
		if err != nil {
			log.Fatal(err)
		}
		order.PriceOrder()
		jsonOrder, _ := json.MarshalIndent(order, "    ", "  ")

		fmt.Fprintf(w, string(jsonOrder))
		//w.Write([]byte("Hello " + order.OrderPriceInfo.Amount))
	} else {
		if r.URL.Query().Get("forceGet") == "true" {
			adjustments := pricingadjustment.PricingAdjustment{AdjustmentDescription: "raw", Amount: 12.0}
			priceInfos := priceinfo.PriceInfo{Type: priceinfo.OrderPriceInfo, Amount: 12.0, Adjustments: adjustments, ProcessedTime: time.Now(), CurrencyCode: "USD"}
			w.Header().Set("Content-Type", "application/json; charset=UTF-8")
			json.NewEncoder(w).Encode(priceInfos)
		} else {
			w.Write([]byte("not supported"))
		}
	}

}
