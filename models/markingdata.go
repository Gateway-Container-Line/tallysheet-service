package models

type MarkingData struct {
	//gorm.Model
	ID      int    `gorm:"primaryKey"`
	BCRefer string `gorm:"size:100"`
	//IdTally   uint
	PageTally int16  `json:"page_tally" gorm:"size:30"`
	Marking   string `json:"marking" gorm:"size:150"`

	//embedded
	MarkingDetail MarkingDetail `json:"marking_detail" gorm:"embedded"`

	//Refer         string           `gorm:"index"`
	//has many
	//MarkingDetail []MarkingDetail `json:"marking_detail"`
	//MarkingDetail []*MarkingDetail `json:"marking_detail" gorm:"foreignKey:IdMarkingData;association_foreignKey:IdMarkingData;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	//many2many
	//MarkingDetail []*MarkingDetail `json:"marking_detail" gorm:"many2many:MarkingData_MarkingDetail;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`

	//MarkingDetailID uint             `gorm:"unique;index"`
	//MarkingDetail   []*MarkingDetail `json:"marking_detail,omitempty" gorm:"foreignKey:DetailID;references:MarkingDetailID;association_foreignKey:DetailID"`
	//MarkingDetail []*MarkingDetail `json:"marking_detail,omitempty" gorm:"many2many:MarkingData_MarkingDetail;foreignKey:Marking;joinForeignKey:MarkingDetailID;References:DetailID;joinReferences:MarkingDetail"`
	//foreignKey:Refer;joinForeignKey:UserReferID;References:UserRefer;joinReferences:ProfileRefer"`
	//  Refer    uint      `gorm:"index:,unique"`
	//Marking string `json:"marking" gorm:"varchar(150)"`
	//MarkingDetail []struct {
	//	WarehousePackage   string           `json:"warehouse_package" gorm:"varchar(30)"`
	//	WarehouseQuantity  int              `json:"warehouse_quantity" gorm:"varchar(10)"`
	//	WarehouseMeas      int16            `json:"warehouse_meas" gorm:"varchar(30)"`
	//	WarehouseDimension map[string]int16 `json:"warehouse_dimension"`
	//	//WarehouseDimension struct {
	//		Width  int `json:"warehouse_dimension_width" gorm:"varchar(30)"`
	//		Height int `json:"warehouse_dimension_height" gorm:"varchar(30)"`
	//		Length int `json:"warehouse_dimension_length" gorm:"varchar(30)"`
	//	//} `json:"warehouse_dimension"`
	//} `json:"marking_detail,omitempty"`
}

type MarkingDetail struct {
	//ID            uint `gorm:"primaryKey"`
	//IdMarkingData uint
	//gorm.Model
	//MarkingDetail      *string            `gorm:"primaryKey"`
	WarehousePackage   string             `json:"warehouse_package" gorm:"varchar(30)"`
	WarehouseQuantity  int                `json:"warehouse_quantity" gorm:"varchar(10)"`
	WarehouseMeas      int16              `json:"warehouse_meas" gorm:"varchar(30)"`
	WarehouseDimension WarehouseDimension `json:"warehouse_dimension" gorm:"embedded"`
}

type WarehouseDimension struct {
	WarehouseDimensionWidth  int `json:"warehouse_dimension_width" gorm:"varchar(30)"`
	WarehouseDimensionHeight int `json:"warehouse_dimension_height" gorm:"varchar(30)"`
	WarehouseDimensionLength int `json:"warehouse_dimension_length" gorm:"varchar(30)"`
}
