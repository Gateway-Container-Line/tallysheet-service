package models

import (
	"database/sql"
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
	//VesselTally      string `gorm:"varchar(30)" json:"vessel_tally"`       //buat
	SizeinTally int8 `gorm:"varchar(30)" json:"size_in_tally"` //buat
	//Voyage                  string `gorm:"varchar(30)" json:"voyage"`                  //voyage
	StuffingPlanDate string         `gorm:"date" json:"stuffingplan_date"`     //buat date time
	GodownLocation   sql.NullString `gorm:"varchar(10)" json:"godownlocation"` //no rak buat

	DimensionLTally int8 `gorm:"varchar(30)" json:"dimension_l_tally"`
	DimensionWTally int8 `gorm:"varchar(30)" json:"dimension_w_tally"`
	DimensionHTally int8 `gorm:"varchar(30)" json:"dimension_h_tally"`
	QTYTally        int8 `gorm:"varchar(30)" json:"qty_tally"`
	CbmTally        int8 `gorm:"varchar(30)" json:"cbm_Tally"`

	Condition   `gorm:"embeded"`
	StatusTally sql.NullString `gorm:"varchar(30)" json:"status_tally"`

	SignSuratJalan    string `gorm:"varchar(50)" json:"sign_surat_jalan"`
	SignDokumenExport string `gorm:"varchar(50)" json:"sign_dokumen_export"`

	FullnameTally string    `gorm:"varchar(150)" json:"fullname_tally"`
	CreatedAt     time.Time `gorm:"date"`
	UpdateAt      time.Time `gorm:"date"`
}

type Condition struct {
	InCondition     `gorm:"embeded"`
	OutCondition    `gorm:"embeded"`
	AlasanCondition sql.NullString `gorm:"varchar(50)" json:"alasan_condition"`
}

type InCondition struct {
	InGood   sql.NullInt16 `gorm:"varchar(10)" json:"condition_in_good"`
	InDamage sql.NullInt16 `gorm:"varchar(30)" json:"condition_in_damage"`
	InShort  sql.NullInt16 `gorm:"varchar(30)" json:"condition_in_short"`
	InOver   sql.NullInt16 `gorm:"varchar(30)" json:"condition_in_over"`
}

type OutCondition struct {
	OutGood   sql.NullInt16 `gorm:"varchar(10)" json:"condition_in_good"`
	OutDamage sql.NullInt16 `gorm:"varchar(30)" json:"condition_in_damage"`
	OutShort  sql.NullInt16 `gorm:"varchar(30)" json:"condition_in_short"`
	OutOver   sql.NullInt16 `gorm:"varchar(30)" json:"condition_in_over"`
}
