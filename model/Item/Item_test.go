package item

import (
	"testing"
)

func TestCreateItem(t *testing.T) {
	itm := NewItem("1", 1, 0.99)

	if itm == nil {
		t.Errorf("NewItems returned a nil expected an Item!!")
	}

	if itm.ItemPriceInfo.Type != "ItemPriceInfo" {
		t.Errorf("NewItems returned wrong PriceInfo Type, got %s, expected %s", itm.ItemPriceInfo.Type, "ItemPriceInfo")
	}

	if itm.Sku != "1" {
		t.Errorf("NewItems Sku returned %s expected %s", itm.Sku, "1")
	}

	if itm.Quantity != 1 {
		t.Errorf("NewItems Sku returned %d expected %d", itm.Quantity, 1)
	}

	if itm.ItemPriceInfo.UnitPrice != 0.99 {
		t.Errorf("NewItems Sku returned %f expected %f", itm.ItemPriceInfo.UnitPrice, 0.99)
	}
}

func TestPriceItemNormal(t *testing.T) {
	itm := NewItem("100001", 2, 0.99)

	if itm == nil {
		t.Errorf("PriceItem returned a nil expected an Item!!")
	}

	itm.PriceItem()
	if itm.ItemPriceInfo.Amount != 49.90 {
		t.Errorf("PriceItem did not price correctly, expected %f, got %f.", 1.98, itm.ItemPriceInfo.Amount)
	}
}

func TestPriceItemZero(t *testing.T) {
	itm := NewItem("100001", 0, 0.99)

	if itm == nil {
		t.Errorf("PriceItem returned a nil expected an Item!!")
	}

	itm.PriceItem()
	if itm.ItemPriceInfo.Amount != 0 {
		t.Errorf("PriceItem did not price correctly, expected %f, got %f.", 0, itm.ItemPriceInfo.Amount)
	}
}
