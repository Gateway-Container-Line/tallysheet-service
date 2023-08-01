package models

// Data From BC API

type BookingDetail struct {
	BookingCode      string `gorm:"unique" json:"booking_code"`
	DestinationCity  string `gorm:"varchar(100)" json:"destination_city"`
	Vessel           string `gorm:"varchar(50)" json:"vessel"`
	ConnectingVessel string `gorm:"varchar(50)" json:"connecting_vessel"`
}

type GoodsDetail struct {
	ShipperName string `gorm:"varchar(100)" json:"shipper_name"`
	PackageType string `gorm:"varchar(50)" json:"package_type"`
	Quantity    int16  `gorm:"size:8" json:"quantity"`    //party: quantity +package type
	Volume      string `gorm:"varchar(10)" json:"volume"` //meas
	Weight      string `gorm:"varchar(10)" json:"weight"`
	//Marking            string  `gorm:"varchar(150)" json:"marking"`
	DescriptionOfGoods string `gorm:"size:800" json:"description_of_goods"`
	DescriptionPrint   string `gorm:"size:200" json:"description_print"`
}

type BookingConfirmation struct {
	BookingDetail `gorm:"embeded"`
	GoodsDetail   `gorm:"embeded"`
}
