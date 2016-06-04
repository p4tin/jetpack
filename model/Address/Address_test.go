package address

import (
	"testing"
)

func TestCreateAddress(t *testing.T) {
	addr := NewAddress()

	if addr == nil {
		t.Errorf("NewAddress returned a nil expected an Address!!")
	}
}
