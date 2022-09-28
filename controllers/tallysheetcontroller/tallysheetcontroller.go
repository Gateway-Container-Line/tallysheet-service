package tallysheetcontroller

import (
	"github.com/Gateway-Container-Line/tallysheet-service/helper"
	"github.com/Gateway-Container-Line/tallysheet-service/models"
	"github.com/gorilla/mux"
	"net/http"
)

func TallySheet(w http.ResponseWriter, r *http.Request) {
	var tallysheet []models.TallySheet
	if models.DB.Find(&tallysheet).RowsAffected == 0 {
		response := map[string]string{"message": "Tidak ada tally sheet"}
		helper.ResponseJSON(w, http.StatusBadRequest, response)
		return
	}
	helper.ResponseJSON(w, http.StatusBadRequest, tallysheet)
}

func TallySheetDetail(w http.ResponseWriter, r *http.Request) {
	paramurl := mux.Vars(r)
	bookingCode := paramurl["booking-code"]
	var tallysheet []models.TallySheet
	if models.DB.Where("booking_code = ?", bookingCode).First(&tallysheet).RowsAffected == 0 {
		response := map[string]string{"message": "Tallysheet Not Found!"}
		helper.ResponseJSON(w, http.StatusBadRequest, response)
		return
	}
	helper.ResponseJSON(w, http.StatusBadRequest, tallysheet)
}
