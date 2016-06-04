package pricingadjustment

import (
	"testing"
)

func TestCreatePricingAdjustment(t *testing.T) {
	adj := NewPricingAdjustment()

	if adj == nil {
		t.Errorf("NewPricingAdjustment returned a nil expected a NewPricingAdjustment!!")
	}
	adj.AddDescription("Test1")
	if adj.AdjustmentDescription != "Test1" {
		t.Errorf("NewPricingAdjustment Description returned %s expected%s", adj.AdjustmentDescription, "Test1")
	}
}
