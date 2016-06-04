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
)

func main() {
	mux := mux.NewRouter()
	mux.HandleFunc("/pricing", Pricing)
	log.Fatal(http.ListenAndServe(":4101", mux))
}

func Pricing(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		decoder := json.NewDecoder(r.Body)
		var order order.Order
		err := decoder.Decode(&order)
		if err != nil {
			panic("failed to decode")
		}
		order.PriceOrder()

		w.Write([]byte("Hello " + order.ShippingAddress.Address1))
	} else {
		if r.URL.Query().
			Get("forceGet") == "true" {
			adjustments := pricingadjustment.PricingAdjustment{AdjustmentDescription: "raw", Amount: 12.0}
			priceInfos := priceinfo.PriceInfo{Type: priceinfo.OrderPriceInfo, Amount: 12.0, Adjustments: adjustments, ProcessedTime: time.Now(), CurrencyCode: "USD"}
			w.Header().Set("Content-Type", "application/json; charset=UTF-8")
			json.NewEncoder(w).Encode(priceInfos)
		} else {
			w.Write([]byte("not supported"))
		}
	}

}
