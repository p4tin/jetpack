package priceinfo

import (
	"testing"
)

func TestCreatePriceInfo(t *testing.T) {
	info := NewPriceInfo(3.95)

	if info == nil {
		t.Errorf("NewPriceInfo returned a nil expected a PriceInfo!!")
	}

	if info.UnitPrice != 3.95 {
		t.Errorf("NewPriceInfo UnitPrice expected %f and got %f", 3.95, info.UnitPrice)
	}
}
