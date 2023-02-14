package models

type TallyTable struct {
	ID int `gorm:"autoIncrement"`
	//`gorm:"primaryKey; autoIncrement"`
	NonPallets      `gorm:"embeded"`
	Pallets         `gorm:"embeded"`
	DimensionDetail `gorm:"embeded"`
	QtyTallyTable   `gorm:"embeded"`
	MeasTallyTable  `gorm:"embeded"`
}

type DetailedQuantity struct {
	DetailedQty1  int64 `gorm:"varchar(10)" json:"detailed_qty_1"`
	DetailedQty2  int64 `gorm:"varchar(10)" json:"detailed_qty_2"`
	DetailedQty3  int64 `gorm:"varchar(10)" json:"detailed_qty_3"`
	DetailedQty4  int64 `gorm:"varchar(10)" json:"detailed_qty_4"`
	DetailedQty5  int64 `gorm:"varchar(10)" json:"detailed_qty_5"`
	DetailedQty6  int64 `gorm:"varchar(10)" json:"detailed_qty_6"`
	DetailedQty7  int64 `gorm:"varchar(10)" json:"detailed_qty_7"`
	DetailedQty8  int64 `gorm:"varchar(10)" json:"detailed_qty_8"`
	DetailedQty9  int64 `gorm:"varchar(10)" json:"detailed_qty_9"`
	DetailedQty10 int64 `gorm:"varchar(10)" json:"detailed_qty_10"`
}

type DetailedPackage struct {
	DetailedPackage1  string `gorm:"varchar(50)" json:"detailed_package_1"`
	DetailedPackage2  string `gorm:"varchar(50)" json:"detailed_package_2"`
	DetailedPackage3  string `gorm:"varchar(50)" json:"detailed_package_3"`
	DetailedPackage4  string `gorm:"varchar(50)" json:"detailed_package_4"`
	DetailedPackage5  string `gorm:"varchar(50)" json:"detailed_package_5"`
	DetailedPackage6  string `gorm:"varchar(50)" json:"detailed_package_6"`
	DetailedPackage7  string `gorm:"varchar(50)" json:"detailed_package_7"`
	DetailedPackage8  string `gorm:"varchar(50)" json:"detailed_package_8"`
	DetailedPackage9  string `gorm:"varchar(50)" json:"detailed_package_9"`
	DetailedPackage10 string `gorm:"varchar(50)" json:"detailed_package_10"`
}

type Marking struct {
	Marking1  string `gorm:"varchar(50)" json:"marking_1"`
	Marking2  string `gorm:"varchar(50)" json:"marking_2"`
	Marking3  string `gorm:"varchar(50)" json:"marking_3"`
	Marking4  string `gorm:"varchar(50)" json:"marking_4"`
	Marking5  string `gorm:"varchar(50)" json:"marking_5"`
	Marking6  string `gorm:"varchar(50)" json:"marking_6"`
	Marking7  string `gorm:"varchar(50)" json:"marking_7"`
	Marking8  string `gorm:"varchar(50)" json:"marking_8"`
	Marking9  string `gorm:"varchar(50)" json:"marking_9"`
	Marking10 string `gorm:"varchar(50)" json:"marking_10"`
}

