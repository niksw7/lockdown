package models

type TraderDetails struct {
	Tehsil           string `json:"tehsil"`
	DealerType       string `json:"dealerType"`
	DeliveryLocation string `json:"deliveryLocation"`
	Mobile           string `json:"mobile"`
}
