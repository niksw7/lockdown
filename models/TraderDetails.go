package models

type TraderDetailsRequest struct {
	City              string              `json:"city"`
	DealerType          string              `json:"dealerType"`
	DeliveryLocation    string              `json:"deliveryLocation"`
	Mobile              string              `json:"mobile"`
	DealerInformation   DealerInformation   `json:"dealerInformation"`
	HomeDeliveryDetails HomeDeliveryDetails `json:"homeDeliveryDetails"`
}

type TraderDetailsDb struct {
	City              string
	DealerType          string
	DeliveryLocation    string
	Mobile              string
	DealerInformation   DealerInformation
	HomeDeliveryDetails HomeDeliveryDetails
	RegistrationDate    string
	Id                  int
}

type CsvModel struct {
	City              string
	DealerType          string
	DeliveryLocation    string
	Mobile              string
	DealerInformation   DealerInformation
	HomeDeliveryDetails HomeDeliveryDetails
	ApplicationDate     string
}

type DealerInformation struct {
	ShopName    string `json:"shopName"`
	ShopAddress string `json:"shopAddress"`
	PhoneNumber string `json:"phoneNumber"`
	Email       string `json:"email"`
	ShopType    string `json:"shopType"`
}
type HomeDeliveryDetails struct {
	HomeDeliveryNumber string         `json:"homeDeliveryNumber"`
	AgentDetails       AgentDetails   `json:"agentDetails"`
	VehicleDetails     VehicleDetails `json:"vehicleDetails"`
}

type AgentDetails struct {
	AgentName   string `json:"agentName"`
	AgentAge    int    `json:"agentAge"`
	AgentMobile string `json:"agentMobile"`
}

type VehicleDetails struct {
	Type   string `json:"type"`
	Number string `json:"number"`
}
