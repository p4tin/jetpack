package order

import (
	"github.com/p4tin/jetpack/model/Item"
	"github.com/p4tin/jetpack/model/Promotion"
	"testing"
)

func TestCreateOrder(t *testing.T) {
	o := NewOrder()

	if o == nil {
		t.Errorf("NewOrder returned a nil expected an Order!!")
	}

	if o.OrderPriceInfo.Type != "OrderPriceInfo" {
		t.Errorf("NewOrder returned wrong PriceInfo Type, got %s, expected %s", o.OrderPriceInfo.Type, "OrderPriceInfo")
	}
}

func TestAddItemToOrder(t *testing.T) {
	o := NewOrder()

	if o == nil {
		t.Errorf("TestAddItemToOrder returned a nil expected an Order!!")
	}
	o.AddItem(*item.NewItem("1", 1, 0.99))
	o.AddItem(*item.NewItem("2", 2, 1.00))

	if len(o.Items) != 2 {
		t.Errorf("TestAddItemToOrder len returned %d expected %d", len(o.Items), 2)
	}
}

func TestPriceOrder(t *testing.T) {
	o := NewOrder()

	if o == nil {
		t.Errorf("TestPriceOrder returned a nil expected an Order!!")
	}
	o.AddItem(*item.NewItem("1", 1, 0.99))
	o.AddItem(*item.NewItem("2", 2, 1.00))

	o.PriceOrder()

	if o.Items[0].ItemPriceInfo.Amount != 0.99 {
		t.Errorf("TestPriceOrder-Item returned %f expected %f", o.Items[0].ItemPriceInfo.Amount, 0.99)
	}

	if o.Items[1].ItemPriceInfo.Amount != 2.00 {
		t.Errorf("TestPriceOrder-Item returned %f expected %f", o.Items[0].ItemPriceInfo.Amount, 2.00)
	}

	if o.OrderPriceInfo.Amount != 2.99 {
		t.Errorf("TestPriceOrder-Order returned %f expected %f", o.OrderPriceInfo.Amount, 2.99)
	}
}

func TestPriceOrderWithPromo(t *testing.T) {
	promo := promotion.NewPromotion("OrderPromo1", promotion.OrderPromotion, 10.0)
	if promo == nil {
		t.Errorf("TestPriceOrderWithPromo returned a nil expected a promotion.")
	}

	o := NewOrder()
	if o == nil {
		t.Errorf("TestPriceOrderWithPromo returned a nil expected an Order!!")
	}
	o.AddItem(*item.NewItem("1", 1, 0.99))
	o.AddItem(*item.NewItem("2", 2, 1.00))

	o.PriceOrder()
	if o.OrderPriceInfo.Amount != 2.69 {
		t.Errorf("TestPriceOrder-Order returned %f expected %f", o.OrderPriceInfo.Amount, 2.69)
	}
}
