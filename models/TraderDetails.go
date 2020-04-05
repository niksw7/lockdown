package models

type TraderDetailsRequest struct {
	Tehsil           string `json:"tehsil"`
	DealerType       string `json:"dealerType"`
	DeliveryLocation string `json:"deliveryLocation"`
	Mobile           string `json:"mobile"`
}

type TraderDetailsDbRequest struct {
	Tehsil           string
	DealerType       string
	DeliveryLocation string
	Mobile           string
	RegistrationDate string
	Id               int
}
