package contactinfo

import (
	"github.com/p4tin/jetpack/model/Address"
)

type ContactInfo struct {
	Address   address.Address `json:"address"`
	HomePhone string          `json:"homePhone"`
	Mobile    string          `json:"mobile"`
	Facebook  string          `json:"fackbook"`
}

func NewContactInfo() *ContactInfo {
	contact := new(ContactInfo)
	return contact
}
