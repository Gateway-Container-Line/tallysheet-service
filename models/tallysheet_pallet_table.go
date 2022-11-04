package models

type TallyTable struct {
	IdTable    int `gorm:"primaryKey; autoIncrement"`
	NonPallets `gorm:"embeded"`
	Pallets    `gorm:"embeded"`
}

type Pallets struct {
	//row 1
	PalletCol1Row1 string `gorm:"varchar(5)" json:"pallet_col1row1"`
	PalletCol2Row1 string `gorm:"varchar(5)" json:"pallet_col2row1"`
	PalletCol3Row1 string `gorm:"varchar(5)" json:"pallet_col3row1"`
	PalletCol4Row1 string `gorm:"varchar(5)" json:"pallet_col4row1"`
	PalletCol5Row1 string `gorm:"varchar(5)" json:"pallet_col5row1"`
	PalletCol6Row1 string `gorm:"varchar(5)" json:"pallet_col6row1"`
	PalletCol7Row1 string `gorm:"varchar(5)" json:"pallet_col7row1"`
	PalletCol8Row1 string `gorm:"varchar(5)" json:"pallet_col8row1"`
	PalletCol9Row1 string `gorm:"varchar(5)" json:"pallet_col9row1"`
	//row 2
	PalletCol1Row2 string `gorm:"varchar(5)" json:"pallet_col1row2"`
	PalletCol2Row2 string `gorm:"varchar(5)" json:"pallet_col2row2"`
	PalletCol3Row2 string `gorm:"varchar(5)" json:"pallet_col3row2"`
	PalletCol4Row2 string `gorm:"varchar(5)" json:"pallet_col4row2"`
	PalletCol5Row2 string `gorm:"varchar(5)" json:"pallet_col5row2"`
	PalletCol6Row2 string `gorm:"varchar(5)" json:"pallet_col6row2"`
	PalletCol7Row2 string `gorm:"varchar(5)" json:"pallet_col7row2"`
	PalletCol8Row2 string `gorm:"varchar(5)" json:"pallet_col8row2"`
	PalletCol9Row2 string `gorm:"varchar(5)" json:"pallet_col9row2"`
	//row 3
	PalletCol1Row3 string `gorm:"varchar(5)" json:"pallet_col1row3"`
	PalletCol2Row3 string `gorm:"varchar(5)" json:"pallet_col2row3"`
	PalletCol3Row3 string `gorm:"varchar(5)" json:"pallet_col3row3"`
	PalletCol4Row3 string `gorm:"varchar(5)" json:"pallet_col4row3"`
	PalletCol5Row3 string `gorm:"varchar(5)" json:"pallet_col5row3"`
	PalletCol6Row3 string `gorm:"varchar(5)" json:"pallet_col6row3"`
	PalletCol7Row3 string `gorm:"varchar(5)" json:"pallet_col7row3"`
	PalletCol8Row3 string `gorm:"varchar(5)" json:"pallet_col8row3"`
	PalletCol9Row3 string `gorm:"varchar(5)" json:"pallet_col9row3"`
	//row 4
	PalletCol1Row4 string `gorm:"varchar(5)" json:"pallet_col1row4"`
	PalletCol2Row4 string `gorm:"varchar(5)" json:"pallet_col2row4"`
	PalletCol3Row4 string `gorm:"varchar(5)" json:"pallet_col3row4"`
	PalletCol4Row4 string `gorm:"varchar(5)" json:"pallet_col4row4"`
	PalletCol5Row4 string `gorm:"varchar(5)" json:"pallet_col5row4"`
	PalletCol6Row4 string `gorm:"varchar(5)" json:"pallet_col6row4"`
	PalletCol7Row4 string `gorm:"varchar(5)" json:"pallet_col7row4"`
	PalletCol8Row4 string `gorm:"varchar(5)" json:"pallet_col8row4"`
	PalletCol9Row4 string `gorm:"varchar(5)" json:"pallet_col9row4"`
}