type Pallets struct {
	//row 1
	PalletCol1Row1 string `gorm:"varchar(5)" json:"pallet_col1row1"`
	PalletCol2Row1 string `gorm:"varchar(5)" json:"pallet_col2row1"`
	PalletCol3Row1 string `gorm:"varchar(5)" json:"pallet_col3row1"`
	PalletCol4Row1 string `gorm:"varchar(5)" json:"pallet_col4row1"`
	PalletCol5Row1 string `gorm:"varchar(5)" json:"pallet_col5row1"`
	//row 2
	PalletCol1Row2 string `gorm:"varchar(5)" json:"pallet_col1row2"`
	PalletCol2Row2 string `gorm:"varchar(5)" json:"pallet_col2row2"`
	PalletCol3Row2 string `gorm:"varchar(5)" json:"pallet_col3row2"`
	PalletCol4Row2 string `gorm:"varchar(5)" json:"pallet_col4row2"`
	PalletCol5Row2 string `gorm:"varchar(5)" json:"pallet_col5row2"`
	//row 3
	PalletCol1Row3 string `gorm:"varchar(5)" json:"pallet_col1row3"`
	PalletCol2Row3 string `gorm:"varchar(5)" json:"pallet_col2row3"`
	PalletCol3Row3 string `gorm:"varchar(5)" json:"pallet_col3row3"`
	PalletCol4Row3 string `gorm:"varchar(5)" json:"pallet_col4row3"`
	PalletCol5Row3 string `gorm:"varchar(5)" json:"pallet_col5row3"`
	//row 4
	PalletCol1Row4 string `gorm:"varchar(5)" json:"pallet_col1row4"`
	PalletCol2Row4 string `gorm:"varchar(5)" json:"pallet_col2row4"`
	PalletCol3Row4 string `gorm:"varchar(5)" json:"pallet_col3row4"`
	PalletCol4Row4 string `gorm:"varchar(5)" json:"pallet_col4row4"`
	PalletCol5Row4 string `gorm:"varchar(5)" json:"pallet_col5row4"`
	//row 5
	PalletCol1Row5 string `gorm:"varchar(5)" json:"pallet_col1row5"`
	PalletCol2Row5 string `gorm:"varchar(5)" json:"pallet_col2row5"`
	PalletCol3Row5 string `gorm:"varchar(5)" json:"pallet_col3row5"`
	PalletCol4Row5 string `gorm:"varchar(5)" json:"pallet_col4row5"`
	PalletCol5Row5 string `gorm:"varchar(5)" json:"pallet_col5row5"`
	//row 6
	PalletCol1Row6 string `gorm:"varchar(5)" json:"pallet_col1row6"`
	PalletCol2Row6 string `gorm:"varchar(5)" json:"pallet_col2row6"`
	PalletCol3Row6 string `gorm:"varchar(5)" json:"pallet_col3row6"`
	PalletCol4Row6 string `gorm:"varchar(5)" json:"pallet_col4row6"`
	PalletCol5Row6 string `gorm:"varchar(5)" json:"pallet_col5row6"`
	//row 7
	PalletCol1Row7 string `gorm:"varchar(5)" json:"pallet_col1row7"`
	PalletCol2Row7 string `gorm:"varchar(5)" json:"pallet_col2row7"`
	PalletCol3Row7 string `gorm:"varchar(5)" json:"pallet_col3row7"`
	PalletCol4Row7 string `gorm:"varchar(5)" json:"pallet_col4row7"`
	PalletCol5Row7 string `gorm:"varchar(5)" json:"pallet_col5row7"`
	//row 8
	PalletCol1Row8 string `gorm:"varchar(5)" json:"pallet_col1row8"`
	PalletCol2Row8 string `gorm:"varchar(5)" json:"pallet_col2row8"`
	PalletCol3Row8 string `gorm:"varchar(5)" json:"pallet_col3row8"`
	PalletCol4Row8 string `gorm:"varchar(5)" json:"pallet_col4row8"`
	PalletCol5Row8 string `gorm:"varchar(5)" json:"pallet_col5row8"`
	//row 9
	PalletCol1Row9 string `gorm:"varchar(5)" json:"pallet_col1row9"`
	PalletCol2Row9 string `gorm:"varchar(5)" json:"pallet_col2row9"`
	PalletCol3Row9 string `gorm:"varchar(5)" json:"pallet_col3row9"`
	PalletCol4Row9 string `gorm:"varchar(5)" json:"pallet_col4row9"`
	PalletCol5Row9 string `gorm:"varchar(5)" json:"pallet_col5row9"`
	//row 10
	PalletCol1Row10 string `gorm:"varchar(5)" json:"pallet_col1row10"`
	PalletCol2Row10 string `gorm:"varchar(5)" json:"pallet_col2row10"`
	PalletCol3Row10 string `gorm:"varchar(5)" json:"pallet_col3row10"`
	PalletCol4Row10 string `gorm:"varchar(5)" json:"pallet_col4row10"`
	PalletCol5Row10 string `gorm:"varchar(5)" json:"pallet_col5row10"`
}

