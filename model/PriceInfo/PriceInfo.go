package priceinfo

import (
	"JetPack/model/PricingAdjustment"
	"time"
)

type PriceInfoType string

const (
	OrderPriceInfo    PriceInfoType = "OrderPriceInfo"
	ItemPriceInfo     PriceInfoType = "ItemPriceInfo"
	TaxPriceInfo      PriceInfoType = "TaxPriceInfo"
	ShippingPriceInfo PriceInfoType = "ShippingPriceInfo"
)

type PriceInfo struct {
	Type          PriceInfoType                       `json:"type"`
	CurrencyCode  string                              `json:"currencyCode"`
	ProcessedTime time.Time                           `json:"processedTime"`
	Amount        float32                             `json:"amount"`
	UnitPrice     float32                             `json:"unitPrice"`
	Discounted    bool                                `json:"discouted"`
	Adjustments   pricingadjustment.PricingAdjustment `json:"adjustments"`
}

func NewPriceInfo(pPrice float32) *PriceInfo {
	info := new(PriceInfo)
	info.UnitPrice = pPrice
	return info
}
