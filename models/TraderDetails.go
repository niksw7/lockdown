package models

type TraderDetailsRequest struct {
	DeliveryLocation    DeliveryLocation    `json:"deliveryLocation"`
	ShopDetails   ShopDetails   `json:"shopDetails"`
	HomeDeliveryInfo HomeDeliveryInfo `json:"homeDeliveryInfo"`
}

type TraderDetailsDb struct {
	DeliveryLocation    DeliveryLocation
	ShopDetails   ShopDetails
	HomeDeliveryInfo HomeDeliveryInfo
	RegistrationDate    string
	Id                  int
}

type CsvModel struct {
	DeliveryLocation    DeliveryLocation
	ShopDetails   ShopDetails
	HomeDeliveryInfo HomeDeliveryInfo
	ApplicationDate     string
}

type ShopDetails struct {
	Name    string `json:"name"`
	Address string `json:"address"`
	OwnerMobile string `json:"ownerMobile"`
	Email       string `json:"email"`
	Type    string `json:"type"`
}
type HomeDeliveryInfo struct {
	HomeDeliveryNumber string         `json:"homeDeliveryNumber"`
	AgentInfo       AgentInfo   `json:"agentInfo"`
	VehicleInfo     VehicleInfo `json:"vehicleInfo"`
}

type AgentInfo struct {
	AgentName   string `json:"agentName"`
	AgentAge    int    `json:"agentAge"`
	AgentMobile string `json:"agentMobile"`
}

type VehicleInfo struct {
	VehicleType   string `json:"vehicleType"`
	VehicleNumber string `json:"vehicleNumber"`
}
type DeliveryLocation struct {
	Area string `json:"area"`
	City string `json:"city"`
}
