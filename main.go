package main

import (
	"github.com/Gateway-Container-Line/tallysheet-service/controllers/bookingconfirmationcontroller"
	"github.com/Gateway-Container-Line/tallysheet-service/controllers/documentconfirmationcontroller"
	"github.com/Gateway-Container-Line/tallysheet-service/controllers/tallysheetcontroller"
	"github.com/Gateway-Container-Line/tallysheet-service/models"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"log"
	"net/http"
)

func main() {
	models.ConnectDatabase()
	logrus.Debug("Server running up...")
	r := mux.NewRouter().StrictSlash(true).UseEncodedPath()

	//Get Data From BookingCode
	r.HandleFunc("/api/quotation-data", bookingconfirmationcontroller.GetBookingConfirmationData).Methods("GET")

	//List all tally
	r.HandleFunc("/api/tally-sheet", tallysheetcontroller.TallySheet).Methods("GET")

	//tally sheet detail
	//r.HandleFunc("/api/tally-sheet",tallysheetcontroller.TallySheetDetail).Methods("GET")
	//router := mux.NewRouter().StrictSlash(true).UseEncodedPath()
	r.HandleFunc("/api/tally-sheet/{booking-code}", tallysheetcontroller.TallySheetDetail).Methods("GET")

	//input tally
	r.HandleFunc("/api/tally-sheet", tallysheetcontroller.InputTallyForm).Methods("POST")

	//update tally
	r.HandleFunc("/api/tally-sheet/{booking-code}", tallysheetcontroller.UpdateTallyForm).Methods("PUT")

	//delete tally
	r.HandleFunc("/api/tally-sheet/{booking-code}", tallysheetcontroller.DeleteTallySheet).Methods("DELETE")

	//update tally surat jalan
	r.HandleFunc("/api/surat-jalan/{booking-code}", documentconfirmationcontroller.ConfirmationSuratJalan).Methods("PUT")

	//update tally doc export
	r.HandleFunc("/api/document-export/{booking-code}", documentconfirmationcontroller.ConfirmationDocumentExport).Methods("PUT")

	log.Fatal(http.ListenAndServe(":8081", r))
}
