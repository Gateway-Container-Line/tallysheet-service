package models

import "github.com/shopspring/decimal"

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
	//  Refer    uint      `gorm:"index:,unique"
}

type MarkingDetail struct {
	WarehousePackage   string             `json:"warehouse_package" gorm:"varchar(30)"`
	WarehouseQuantity  int                `json:"warehouse_quantity" gorm:"varchar(10)"`
	WarehouseMeas      decimal.Decimal    `json:"warehouse_meas" gorm:"type:decimal(18,4);"`
	WarehouseDimension WarehouseDimension `json:"warehouse_dimension" gorm:"embedded"`
}

type WarehouseDimension struct {
	WarehouseDimensionWidth  int `json:"warehouse_dimension_width" gorm:"varchar(30)"`
	WarehouseDimensionHeight int `json:"warehouse_dimension_height" gorm:"varchar(30)"`
	WarehouseDimensionLength int `json:"warehouse_dimension_length" gorm:"varchar(30)"`
}
