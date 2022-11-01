package tallysheetcontroller

import (
	"github.com/Gateway-Container-Line/tallysheet-service/helper"
	"github.com/Gateway-Container-Line/tallysheet-service/models"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm/clause"
	"net/http"
)

func TallySheet(w http.ResponseWriter, r *http.Request) {
	var tallysheet []models.TallySheet
	if models.DB.Preload(clause.Associations).Find(&tallysheet).RowsAffected == 0 {
		response := map[string]string{"message": "Tidak ada tally sheet"}
		helper.ResponseJSON(w, http.StatusBadRequest, response)
		return
	}
	helper.ResponseJSON(w, http.StatusBadRequest, tallysheet)
}

func TallySheetDetail(w http.ResponseWriter, r *http.Request) {
	//if err := r.ParseForm(); err != nil {
	//	http.Error(w, err.Error(), http.StatusInternalServerError)
	//	return
	//}

	paramurl := mux.Vars(r)
	//paramurl := r.FormValue("booking-code")
	//logrus.Info("Paramurl : " + paramurl)

	//bookingCode := r.URL.Query().Get("booking-code")

	//logrus.Info(paramurl)
	//bookingCode, err := url.QueryUnescape(paramurl)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//bookingCode := EncodeParam(s)
	bookingCode := paramurl["booking-code"]

	logrus.Info("BC : " + bookingCode)
	var tallysheet []models.TallySheet
	if models.DB.Where("booking_code = ?", bookingCode).Preload(clause.Associations).First(&tallysheet).RowsAffected == 0 {
		response := map[string]string{"message": "Tallysheet Not Found!"}
		helper.ResponseJSON(w, http.StatusNotFound, response)
		return
	}
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST")
	w.Header().Set("Content-Type", "application/json")
	helper.ResponseJSON(w, http.StatusOK, tallysheet)
}
