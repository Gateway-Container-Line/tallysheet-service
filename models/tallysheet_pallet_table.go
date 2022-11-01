package models

type TallyTable struct {
	IdTable   int    `gorm:"primaryKey; autoIncrement"`
	NonPallet string `gorm:"varchar(30)" json:"nonPallet"`
	Pallet    `gorm:"embeded"`
}

type Pallet struct {
	//row 1
	Col1Row1 string `gorm:"varchar(5)" json:"col1row1"`
	Col2Row1 string `gorm:"varchar(5)" json:"col2row1"`
	Col3Row1 string `gorm:"varchar(5)" json:"col3row1"`
	Col4Row1 string `gorm:"varchar(5)" json:"col4row1"`
	Col5Row1 string `gorm:"varchar(5)" json:"col5row1"`
	Col6Row1 string `gorm:"varchar(5)" json:"col6row1"`
	Col7Row1 string `gorm:"varchar(5)" json:"col7row1"`
	Col8Row1 string `gorm:"varchar(5)" json:"col8row1"`
	Col9Row1 string `gorm:"varchar(5)" json:"col9row1"`
	//row 2
	Col1Row2 string `gorm:"varchar(5)" json:"col1row2"`
	Col2Row2 string `gorm:"varchar(5)" json:"col2row2"`
	Col3Row2 string `gorm:"varchar(5)" json:"col3row2"`
	Col4Row2 string `gorm:"varchar(5)" json:"col4row2"`
	Col5Row2 string `gorm:"varchar(5)" json:"col5row2"`
	Col6Row2 string `gorm:"varchar(5)" json:"col6row2"`
	Col7Row2 string `gorm:"varchar(5)" json:"col7row2"`
	Col8Row2 string `gorm:"varchar(5)" json:"col8row2"`
	Col9Row2 string `gorm:"varchar(5)" json:"col9row2"`
	//row 3
	Col1Row3 string `gorm:"varchar(5)" json:"col1row3"`
	Col2Row3 string `gorm:"varchar(5)" json:"col2row3"`
	Col3Row3 string `gorm:"varchar(5)" json:"col3row3"`
	Col4Row3 string `gorm:"varchar(5)" json:"col4row3"`
	Col5Row3 string `gorm:"varchar(5)" json:"col5row3"`
	Col6Row3 string `gorm:"varchar(5)" json:"col6row3"`
	Col7Row3 string `gorm:"varchar(5)" json:"col7row3"`
	Col8Row3 string `gorm:"varchar(5)" json:"col8row3"`
	Col9Row3 string `gorm:"varchar(5)" json:"col9row3"`
	//row 4
	Col1Row4 string `gorm:"varchar(5)" json:"col1row4"`
	Col2Row4 string `gorm:"varchar(5)" json:"col2row4"`
	Col3Row4 string `gorm:"varchar(5)" json:"col3row4"`
	Col4Row4 string `gorm:"varchar(5)" json:"col4row4"`
	Col5Row4 string `gorm:"varchar(5)" json:"col5row4"`
	Col6Row4 string `gorm:"varchar(5)" json:"col6row4"`
	Col7Row4 string `gorm:"varchar(5)" json:"col7row4"`
	Col8Row4 string `gorm:"varchar(5)" json:"col8row4"`
	Col9Row4 string `gorm:"varchar(5)" json:"col9row4"`
}

//func (tallytable *TallyTable) AfterUpdate(tx *gorm.DB) (err error) {
//	var tallysheet []TallySheet
//	for i, _ := range tallysheet {
//		DB.Table("tally_sheets").Find(&tallysheet[i].TallyTableIdTable)
//		DB.Table("tally_tables").Where("IdTable = ?", tallysheet[i].TallyTableIdTable).Delete(&tallytable)
//	}
//	return nil
//}