type NonPallets struct {
	//row 1
	NonPalletCol1Row1 int16 `gorm:"varchar(5)" json:"non_pallet_col1row1"`
	NonPalletCol2Row1 int16 `gorm:"varchar(5)" json:"non_pallet_col2row1"`
	NonPalletCol3Row1 int16 `gorm:"varchar(5)" json:"non_pallet_col3row1"`
	NonPalletCol4Row1 int16 `gorm:"varchar(5)" json:"non_pallet_col4row1"`
	NonPalletCol5Row1 int16 `gorm:"varchar(5)" json:"non_pallet_col5row1"`
	//row 2
	NonPalletCol1Row2 int16 `gorm:"varchar(5)" json:"non_pallet_col1row2"`
	NonPalletCol2Row2 int16 `gorm:"varchar(5)" json:"non_pallet_col2row2"`
	NonPalletCol3Row2 int16 `gorm:"varchar(5)" json:"non_pallet_col3row2"`
	NonPalletCol4Row2 int16 `gorm:"varchar(5)" json:"non_pallet_col4row2"`
	NonPalletCol5Row2 int16 `gorm:"varchar(5)" json:"non_pallet_col5row2"`
	//row 3
	NonPalletCol1Row3 int16 `gorm:"varchar(5)" json:"non_pallet_col1row3"`
	NonPalletCol2Row3 int16 `gorm:"varchar(5)" json:"non_pallet_col2row3"`
	NonPalletCol3Row3 int16 `gorm:"varchar(5)" json:"non_pallet_col3row3"`
	NonPalletCol4Row3 int16 `gorm:"varchar(5)" json:"non_pallet_col4row3"`
	NonPalletCol5Row3 int16 `gorm:"varchar(5)" json:"non_pallet_col5row3"`
	//row 4
	NonPalletCol1Row4 int16 `gorm:"varchar(5)" json:"non_pallet_col1row4"`
	NonPalletCol2Row4 int16 `gorm:"varchar(5)" json:"non_pallet_col2row4"`
	NonPalletCol3Row4 int16 `gorm:"varchar(5)" json:"non_pallet_col3row4"`
	NonPalletCol4Row4 int16 `gorm:"varchar(5)" json:"non_pallet_col4row4"`
	NonPalletCol5Row4 int16 `gorm:"varchar(5)" json:"non_pallet_col5row4"`
	//row 5
	NonPalletCol1Row5 int16 `gorm:"varchar(5)" json:"non_pallet_col1row5"`
	NonPalletCol2Row5 int16 `gorm:"varchar(5)" json:"non_pallet_col2row5"`
	NonPalletCol3Row5 int16 `gorm:"varchar(5)" json:"non_pallet_col3row5"`
	NonPalletCol4Row5 int16 `gorm:"varchar(5)" json:"non_pallet_col4row5"`
	NonPalletCol5Row5 int16 `gorm:"varchar(5)" json:"non_pallet_col5row5"`
	//row 6
	NonPalletCol1Row6 int16 `gorm:"varchar(5)" json:"non_pallet_col1row6"`
	NonPalletCol2Row6 int16 `gorm:"varchar(5)" json:"non_pallet_col2row6"`
	NonPalletCol3Row6 int16 `gorm:"varchar(5)" json:"non_pallet_col3row6"`
	NonPalletCol4Row6 int16 `gorm:"varchar(5)" json:"non_pallet_col4row6"`
	NonPalletCol5Row6 int16 `gorm:"varchar(5)" json:"non_pallet_col5row6"`
	//row 7
	NonPalletCol1Row7 int16 `gorm:"varchar(5)" json:"non_pallet_col1row7"`
	NonPalletCol2Row7 int16 `gorm:"varchar(5)" json:"non_pallet_col2row7"`
	NonPalletCol3Row7 int16 `gorm:"varchar(5)" json:"non_pallet_col3row7"`
	NonPalletCol4Row7 int16 `gorm:"varchar(5)" json:"non_pallet_col4row7"`
	NonPalletCol5Row7 int16 `gorm:"varchar(5)" json:"non_pallet_col5row7"`
	//row 8
	NonPalletCol1Row8 int16 `gorm:"varchar(5)" json:"non_pallet_col1row8"`
	NonPalletCol2Row8 int16 `gorm:"varchar(5)" json:"non_pallet_col2row8"`
	NonPalletCol3Row8 int16 `gorm:"varchar(5)" json:"non_pallet_col3row8"`
	NonPalletCol4Row8 int16 `gorm:"varchar(5)" json:"non_pallet_col4row8"`
	NonPalletCol5Row8 int16 `gorm:"varchar(5)" json:"non_pallet_col5row8"`
	//row 9
	NonPalletCol1Row9 int16 `gorm:"varchar(5)" json:"non_pallet_col1row9"`
	NonPalletCol2Row9 int16 `gorm:"varchar(5)" json:"non_pallet_col2row9"`
	NonPalletCol3Row9 int16 `gorm:"varchar(5)" json:"non_pallet_col3row9"`
	NonPalletCol4Row9 int16 `gorm:"varchar(5)" json:"non_pallet_col4row9"`
	NonPalletCol5Row9 int16 `gorm:"varchar(5)" json:"non_pallet_col5row9"`
	//row 10
	NonPalletCol1Row10 int16 `gorm:"varchar(5)" json:"non_pallet_col1row10"`
	NonPalletCol2Row10 int16 `gorm:"varchar(5)" json:"non_pallet_col2row10"`
	NonPalletCol3Row10 int16 `gorm:"varchar(5)" json:"non_pallet_col3row10"`
	NonPalletCol4Row10 int16 `gorm:"varchar(5)" json:"non_pallet_col4row10"`
	NonPalletCol5Row10 int16 `gorm:"varchar(5)" json:"non_pallet_col5row10"`
}

