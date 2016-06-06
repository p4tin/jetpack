package priceinfo

import (
	"github.com/p4tin/jetpack/model/PricingAdjustment"
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
	Type          PriceInfoType                       `json:"type,omitempty"`
	CurrencyCode  string                              `json:"currencyCode,omitempty"`
	ProcessedTime time.Time                           `json:"processedTime,omitempty"`
	Amount        float64                             `json:"amount,omitempty"`
	UnitPrice     float64                             `json:"unitPrice,omitempty"`
	Discounted    bool                                `json:"discouted,omitempty"`
	Adjustments   pricingadjustment.PricingAdjustment `json:"adjustments,omitempty"`
}

func NewPriceInfo(pPrice float64) *PriceInfo {
	info := new(PriceInfo)
	info.UnitPrice = pPrice
	return info
}
