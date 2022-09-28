package models

type BookingDetail struct {
	BookingCode     string `gorm:"varchar(18)" json:"booking_code"`
	DestinationCity string `gorm:"varchar(100)" json:"destination_city"`
	Voyage          string `gorm:"varchar(50)" json:"voyage"`
}

type GoodsDetail struct {
	ShipperName        string `gorm:"varchar(100)" json:"shipper_name"`
	PackageType        string `gorm:"varchar(20)" json:"package_type"`
	Quantity           int64  `gorm:"varchar(10)" json:"quantity"` //party
	Volume             int64  `gorm:"varchar(10)" json:"volume"`   //meas
	Weight             int64  `gorm:"varchar(10)" json:"weight"`
	Marking            string `gorm:"varchar(150)" json:"marking"`
	DescriptionOfGoods string `gorm:"varchar(500)" json:"description_of_goods"`
}

type BookingConfirmationData struct {
	BookingDetail `gorm:"embeded"`
	GoodsDetail   `gorm:"embeded"`
}

type BookingConfirmation struct {
	BookingConfirmationData `gorm:"embeded"`
}