type DimensionDetail struct {
	// Pallets
	// Row 1
	PalletRow1Length int16 `gorm:"varchar(10)" json:"pallet_row1_length"`
	PalletRow1Width  int16 `gorm:"varchar(10)" json:"pallet_row1_width"`
	PalletRow1Height int16 `gorm:"varchar(10)" json:"pallet_row1_height"`
	// Row 2
	PalletRow2Length int16 `gorm:"varchar(10)" json:"pallet_row2_length"`
	PalletRow2Width  int16 `gorm:"varchar(10)" json:"pallet_row2_width"`
	PalletRow2Height int16 `gorm:"varchar(10)" json:"pallet_row2_height"`
	// Row 3
	PalletRow3Length int16 `gorm:"varchar(10)" json:"pallet_row3_length"`
	PalletRow3Width  int16 `gorm:"varchar(10)" json:"pallet_row3_width"`
	PalletRow3Height int16 `gorm:"varchar(10)" json:"pallet_row3_height"`
	// Row 4
	PalletRow4Length int16 `gorm:"varchar(10)" json:"pallet_row4_length"`
	PalletRow4Width  int16 `gorm:"varchar(10)" json:"pallet_row4_width"`
	PalletRow4Height int16 `gorm:"varchar(10)" json:"pallet_row4_height"`
	// Row 5
	PalletRow5Length int16 `gorm:"varchar(10)" json:"pallet_row5_length"`
	PalletRow5Width  int16 `gorm:"varchar(10)" json:"pallet_row5_width"`
	PalletRow5Height int16 `gorm:"varchar(10)" json:"pallet_row5_height"`
	// Row 6
	PalletRow6Length int16 `gorm:"varchar(10)" json:"pallet_row6_length"`
	PalletRow6Width  int16 `gorm:"varchar(10)" json:"pallet_row6_width"`
	PalletRow6Height int16 `gorm:"varchar(10)" json:"pallet_row6_height"`
	// Row 7
	PalletRow7Length int16 `gorm:"varchar(10)" json:"pallet_row7_length"`
	PalletRow7Width  int16 `gorm:"varchar(10)" json:"pallet_row7_width"`
	PalletRow7Height int16 `gorm:"varchar(10)" json:"pallet_row7_height"`
	// Row 8
	PalletRow8Length int16 `gorm:"varchar(10)" json:"pallet_row8_length"`
	PalletRow8Width  int16 `gorm:"varchar(10)" json:"pallet_row8_width"`
	PalletRow8Height int16 `gorm:"varchar(10)" json:"pallet_row8_height"`
	// Row 9
	PalletRow9Length int16 `gorm:"varchar(10)" json:"pallet_row9_length"`
	PalletRow9Width  int16 `gorm:"varchar(10)" json:"pallet_row9_width"`
	PalletRow9Height int16 `gorm:"varchar(10)" json:"pallet_row9_height"`
	// Row 10
	PalletRow10Length int16 `gorm:"varchar(10)" json:"pallet_row10_length"`
	PalletRow10Width  int16 `gorm:"varchar(10)" json:"pallet_row10_width"`
	PalletRow10Height int16 `gorm:"varchar(10)" json:"pallet_row10_height"`

	// NonPallets
	// Row 1
	NonPalletRow1Length int16 `gorm:"varchar(10)" json:"non_pallet_row1_length"`
	NonPalletRow1Width  int16 `gorm:"varchar(10)" json:"non_pallet_row1_width"`
	NonPalletRow1Height int16 `gorm:"varchar(10)" json:"non_pallet_row1_height"`
	// Row 2
	NonPalletRow2Length int16 `gorm:"varchar(10)" json:"non_pallet_row2_length"`
	NonPalletRow2Width  int16 `gorm:"varchar(10)" json:"non_pallet_row2_width"`
	NonPalletRow2Height int16 `gorm:"varchar(10)" json:"non_pallet_row2_height"`
	// Row 3
	NonPalletRow3Length int16 `gorm:"varchar(10)" json:"non_pallet_row3_length"`
	NonPalletRow3Width  int16 `gorm:"varchar(10)" json:"non_pallet_row3_width"`
	NonPalletRow3Height int16 `gorm:"varchar(10)" json:"non_pallet_row3_height"`
	// Row 4
	NonPalletRow4Length int16 `gorm:"varchar(10)" json:"non_pallet_row4_length"`
	NonPalletRow4Width  int16 `gorm:"varchar(10)" json:"non_pallet_row4_width"`
	NonPalletRow4Height int16 `gorm:"varchar(10)" json:"non_pallet_row4_height"`
	// Row 5
	NonPalletRow5Length int16 `gorm:"varchar(10)" json:"non_pallet_row5_length"`
	NonPalletRow5Width  int16 `gorm:"varchar(10)" json:"non_pallet_row5_width"`
	NonPalletRow5Height int16 `gorm:"varchar(10)" json:"non_pallet_row5_height"`
	// Row 6
	NonPalletRow6Length int16 `gorm:"varchar(10)" json:"non_pallet_row6_length"`
	NonPalletRow6Width  int16 `gorm:"varchar(10)" json:"non_pallet_row6_width"`
	NonPalletRow6Height int16 `gorm:"varchar(10)" json:"non_pallet_row6_height"`
	// Row 7
	NonPalletRow7Length int16 `gorm:"varchar(10)" json:"non_pallet_row7_length"`
	NonPalletRow7Width  int16 `gorm:"varchar(10)" json:"non_pallet_row7_width"`
	NonPalletRow7Height int16 `gorm:"varchar(10)" json:"non_pallet_row7_height"`
	// Row 8
	NonPalletRow8Length int16 `gorm:"varchar(10)" json:"non_pallet_row8_length"`
	NonPalletRow8Width  int16 `gorm:"varchar(10)" json:"non_pallet_row8_width"`
	NonPalletRow8Height int16 `gorm:"varchar(10)" json:"non_pallet_row8_height"`
	// Row 9
	NonPalletRow9Length int16 `gorm:"varchar(10)" json:"non_pallet_row9_length"`
	NonPalletRow9Width  int16 `gorm:"varchar(10)" json:"non_pallet_row9_width"`
	NonPalletRow9Height int16 `gorm:"varchar(10)" json:"non_pallet_row9_height"`
	// Row 10
	NonPalletRow10Length int16 `gorm:"varchar(10)" json:"non_pallet_row10_length"`
	NonPalletRow10Width  int16 `gorm:"varchar(10)" json:"non_pallet_row10_width"`
	NonPalletRow10Height int16 `gorm:"varchar(10)" json:"non_pallet_row10_height"`
}

