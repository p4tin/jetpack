package contactinfo

import (
	"testing"
)

func TestCreateContactInfo(t *testing.T) {
	contact := NewContactInfo()

	if contact == nil {
		t.Errorf("NewContactInfo returned a nil expected a ContactInfo!!")
	}
}
