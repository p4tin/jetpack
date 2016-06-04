package promotion

type PromotionType string

const (
	OrderPromotion    PromotionType = "OrderPromotion"
	ItemPromotion     PromotionType = "ItemPromotion"
	TaxPromotion      PromotionType = "TaxPromotion"
	ShippingPromotion PromotionType = "ShippingPromotion"
)

type Promotion struct {
	Name       string
	Type       PromotionType
	PercentOff float32
}

var Promotions []Promotion

func NewPromotion(pName string, pType PromotionType, pPercentOff float32) *Promotion {
	promo := new(Promotion)
	promo.Name = pName
	promo.Type = pType
	promo.PercentOff = pPercentOff
	Promotions = append(Promotions, *promo)
	return promo
}
