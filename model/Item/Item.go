package item

import (
	"github.com/p4tin/jetpack/model/PriceInfo"
)

// Item is a basket/shopping cart item that contains the Sku#, the quantity and a PriceInfo that basically has the Sku unit price.
// After the call to PriceItem the PriceInfo will then contain the true discounted price for this shopping cart line item.
type Item struct {
	Sku           string              `json:"type"`
	Quantity      int                 `json:"quantity"`
	ItemPriceInfo priceinfo.PriceInfo `json:"itemPriceInfo"`
}

func NewItem(pSku string, pQuantity int, pPrice float32) *Item {
	itm := new(Item)
	itm.Sku = pSku
	itm.Quantity = pQuantity
	itm.ItemPriceInfo = *priceinfo.NewPriceInfo(pPrice)
	itm.ItemPriceInfo.Type = priceinfo.ItemPriceInfo
	return itm
}

// PriceItem is normally called from the PriceOrder call and returns an item fully priced
// and discounted for item level promotions.
func (i *Item) PriceItem() {
	//Call Pre Calculators here
	i.ItemPriceInfo.Amount = i.ItemPriceInfo.UnitPrice * float32(i.Quantity)
	//Call post Calculators here
	//	log.Printf("ShuId = %s, Item Price = %.2f, quantity = %d, calculatedTotal = %.2f", i.Sku, i.ItemPriceInfo.UnitPrice, i.Quantity, i.ItemPriceInfo.Amount)
}
