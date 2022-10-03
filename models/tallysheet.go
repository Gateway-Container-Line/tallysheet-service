package models

import (
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"time"
)

type TallySheet struct {
	IdTally             uint `gorm:"primaryKey;autoIncrement" json:"id_tally"` //buat
	BookingConfirmation `gorm:"embeded"`
	//Booking_Code      string `gorm:"varchar(16)" json:"booking_code"`
	PageTally int64  `gorm:"varchar(10)" json:"page_tally"` //buat
	DateTally string `gorm:"date" json:"date_tally"`        //buat
	TruckNo   int64  `gorm:"varchar(10)" json:"truck_no"`   //buat

	ContainerandSealNo string `gorm:"varchar(50)" json:"container_seal_no"` //buat
	SizeinTally        int8   `gorm:"varchar(30)" json:"size_in_tally"`     //buat
	//Voyage                  string `gorm:"varchar(30)" json:"voyage"`                  //voyage
	StuffingPlanDate string `gorm:"date" json:"stuffingplan_date"`     //buat date time
	GodownLocation   string `gorm:"varchar(10)" json:"godownlocation"` //no rak buat

	DimensionLTally int8 `gorm:"varchar(30)" json:"dimension_l_tally"`
	DimensionWTally int8 `gorm:"varchar(30)" json:"dimension_w_tally"`
	DimensionHTally int8 `gorm:"varchar(30)" json:"dimension_h_tally"`
	QTYTally        int8 `gorm:"varchar(30)" json:"qty_tally"`
	CbmTally        int8 `gorm:"varchar(30)" json:"cbm_Tally"`

	Condition   `gorm:"embeded"`
	StatusTally string `gorm:"varchar(30)" json:"status_tally"`

	SignSuratJalan    string `gorm:"varchar(50)" json:"sign_surat_jalan"`
	SignDokumenExport string `gorm:"varchar(50)" json:"sign_dokumen_export"`

	FullnameTally string `gorm:"varchar(150)" json:"fullname_tally"`
	CreatedAt     string `gorm:"varchar(50)"`
	UpdateAt      string `gorm:"varchar(50)"`
}

type Condition struct {
	InCondition     `gorm:"embeded"`
	OutCondition    `gorm:"embeded"`
	AlasanCondition string `gorm:"varchar(50)" json:"alasan_condition"`
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

func (tallysheet *TallySheet) BeforeUpdate(tx *gorm.DB) (err error) {
	logrus.Info("Test before update")
	//if tx.Statement.Changed() {
	//	tallysheet.UpdateAt = time.Now().Format("02-01-2006 15:04:05 Mon")
	//	//tx.Statement.SetColumn("update_at", time.Now().Format("02-01-2006 15:04:05 Mon"))
	//	tx.Model(&tallysheet).Updates(&tallysheet)
	//}
	//return nil
	tallysheet.UpdateAt = time.Now().Format("02-01-2006 15:04:05 Mon")
	//DB.Model(&tallysheet.UpdateAt).Where("booking_code = ?", &tallysheet.BookingCode).Updates(&tallysheet.UpdateAt)
	DB.Model(&tallysheet.UpdateAt).Save(&tallysheet.UpdateAt)
	return nil
}

func (tallysheet *TallySheet) BeforeCreate(tx *gorm.DB) (err error) {
	logrus.Info("Test Before Create")
	tallysheet.CreatedAt = time.Now().Format("02-01-2006 15:04:05 Mon")
	DB.Model(&tallysheet.CreatedAt).Save(&tallysheet.CreatedAt)
	//DB.Model(&tallysheet.UpdateAt).Where("booking_code = ?", &tallysheet.BookingCode).Updates(&tallysheet.UpdateAt)
	return nil
}
