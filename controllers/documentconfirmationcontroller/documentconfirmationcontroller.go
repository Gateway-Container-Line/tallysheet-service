package documentconfirmationcontroller

import (
	"github.com/Gateway-Container-Line/tallysheet-service/helper"
	"github.com/Gateway-Container-Line/tallysheet-service/models"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"net/http"
	"time"
)

func ConfirmationSuratJalan(w http.ResponseWriter, r *http.Request) {
	paramurl := mux.Vars(r)
	bookingCode := paramurl["booking-code"]

	//contoh saja, nanti gampang diganti setelah proses jwt dan middleware
	nama_tallyman := "zakki"

	logrus.Debug("Berhasil mendapat data. data booking code : ", bookingCode)

	var tallysheet models.TallySheet
	tallysheet.SignSuratJalan = time.Now().Format("02-01-2006 15:04:05 Mon ") + nama_tallyman
	//tallysheet.SignSuratJalan = fmt.Sprintf("%s%s", waktu_sekarang.String(), nama_tallyman)
	logrus.Debug("Sign Surat Jalan : ", tallysheet.SignSuratJalan)

	// input ke database
	if models.DB.Where("booking_code = ?", bookingCode).Updates(&tallysheet).RowsAffected == 0 {
		response := map[string]string{"message": "Tidak dapat mengupdate signsuratjalan"}
		helper.ResponseJSON(w, http.StatusBadRequest, response)
		return
	}
	response := map[string]string{"message": "success"}
	helper.ResponseJSON(w, http.StatusOK, response)

	logrus.Info("Berhasil Menyerahkan surat jalan")
}

func ConfirmationDocumentExport(w http.ResponseWriter, r *http.Request) {
	paramurl := mux.Vars(r)
	bookingCode := paramurl["booking-code"]

	//contoh saja, nanti gampang diganti setelah proses jwt dan middleware
	nama_tallyman := "zakki"

	logrus.Debug("Berhasil mendapat data. data booking_code : ", bookingCode)
	var tallysheet models.TallySheet
	tallysheet.SignDocumentExport = time.Now().Format("02-01-2006 15:04:05 Mon ") + nama_tallyman
	logrus.Debug("Sign Surat Jalan : ", tallysheet.SignSuratJalan)

	// input ke database
	if models.DB.Where("booking_code = ?", bookingCode).Updates(&tallysheet).RowsAffected == 0 {
		response := map[string]string{"message": "Tidak dapat mengupdate signdocumentexport"}
		helper.ResponseJSON(w, http.StatusBadRequest, response)
		return
	}
	response := map[string]string{"message": "success"}
	helper.ResponseJSON(w, http.StatusOK, response)

	logrus.Info("Berhasil Menyerahkan dokumen eksport")
}
