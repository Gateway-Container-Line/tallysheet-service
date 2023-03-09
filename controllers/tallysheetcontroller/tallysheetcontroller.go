package tallysheetcontroller

import (
	"encoding/json"
	"github.com/Gateway-Container-Line/tallysheet-service/helper"
	"github.com/Gateway-Container-Line/tallysheet-service/models"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"io"
	"log"
	"net/http"
)

func TallySheet(w http.ResponseWriter, r *http.Request) {
	var tallysheet []models.TallySheet
	if err := models.DB.Preload(clause.Associations).Find(&tallysheet).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			helper.ResponseError(w, http.StatusNotFound, "There is no Tallysheet!")
			return
		default:
			helper.ResponseError(w, http.StatusInternalServerError, err.Error())
			return
		}
	}
	//if models.DB.Preload(clause.Associations).Find(&tallysheet).RowsAffected == 0 {
	//	response := map[string]string{"message": "Tidak ada tally sheet"}
	//	helper.ResponseJSON(w, http.StatusBadRequest, response)
	//	return
	//}
	helper.ResponseJSON(w, http.StatusOK, tallysheet)
}

func TallySheetDetail(w http.ResponseWriter, r *http.Request) {

	paramurl := mux.Vars(r)
	bookingCode := paramurl["booking-code"]
	logrus.Info("BC : " + bookingCode)

	var tallysheet models.TallySheet
	if err := models.DB.Where("booking_code = ?", bookingCode).Preload(clause.Associations).First(&tallysheet).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			helper.ResponseError(w, http.StatusNotFound, "Tallysheet Not Found")
			return
		default:
			helper.ResponseError(w, http.StatusInternalServerError, err.Error())
			return
		}
	}
	//if models.DB.Where("booking_code = ?", bookingCode).Preload(clause.Associations).First(&tallysheet).RowsAffected == 0 {
	//	response := map[string]string{"message": "Tallysheet Not Found!"}
	//	helper.ResponseJSON(w, http.StatusNotFound, response)
	//	return
	//}
	tallysheet.DateTally = helper.TruncateDateText(tallysheet.DateTally, 16)
	helper.ResponseJSON(w, http.StatusOK, tallysheet)
}

func TallyNotInRack(w http.ResponseWriter, r *http.Request) {
	logrus.Info("Bukan BC")
	var tallysheet []models.TallySheet
	if err := models.DB.Where("racking_status = 'false' OR ''").Preload(clause.Associations).Find(&tallysheet).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			helper.ResponseError(w, http.StatusNotFound, "Tallysheet Not Found")
			return
		default:
			helper.ResponseError(w, http.StatusInternalServerError, err.Error())
			return
		}
	}
	//if models.DB.Where("godown_location = ''").Preload(clause.Associations).Find(&tallysheet).RowsAffected == 0 {
	//	response := map[string]string{"message": "Tidak ada tally sheet"}
	//	helper.ResponseJSON(w, http.StatusBadRequest, response)
	//	return
	//}
	helper.ResponseJSON(w, http.StatusOK, tallysheet)
}

type OutputRequestQuoteTally struct {
	statusCode int
	error      error
	data       models.TallySheet
	meta       struct {
		SubmitMethod string
	}
}

func requestGETQuote(result chan<- OutputRequestQuoteTally, bookingCode string) {
	logrus.Info("Request Quote API ...")
	quotation, err := http.Get("https://gateway-cl.com/api/quotation_gateway?X-API-KEY=gateway-fms&booking_code=" + bookingCode)
	if err != nil {
		//close(result)
		log.Fatal(err)
		logrus.Error("Tidak bisa mengambil API!")
	}

	defer quotation.Body.Close()

	quotationData, _ := io.ReadAll(quotation.Body)

	var quotationObject OutputRequestQuoteTally
	json.Unmarshal(quotationData, &quotationObject)

	//statusCode <- quotation.StatusCode
	quotationObject.statusCode = quotation.StatusCode
	quotationObject.meta.SubmitMethod = "POST"
	quotationObject.error = err
	result <- quotationObject
	//error <- err
	//close(result)
	//close(statusCode)
	//close(error)
	//return quotationData,quotation.StatusCode, err
}

func requestGETTally(result chan<- OutputRequestQuoteTally, bookingCode string) {
	logrus.Info("Request Tallysheet API ...")
	tallyRequest, err := http.Get("http://localhost:8081/api/tally-sheet/" + bookingCode)
	if err != nil {
		//close(result)
		//close(statusCode)
		//close(error)
		log.Fatal(err)
		logrus.Error("Tidak bisa mengambil API!")
	}

	defer tallyRequest.Body.Close()

	tallyData, _ := io.ReadAll(tallyRequest.Body)

	//quotationstring := string(quotationData)
	//fmt.Println(quotationstring)

	//var quotationObject models.BookingConfirmation
	//json.Unmarshal(quotationData, &quotationObject)
	var tallysheetObject OutputRequestQuoteTally
	json.Unmarshal(tallyData, &tallysheetObject)

	tallysheetObject.statusCode = tallyRequest.StatusCode
	tallysheetObject.meta.SubmitMethod = "PUT"
	tallysheetObject.error = err
	//statusCode <- tallyRequest.StatusCode
	result <- tallysheetObject
	//error <- err
	//return tallyData, tallyRequest.StatusCode,err
	//close(result)
	//close(statusCode)
	//close(error)
}

func CargoInGETQuoteTally(w http.ResponseWriter, r *http.Request) {
	// Create application context.
	//ctx, cancel := context.WithCancel(context.Background())

	paramurl := mux.Vars(r)
	bookingCode := paramurl["booking-code"]
	logrus.Info("BC : " + bookingCode)

	if bookingCode == "" {
		helper.ResponseError(w, http.StatusBadRequest, "Cannot Empty bookingCode")
	}

	//Create Channel
	resultTally := make(chan OutputRequestQuoteTally, 1)
	resultQuote := make(chan OutputRequestQuoteTally, 1)
	//quotationChanCode := make(chan int)
	//quotationChanError := make(chan error)
	//tallysheetChanResult := make(chan []byte)
	//tallysheetChanCode := make(chan int)
	//tallysheetChanError := make(chan error)

	//requestQuote, errQuote := requestGETQuote(ctx, bookingCode)
	go requestGETTally(resultTally, bookingCode)
	go requestGETQuote(resultQuote, bookingCode)
	//requestTally, errTally := go requestGETTally(ctx, bookingCode)

	var quotationResult = <-resultQuote
	var tallysheetResult = <-resultTally
	//var quotationError = <-quotationChanError
	//var tallysheetResult = <-tallysheetChanResult
	//var tallysheetCode = <-tallysheetChanCode
	//var tallysheetError = <-tallysheetChanError

	if tallysheetResult.statusCode == 200 {
		helper.ResponseJSON(w, http.StatusOK, tallysheetResult)
	} else {
		helper.ResponseJSON(w, http.StatusOK, quotationResult)
	}

	//logrus.Info(quotationCode)
	//logrus.Info(quotationError)
	//logrus.Info(tallysheetError)
	//
	close(resultTally)
	close(resultQuote)
	//close(quotationChanError)
	//close(tallysheetChanResult)
	//close(tallysheetChanCode)
	//close(tallysheetChanError)
	//if quotationError != nil && tallysheetError != nil {
	//	helper.ResponseError(w, http.StatusInternalServerError, "ERR QUOTE :"+utils.ToString(quotationError)+" \n ERR TALLY : "+utils.ToString(tallysheetError))
	//}
	//cancel()
}
