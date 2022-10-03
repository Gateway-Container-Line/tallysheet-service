package tallysheetcontroller

import (
	"github.com/Gateway-Container-Line/tallysheet-service/helper"
	"github.com/Gateway-Container-Line/tallysheet-service/models"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"net/http"
)

func UpdateGodownLocation(w http.ResponseWriter, r *http.Request) {
	paramurl := mux.Vars(r)
	bookingCode := paramurl["booking-code"]
	rackName := paramurl["rack-name"]

	// mengambil id rak dari url

	logrus.Info("Berhasil mendapat booking code dari url. data : " + bookingCode)

	// input data ke database
	var tallyInput models.TallySheet
	tallyInput.GodownLocation = rackName
	if models.DB.Where("booking_code = ?", bookingCode).Updates(&tallyInput).RowsAffected == 0 {
		response := map[string]string{"message": "Tallysheet Not Found!"}
		helper.ResponseJSON(w, http.StatusBadRequest, response)
		return
	}
	logrus.Info("TallySheet Updated Successfully!")
}
