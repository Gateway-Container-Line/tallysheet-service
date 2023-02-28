package models

import (
	"gorm.io/datatypes"
	"gorm.io/gorm"
	"time"
)

type TallySheet struct {
	gorm.Model
	// index

	// Untuk sementara orm booking code dilakukan embeded yang mana harusnya tdak boleh
	BookingConfirmation `gorm:"embedded"`
	//BookingConfirmation []BookingConfirmation

	//DateTally time.Time `json:"date_tally"`

	TruckNo string `gorm:"varchar(10)" json:"truck_no"` //no seri truck nya

	PartyTally       string    `gorm:"varchar(20)" json:"party_tally"` // quantity + packing
	ContainerNo      string    `gorm:"varchar(50)" json:"container_no"`
	SealNo           string    `gorm:"varchar(50)" json:"seal_no"`
	Size             string    `gorm:"varchar(30)" json:"size"` //size dari container
	StuffingPlanDate time.Time `json:"stuffingplan_date"`       //buat date time

	RackingStatus  string `gorm:"varchar(5)" json:"racking_status"`  // ada 3 condition true false loaded
	GodownLocation uint   `gorm:"varchar(10)" json:"godownlocation"` //no rak buat

	//TallyTableID int
	//TallyTable   *TallyTable `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`

	Condition `gorm:"embedded"`
	//Condition   []Condition

	//MarkingData datatypes.JSONType[MarkingData] `json:"marking_data"`
	//MarkingDataID uint
	//MarkingData   []MarkingData `json:"marking_data,omitempty"`
	//MarkingDataID int16
	MarkingData []*MarkingData `gorm:"foreignKey:BCRefer;association_foreignKey:BCRefer;references:BookingCode;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"marking_data"`

	RackingData datatypes.JSONType[RackingData] `json:"racking_data,omitempty"`

	//jika ada keterangan terhadap tally tulis di status tally
	StatusTally string `gorm:"varchar(30)" json:"status_tally"`

	PICTallyman string `gorm:"varchar(150)" json:"PIC_Tallyman"`
}

type Condition struct {
	//InCondition  `gorm:"embeded"`
	//OutCondition `gorm:"embeded"`

	InCondition  datatypes.JSONType[InCondition]  `json:"in_condition,omitempty"`
	OutCondition datatypes.JSONType[OutCondition] `json:"out_condition,omitempty"`

	//alasan dari kondisi yang ada
	//null jika tidak ada masalah
	AlasanCondition string `gorm:"varchar(50)" json:"alasan_condition,omitempty"`

	//keterangan mengenai surat jalan dan doc eksport
	SignSuratJalan     string `gorm:"varchar(50)" json:"sign_surat_jalan"`
	SignDocumentExport string `gorm:"varchar(50)" json:"sign_document_export"`
}

type InCondition []struct {
	ArrivalNumber       int8  `json:"arrival_number"`
	DateTally           int64 `gorm:"autoCreateTime"`
	DetailedInCondition struct {
		Good   int16 `gorm:"size:30" json:"good"`
		Damage int16 `gorm:"size:30" json:"damage"`
		Short  int16 `gorm:"size:30" json:"short"`
		Over   int16 `gorm:"size:30" json:"over"`
	} `json:"detailed_in_condition"`
	TotalArrivalGoods int16  `gorm:"varchar(30)" json:"total_arrival_goods"`
	ArrivalNotes      string `gorm:"varchar(100)" json:"arrival_notes"`
}

type OutCondition []struct {
	ExitingNumber        int8           `json:"exiting_number"`
	DateTally            datatypes.Date `json:"date_tally_out"`
	TimeTally            datatypes.Time `json:"time_tally_out"`
	DetailedOutCondition struct {
		Good   int16 `gorm:"varchar(10)" json:"good"`
		Damage int16 `gorm:"varchar(30)" json:"damage"`
		Short  int16 `gorm:"varchar(30)" json:"short"`
		Over   int16 `gorm:"varchar(30)" json:"over"`
	} `json:"detailed_out_condition"`
	ExitingStatus struct {
		StatusExit        string `gorm:"varchar(80)" json:"status_exit"`
		ColoadDestination string `gorm:"varchar(100)" json:"coload_destination,omitempty"`
		KeluarSementara   string `gorm:"varchar(100)" json:"keluar_sementara"`
	}
	TotalExitingGoods int16 `gorm:"varchar(30)"`
}

type inputTallySheet struct {
	TallySheet       `gorm:"embedded"`
	DateTally        time.Time `json:"date_tally"`
	StuffingPlanDate time.Time `json:"stuffingplan_date"`
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