type QtyTallyTable struct {
	// Pallets
	PalletQTYRow1  int16 `gorm:"varchar(10)" json:"pallet_qty_row1"`
	PalletQTYRow2  int16 `gorm:"varchar(10)" json:"pallet_qty_row2"`
	PalletQTYRow3  int16 `gorm:"varchar(10)" json:"pallet_qty_row3"`
	PalletQTYRow4  int16 `gorm:"varchar(10)" json:"pallet_qty_row4"`
	PalletQTYRow5  int16 `gorm:"varchar(10)" json:"pallet_qty_row5"`
	PalletQTYRow6  int16 `gorm:"varchar(10)" json:"pallet_qty_row6"`
	PalletQTYRow7  int16 `gorm:"varchar(10)" json:"pallet_qty_row7"`
	PalletQTYRow8  int16 `gorm:"varchar(10)" json:"pallet_qty_row8"`
	PalletQTYRow9  int16 `gorm:"varchar(10)" json:"pallet_qty_row9"`
	PalletQTYRow10 int16 `gorm:"varchar(10)" json:"pallet_qty_row10"`

	// Non Pallets
	NonPalletQTYRow1  int16 `gorm:"varchar(10)" json:"non_pallet_qty_row1"`
	NonPalletQTYRow2  int16 `gorm:"varchar(10)" json:"non_pallet_qty_row2"`
	NonPalletQTYRow3  int16 `gorm:"varchar(10)" json:"non_pallet_qty_row3"`
	NonPalletQTYRow4  int16 `gorm:"varchar(10)" json:"non_pallet_qty_row4"`
	NonPalletQTYRow5  int16 `gorm:"varchar(10)" json:"non_pallet_qty_row5"`
	NonPalletQTYRow6  int16 `gorm:"varchar(10)" json:"non_pallet_qty_row6"`
	NonPalletQTYRow7  int16 `gorm:"varchar(10)" json:"non_pallet_qty_row7"`
	NonPalletQTYRow8  int16 `gorm:"varchar(10)" json:"non_pallet_qty_row8"`
	NonPalletQTYRow9  int16 `gorm:"varchar(10)" json:"non_pallet_qty_row9"`
	NonPalletQTYRow10 int16 `gorm:"varchar(10)" json:"non_pallet_qty_row10"`
}

