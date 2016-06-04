package order

import (
	"math"
	"sync"

	"JetPack/model/Address"
	"JetPack/model/Item"
	"JetPack/model/PriceInfo"
	"JetPack/model/PricingAdjustment"
	"JetPack/model/Promotion"
)

type Order struct {
	Items              []item.Item                           `json:"items"`
	OrderPriceInfo     priceinfo.PriceInfo                   `json: "orderpriceInfo"`
	ShippingAddress    address.Address                       `json:"shippingAddress"`
	PricingAdjustments []pricingadjustment.PricingAdjustment `json:"PricingAdjsutments"`
}

func NewOrder() *Order {
	o := new(Order)
	o.OrderPriceInfo.Type = priceinfo.OrderPriceInfo
	return o
}

func (o *Order) AddItem(itm item.Item) {
	o.Items = append(o.Items, itm)
}

//
// This Function Calculates the Order Prices and returns each level's price in an PriceInfo struct
//
func (o *Order) PriceOrder() {
	var wg sync.WaitGroup
	//loop through all items
	wg.Add(len(o.Items))
	for i := 0; i < len(o.Items); i++ {
		go func(i int) {
			defer wg.Done()
			o.Items[i].PriceItem()
		}(i)
	}
	wg.Wait()

	sum := float32(0.0)
	for _, itm := range o.Items {
		sum += itm.ItemPriceInfo.Amount
	}
	o.OrderPriceInfo.Amount = sum

	//Check Order level promos and adjust
	o.ApplyOrderLevelPromotions()
}

//
// This Function Calculates the Order Level Prmotions, applies each that apply to the order total in the
// OrderPriceInfo and Also Add a PricingAdjustment to the Order so we can track why the price was changed
// and by how much
//
func (o *Order) ApplyOrderLevelPromotions() {
	for i := 0; i < len(promotion.Promotions); i++ {
		promo := promotion.Promotions[0]
		if promo.Type == promotion.OrderPromotion {
			if promo.PercentOff > 0.0 {
				adjustBy := float32(roundFloat(float64(o.OrderPriceInfo.Amount/100*promo.PercentOff)+.005, 2))
				o.OrderPriceInfo.Amount = o.OrderPriceInfo.Amount - adjustBy
				adj := pricingadjustment.NewPricingAdjustment()
				adj.AddDescription(promo.Name)
				adj.TotalAdjustment = -adjustBy
				o.PricingAdjustments = append(o.PricingAdjustments, *adj)
			}
		}
	}
}

// return rounded version of x with prec precision.
func roundFloat(x float64, prec int) float64 {
	var rounder float64
	pow := math.Pow(10, float64(prec))
	intermed := x * pow
	_, frac := math.Modf(intermed)
	intermed += .5
	x = .5
	if frac < 0.0 {
		x = -.5
		intermed -= 1
	}
	if frac >= x {
		rounder = math.Ceil(intermed)
	} else {
		rounder = math.Floor(intermed)
	}

	return rounder / pow
}
