package pricingadjustment

type PricingAdjustment struct {
	AdjustmentDescription string  `json:"adjustmentDescription"`
	QuantityAdjusted      int     `json:"quantityAdjusted"`
	PricingModel          string  `json:"pricingModel"`
	TotalAdjustment       float64 `json:"totalAdjustment"`
	Amount                float64 `json:"amount"`
}

func NewPricingAdjustment() *PricingAdjustment {
	adj := new(PricingAdjustment)
	return adj
}

func (adj *PricingAdjustment) AddDescription(pText string) {
	adj.AdjustmentDescription = pText
}