type MeasTallyTable struct {
	// Pallets
	PalletMeasRow1  int16 `gorm:"varchar(10)" json:"pallet_meas_row1"`
	PalletMeasRow2  int16 `gorm:"varchar(10)" json:"pallet_meas_row2"`
	PalletMeasRow3  int16 `gorm:"varchar(10)" json:"pallet_meas_row3"`
	PalletMeasRow4  int16 `gorm:"varchar(10)" json:"pallet_meas_row4"`
	PalletMeasRow5  int16 `gorm:"varchar(10)" json:"pallet_meas_row5"`
	PalletMeasRow6  int16 `gorm:"varchar(10)" json:"pallet_meas_row6"`
	PalletMeasRow7  int16 `gorm:"varchar(10)" json:"pallet_meas_row7"`
	PalletMeasRow8  int16 `gorm:"varchar(10)" json:"pallet_meas_row8"`
	PalletMeasRow9  int16 `gorm:"varchar(10)" json:"pallet_meas_row9"`
	PalletMeasRow10 int16 `gorm:"varchar(10)" json:"pallet_meas_row10"`

	// Non Pallets
	NonPalletMeasRow1  int16 `gorm:"varchar(10)" json:"non_pallet_meas_row1"`
	NonPalletMeasRow2  int16 `gorm:"varchar(10)" json:"non_pallet_meas_row2"`
	NonPalletMeasRow3  int16 `gorm:"varchar(10)" json:"non_pallet_meas_row3"`
	NonPalletMeasRow4  int16 `gorm:"varchar(10)" json:"non_pallet_meas_row4"`
	NonPalletMeasRow5  int16 `gorm:"varchar(10)" json:"non_pallet_meas_row5"`
	NonPalletMeasRow6  int16 `gorm:"varchar(10)" json:"non_pallet_meas_row6"`
	NonPalletMeasRow7  int16 `gorm:"varchar(10)" json:"non_pallet_meas_row7"`
	NonPalletMeasRow8  int16 `gorm:"varchar(10)" json:"non_pallet_meas_row8"`
	NonPalletMeasRow9  int16 `gorm:"varchar(10)" json:"non_pallet_meas_row9"`
	NonPalletMeasRow10 int16 `gorm:"varchar(10)" json:"non_pallet_meas_row10"`
}

//func (tallytable *TallyTable) AfterUpdate(tx *gorm.DB) (err error) {
//	var tallysheet []TallySheet
//	for i, _ := range tallysheet {
//		DB.Table("tally_sheets").Find(&tallysheet[i].TallyTableIdTable)
//		DB.Table("tally_tables").Where("IdTable = ?", tallysheet[i].TallyTableIdTable).Delete(&tallytable)
//	}
//	return nil
//}
