package bookingconfirmationcontroller

import (
	"encoding/json"
	"github.com/Gateway-Container-Line/tallysheet-service/models"
	"github.com/sirupsen/logrus"
	"io"
	"log"
	"net/http"
)

func GetBookingConfirmationData(w http.ResponseWriter, r *http.Request) {
	// Health Cek
	//w.Header().Set("Content-Type", "application/json")
	//w.WriteHeader(http.StatusOK)

	// In the future we could report back on the status of our DB, or our cache
	// (e.g. Redis) by performing a simple PING, and include them in the response.
	//io.WriteString(w, `{"alive": true}`)
	logrus.Info("Mengambil API ...")
	quotation, err := http.Get("https://gateway-cl.com/api/quotation_gateway?X-API-KEY=gateway-fms&booking_code=QN-05082022-0001")
	if err != nil {
		log.Fatal(err)
		logrus.Error("Tidak bisa mengambil API!")
	}

	defer quotation.Body.Close()

	quotationData, _ := io.ReadAll(quotation.Body)

	//quotationstring := string(quotationData)
	//fmt.Println(quotationstring)

	var quotationObject models.BookingConfirmation
	json.Unmarshal(quotationData, &quotationObject)

	//fmt.Println(quotationObject.BookingConfirmationData)
	//fmt.Println(quotationObject.BookingConfirmationData[0].BookingDetail[0].BookingCode)

	w.Header().Set("Access-Control-Allow-Origin", "*")
	//w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(quotationData)
}
