package request

type PackageCreateEditModel struct {
	DeliveryAddress string `json:"deliveryAddress"`
	Recipient       string `json:"recipient"`
	Sender          string `json:"sender"`
	Products        []uint `json:"products"`
}
