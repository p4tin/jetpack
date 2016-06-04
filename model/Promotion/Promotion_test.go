package promotion

import (
	"testing"
)

func TestCreatePromotion(t *testing.T) {
	promo := NewPromotion("OrderPromo1", OrderPromotion, 10.0)

	if promo == nil {
		t.Errorf("NewPromotion returned a nil expected a promotion.")
	}

	if promo.PercentOff != 10.0 {
		t.Errorf("NewPromotion PercentOff expected %f, got %f.", promo.PercentOff, 10.0)
	}

	if promo.Name != "OrderPromo1" {
		t.Errorf("NewPromotion PercentOff expected %f, got %f.", promo.Name, "OrderPromo1")
	}

	if promo.Type != OrderPromotion {
		t.Errorf("NewPromotion PercentOff expected %f, got %f.", promo.Type, OrderPromotion)
	}
}
