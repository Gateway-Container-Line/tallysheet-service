package models

type RackingData []struct {
	WarehouseName  string `gorm:"varchar(30)" json:"warehouse_name"`
	RackingMarunda struct {
		LoadingArea                 *int `gorm:"varchar(5)" json:"loading_area,omitempty"`
		RackingLine1                *int `gorm:"varchar(5)" json:"racking_line_1,omitempty"`
		RackingLine2                *int `gorm:"varchar(5)" json:"racking_line_2,omitempty"`
		RackingLine3                *int `gorm:"varchar(5)" json:"racking_line_3,omitempty"`
		RackingLine4                *int `gorm:"varchar(5)" json:"racking_line_4,omitempty"`
		RackingLine5                *int `gorm:"varchar(5)" json:"racking_line_5,omitempty"`
		RackingLine6                *int `gorm:"varchar(5)" json:"racking_line_6,omitempty"`
		RackingLine7                *int `gorm:"varchar(5)" json:"racking_line_7,omitempty"`
		RackingLine8                *int `gorm:"varchar(5)" json:"racking_line_8,omitempty"`
		RackingLine9                *int `gorm:"varchar(5)" json:"racking_line_9,omitempty"`
		RackingLine10               *int `gorm:"varchar(5)" json:"racking_line_10,omitempty"`
		TemporaryLocationFloor3     *int `gorm:"varchar(5)" json:"temporary_location_floor_3,omitempty"`
		TemporaryLocationFloor4     *int `gorm:"varchar(5)" json:"temporary_location_floor_4,omitempty"`
		TemporaryLocationFloor5     *int `gorm:"varchar(5)" json:"temporary_location_floor_5,omitempty"`
		TemporaryLocationFloor6     *int `gorm:"varchar(5)" json:"temporary_location_floor_6,omitempty"`
		TemporaryLocationFloor7     *int `gorm:"varchar(5)" json:"temporary_location_floor_7,omitempty"`
		TemporaryLocationFloor10    *int `gorm:"varchar(5)" json:"temporary_location_floor_10,omitempty"`
		TemporaryLocationFrontFloor *int `gorm:"varchar(5)" json:"temporary_location_front_floor,omitempty"`
		TemporaryLocationSideFloorA *int `gorm:"varchar(5)" json:"temporary_location_side_floor_a,omitempty"`
		TemporaryLocationSideFloorB *int `gorm:"varchar(5)" json:"temporary_location_side_floor_b,omitempty"`
	} `json:"racking_marunda,omitempty"`
	//racking map[string]int16
}
