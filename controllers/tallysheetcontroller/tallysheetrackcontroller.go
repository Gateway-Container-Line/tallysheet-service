package tallysheetcontroller

import (
	"context"
	"github.com/Gateway-Container-Line/tallysheet-service/helper"
	"github.com/Gateway-Container-Line/tallysheet-service/models"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"net/http"
)

func QuantityTally(w http.ResponseWriter, r *http.Request) {
	paramurl := mux.Vars(r)
	bookingCode := paramurl["booking-code"]

	logrus.Info("Berhasil mendapat booking code dari url. data : " + bookingCode)

	var tallysheet models.TallySheet
	if err := models.DB.Distinct("id", "booking_code", "quantity").Where("booking_code = ?", bookingCode).Preload(clause.Associations).First(&tallysheet).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			helper.ResponseError(w, http.StatusNotFound, "Tallysheet Not Found")
			return
		default:
			helper.ResponseError(w, http.StatusInternalServerError, err.Error())
			return
		}
	}
	response := map[string]interface{}{
		"ID":           tallysheet.ID,
		"booking_code": tallysheet.BookingCode,
		"quantity":     tallysheet.Quantity,
	}
	helper.ResponseJSON(w, http.StatusOK, response)

}

//type TallyNotInRackOutput struct {
//	Error                     bool
//	MarundaListTallyNotInRack []models.TallySheet `json:"marunda_list_tally_not_in_rack"`
//}

type CargoNotInRackOutput struct {
	gorm.Model
	BookingCode     string
	ShipperName     string
	DestinationCity string `json:"Destination"`
	ETD             string
	Quantity        int64
	ItemsReceived   int64
	PackageType     string
	RackingStatus   string `json:"RackStatus"`
	ItemsInRack     *int   `json:"ItemInRack"`
}

//func TallyNotInRack(w http.ResponseWriter, r *http.Request) {
//logrus.Info("GET List Tally not in rack")
//var tallysheet []models.TallySheet
//if err := models.DB.Where("racking_status = 'false'").Preload(clause.Associations).Find(&tallysheet).Error; err != nil {
//	switch err {
//	case gorm.ErrRecordNotFound:
//		helper.ResponseError(w, http.StatusNotFound, "Tallysheet Not Found")
//		return
//	default:
//		helper.ResponseError(w, http.StatusInternalServerError, err.Error())
//		return
//	}
//}
//if models.DB.Where("godown_location = ''").Preload(clause.Associations).Find(&tallysheet).RowsAffected == 0 {
//	response := map[string]string{"message": "Tidak ada tally sheet"}
//	helper.ResponseJSON(w, http.StatusBadRequest, response)
//	return
//}
//helper.ResponseJSON(w, http.StatusOK, tallysheet)
//}

func TallyNotInRack(w http.ResponseWriter, r *http.Request) {
	logrus.Info("GET List Tally not in rack")
	var tallysheet models.TallySheet
	var Output []CargoNotInRackOutput
	if err := models.DB.Model(&tallysheet).Session(&gorm.Session{Context: context.Background()}).Where("racking_status = 'false' AND items_received <> items_in_rack").Find(&Output).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			helper.ResponseError(w, http.StatusNotFound, "There was no record cargo not in rack")
			return
		default:
			helper.ResponseError(w, http.StatusInternalServerError, err.Error())
			return
		}
	}
	helper.ResponseJSON(w, http.StatusOK, Output)
}

func TallyNotInRackList() ([]CargoNotInRackOutput, error) {
	logrus.Info("GET List Tally not in rack")
	models.ConnectDatabase()
	var tallysheet models.TallySheet
	var Output []CargoNotInRackOutput
	if err := models.DB.Model(&tallysheet).Session(&gorm.Session{Context: context.Background()}).Where("racking_status = 'false' AND items_received <> items_in_rack").Find(&Output).Error; err != gorm.ErrRecordNotFound {
		//switch err {
		//case gorm.ErrRecordNotFound:
		//	//helper.ResponseError(w, http.StatusNotFound, "There was no record cargo not in rack")
		//	return nil, err
		//default:
		//	//helper.ResponseError(w, http.StatusInternalServerError, err.Error())
		//	return nil, err
		//}
		//if err != gorm.ErrRecordNotFound {
		return Output, err
		//}
	}
	//helper.ResponseJSON(w, http.StatusOK, Output)
	return Output, nil
}

//func UpdateGodownLocation(w http.ResponseWriter, r *http.Request) {
//	paramurl := mux.Vars(r)
//	bookingCode := paramurl["booking-code"]
//	rackName := paramurl["rack-name"]
//
//	// mengambil id rak dari url
//
//	logrus.Info("Berhasil mendapat booking code dari url. data : " + bookingCode)
//
//	// input data ke database
//	var tallyInput models.TallySheet
//	//decoder := json.NewDecoder(r.Body)
//	//if err := decoder.Decode(&tallyInput); err != nil {
//	//	response := map[string]string{"message": err.Error()}
//	//	helper.ResponseJSON(w, http.StatusBadRequest, response)
//	//	return
//	//}
//	//defer r.Body.Close()
//	//var tallysheet models.TallySheet
//	//tallysheet = tallyInput
//	if err := r.ParseForm(); err != nil {
//		http.Error(w, err.Error(), http.StatusInternalServerError)
//		return
//	}
//	//var rackName = r.Form.Get("godownlocation")
//	tallyInput.GodownLocation = rackName
//	if models.DB.Where("booking_code = ?", bookingCode).Updates(&tallyInput).RowsAffected == 0 {
//		response := map[string]string{"message": "Tallysheet Not Found!"}
//		helper.ResponseJSON(w, http.StatusBadRequest, response)
//		return
//	}
//	//response := map[string]string{"message": "TallySheet Updated Successfully!"}
//	//helper.ResponseJSON(w, http.StatusOK, response)
//
//	logrus.Info("TallySheet Updated Successfully!")
//}
