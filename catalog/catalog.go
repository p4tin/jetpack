package catalog

type sku struct {
	Id 		string
	Description 	string
	Price 		float64
	SalePrice	float64
}

var items = map[string]*sku{
	"100001": {Id: "100001", Description: "", Price: 24.95, SalePrice: -1.00 },
	"100002": {Id: "100002", Description: "", Price: 5.35, SalePrice: -1.00 },
	"100003": {Id: "100003", Description: "", Price: 106.50, SalePrice: 99.50 },
}


func GetPrice(itemId string) float64 {
	if items[itemId].SalePrice > 0.0 {
		return items[itemId].SalePrice
	} else {
		return items[itemId].Price
	}
}