type NonPallets struct {
	//row 1
	NonPalletCol1Row1 int `gorm:"varchar(5)" json:"non_pallet_col1row1"`
	NonPalletCol2Row1 int `gorm:"varchar(5)" json:"non_pallet_col2row1"`
	NonPalletCol3Row1 int `gorm:"varchar(5)" json:"non_pallet_col3row1"`
	NonPalletCol4Row1 int `gorm:"varchar(5)" json:"non_pallet_col4row1"`
	NonPalletCol5Row1 int `gorm:"varchar(5)" json:"non_pallet_col5row1"`
	NonPalletCol6Row1 int `gorm:"varchar(5)" json:"non_pallet_col6row1"`
	NonPalletCol7Row1 int `gorm:"varchar(5)" json:"non_pallet_col7row1"`
	NonPalletCol8Row1 int `gorm:"varchar(5)" json:"non_pallet_col8row1"`
	NonPalletCol9Row1 int `gorm:"varchar(5)" json:"non_pallet_col9row1"`
	//row 2
	NonPalletCol1Row2 int `gorm:"varchar(5)" json:"non_pallet_col1row2"`
	NonPalletCol2Row2 int `gorm:"varchar(5)" json:"non_pallet_col2row2"`
	NonPalletCol3Row2 int `gorm:"varchar(5)" json:"non_pallet_col3row2"`
	NonPalletCol4Row2 int `gorm:"varchar(5)" json:"non_pallet_col4row2"`
	NonPalletCol5Row2 int `gorm:"varchar(5)" json:"non_pallet_col5row2"`
	NonPalletCol6Row2 int `gorm:"varchar(5)" json:"non_pallet_col6row2"`
	NonPalletCol7Row2 int `gorm:"varchar(5)" json:"non_pallet_col7row2"`
	NonPalletCol8Row2 int `gorm:"varchar(5)" json:"non_pallet_col8row2"`
	NonPalletCol9Row2 int `gorm:"varchar(5)" json:"non_pallet_col9row2"`
	//row 3
	NonPalletCol1Row3 int `gorm:"varchar(5)" json:"non_pallet_col1row3"`
	NonPalletCol2Row3 int `gorm:"varchar(5)" json:"non_pallet_col2row3"`
	NonPalletCol3Row3 int `gorm:"varchar(5)" json:"non_pallet_col3row3"`
	NonPalletCol4Row3 int `gorm:"varchar(5)" json:"non_pallet_col4row3"`
	NonPalletCol5Row3 int `gorm:"varchar(5)" json:"non_pallet_col5row3"`
	NonPalletCol6Row3 int `gorm:"varchar(5)" json:"non_pallet_col6row3"`
	NonPalletCol7Row3 int `gorm:"varchar(5)" json:"non_pallet_col7row3"`
	NonPalletCol8Row3 int `gorm:"varchar(5)" json:"non_pallet_col8row3"`
	NonPalletCol9Row3 int `gorm:"varchar(5)" json:"non_pallet_col9row3"`
	//row 4
	NonPalletCol1Row4 int `gorm:"varchar(5)" json:"non_pallet_col1row4"`
	NonPalletCol2Row4 int `gorm:"varchar(5)" json:"non_pallet_col2row4"`
	NonPalletCol3Row4 int `gorm:"varchar(5)" json:"non_pallet_col3row4"`
	NonPalletCol4Row4 int `gorm:"varchar(5)" json:"non_pallet_col4row4"`
	NonPalletCol5Row4 int `gorm:"varchar(5)" json:"non_pallet_col5row4"`
	NonPalletCol6Row4 int `gorm:"varchar(5)" json:"non_pallet_col6row4"`
	NonPalletCol7Row4 int `gorm:"varchar(5)" json:"non_pallet_col7row4"`
	NonPalletCol8Row4 int `gorm:"varchar(5)" json:"non_pallet_col8row4"`
	NonPalletCol9Row4 int `gorm:"varchar(5)" json:"non_pallet_col9row4"`
}

//func (tallytable *TallyTable) AfterUpdate(tx *gorm.DB) (err error) {
//	var tallysheet []TallySheet
//	for i, _ := range tallysheet {
//		DB.Table("tally_sheets").Find(&tallysheet[i].TallyTableIdTable)
//		DB.Table("tally_tables").Where("IdTable = ?", tallysheet[i].TallyTableIdTable).Delete(&tallytable)
//	}
//	return nil
//}
