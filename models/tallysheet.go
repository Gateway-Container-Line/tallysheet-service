package models

import (
	"gorm.io/gorm"
)

type TallySheet struct {
	gorm.Model
	// index
	//IdTally uint `gorm:"primaryKey;autoIncrement" json:"id_tally"`

	// Untuk sementara orm booking code dilakukan embeded yang mana harusnya tdak boleh
	BookingConfirmation `gorm:"embeded"`
	//BookingConfirmation []BookingConfirmation

	//PageTally int64  `gorm:"varchar(10)" json:"page_tally"` //tidak dibutuhkan karena jika ada banyak barang maka dibuat bc baru
	DateTally string `gorm:"date" json:"date_tally"`
	TruckNo   string `gorm:"varchar(10)" json:"truck_no"` //tanyakan lagi bentuk nya spt apa

	PartyTally       string `gorm:"varchar(20)" json:"party_tally"` // quantity + packing
	ContainerNo      string `gorm:"varchar(50)" json:"container_no"`
	SealNo           string `gorm:"varchar(50)" json:"seal_no"`
	Size             string `gorm:"varchar(30)" json:"size"`       //size dari container
	StuffingPlanDate string `gorm:"date" json:"stuffingplan_date"` //buat date time

	RackingStatus  string `gorm:"varchar(5)" json:"racking_status"`  // ada 3 condition true false loaded
	GodownLocation uint   `gorm:"varchar(10)" json:"godownlocation"` //no rak buat

	//DimensionLTally float32 `gorm:"varchar(30)" json:"dimension_l_tally"` //dimension dibuat di table
	//DimensionWTally float32 `gorm:"varchar(30)" json:"dimension_w_tally"`
	//DimensionHTally float32 `gorm:"varchar(30)" json:"dimension_h_tally"`
	//QTYTally        int8 `gorm:"varchar(30)" json:"qty_tally"` // sama dengan quantity yang di booking confirmation hanya saja ini actual datanya
	//CbmTally float32 `gorm:"varchar(30)" json:"cbm_Tally"`  // meas dibuat di table

	Condition `gorm:"embeded"`
	//Condition   []Condition
	TallyTableID int
	TallyTable   *TallyTable `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`

	PICTallyman string `gorm:"varchar(150)" json:"PIC_Tallyman"`
	//CreatedAt   string `gorm:"varchar(50)"`
	//UpdateAt    string `gorm:"varchar(50)"`
}

type Condition struct {
	InCondition  `gorm:"embeded"`
	OutCondition `gorm:"embeded"`

	//alasan dari kondisi yang ada
	//null jika tidak ada masalah
	AlasanCondition string `gorm:"varchar(50)" json:"alasan_condition"`

	//jika ada keterangan terhadap tally tulis di status tally
	StatusTally string `gorm:"varchar(30)" json:"status_tally"`

	//keterangan mengenai surat jalan dan doc eksport
	SignSuratJalan    string `gorm:"varchar(50)" json:"sign_surat_jalan"`
	SignDokumenExport string `gorm:"varchar(50)" json:"sign_dokumen_export"`
}

type InCondition struct {
	InGood   int16 `gorm:"varchar(10)" json:"condition_in_good"`
	InDamage int16 `gorm:"varchar(30)" json:"condition_in_damage"`
	InShort  int16 `gorm:"varchar(30)" json:"condition_in_short"`
	InOver   int16 `gorm:"varchar(30)" json:"condition_in_over"`
}

type OutCondition struct {
	OutGood   int16 `gorm:"varchar(10)" json:"condition_out_good"`
	OutDamage int16 `gorm:"varchar(30)" json:"condition_out_damage"`
	OutShort  int16 `gorm:"varchar(30)" json:"condition_out_short"`
	OutOver   int16 `gorm:"varchar(30)" json:"condition_out_over"`
}

//func (tallysheet *TallySheet) BeforeUpdate(tx *gorm.DB) (err error) {
//	logrus.Info("Test before update")
//	//if tx.Statement.Changed() {
//	//	tallysheet.UpdateAt = time.Now().Format("02-01-2006 15:04:05 Mon")
//	//	//tx.Statement.SetColumn("update_at", time.Now().Format("02-01-2006 15:04:05 Mon"))
//	//	tx.Model(&tallysheet).Updates(&tallysheet)
//	//}
//	//return nil
//	tallysheet.UpdateAt = time.Now().Format("02-01-2006 15:04:05 Mon")
//	//DB.Model(&tallysheet.UpdateAt).Where("booking_code = ?", &tallysheet.BookingCode).Updates(&tallysheet.UpdateAt)
//	DB.Model(&tallysheet.UpdateAt).Save(&tallysheet.UpdateAt)
//	return nil
//}
//
//func (tallysheet *TallySheet) BeforeCreate(tx *gorm.DB) (err error) {
//	logrus.Info("Test Before Create")
//	tallysheet.CreatedAt = time.Now().Format("02-01-2006 15:04:05 Mon")
//	DB.Model(&tallysheet.CreatedAt).Save(&tallysheet.CreatedAt)
//	//DB.Model(&tallysheet.UpdateAt).Where("booking_code = ?", &tallysheet.BookingCode).Updates(&tallysheet.UpdateAt)
//	return nil
//}
