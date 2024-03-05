package entities

type Address struct {
	City           string `json:"city"`
	District       string `json:"district"`
	Street         string `json:"street"`
	House          string `json:"house"`
	AdditionalInfo string `json:"additionalInfo"`
}
