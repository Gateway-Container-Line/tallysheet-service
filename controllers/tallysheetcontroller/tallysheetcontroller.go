package tallysheetcontroller

import (
	"encoding/json"
	"github.com/Gateway-Container-Line/tallysheet-service/helper"
	"github.com/Gateway-Container-Line/tallysheet-service/models"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/utils"
	"io"
	"log"
	"net/http"
	"net/url"
	"runtime"
)

// AllTallySheetOutput Output for get All Tallysheet data
type AllTallySheetOutput struct {
	//Code          int
	Error         bool                `json:",omitempty"`
	InventoryData []models.TallySheet `json:"inventory_data"`
	MetaData      struct{}            `json:"meta_data,omitempty"`
}

// TallySheet Search All tallysheet
func TallySheet(w http.ResponseWriter, r *http.Request) {
	var tallysheet []models.TallySheet
	if err := models.DB.Preload(clause.Associations).Find(&tallysheet).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			response := map[string]any{"error": true, "message": "There is no Tallysheet! :("}
			helper.ResponseError(w, http.StatusNotFound, response)
			return
		default:
			response := map[string]any{"error": true, "message": err.Error()}
			helper.ResponseError(w, http.StatusInternalServerError, response)
			return
		}
	}

	var TSOutput AllTallySheetOutput
	//TSOutput.Code = http.StatusOK
	TSOutput.Error = false
	TSOutput.InventoryData = tallysheet
	helper.ResponseJSON(w, http.StatusOK, TSOutput)
}

type TallySheetOutput struct {
	Error          bool              `json:",omitempty"`
	TallysheetData models.TallySheet `json:"tallysheet_data"`
	MetaData       struct{}          `json:"meta_data,omitempty"`
}

func TallySheetDetail(w http.ResponseWriter, r *http.Request) {

	paramurl := mux.Vars(r)
	bookingCode := paramurl["booking-code"]
	bookingCode, _ = url.QueryUnescape(bookingCode)
	logrus.Info("BC : " + bookingCode)

	var tallysheet models.TallySheet
	if err := models.DB.Where("booking_code = ?", bookingCode).Preload(clause.Associations).First(&tallysheet).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			response := map[string]any{"error": true, "message": "Tallysheet Not Found! :("}
			helper.ResponseError(w, http.StatusNotFound, response)
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
	//tallysheet.Error = false
	tallysheet.ETD = helper.TruncateDateText(tallysheet.ETD, 10)
	tallysheet.DateTally = helper.TruncateDateText(tallysheet.DateTally, 16)
	var TSOutput TallySheetOutput
	TSOutput.Error = false
	TSOutput.TallysheetData = tallysheet
	helper.ResponseJSON(w, http.StatusOK, TSOutput)
}

type OutputRequestQuoteTally struct {
	statusCode int
	error      error
	data       models.TallySheet
	meta       struct {
		SubmitMethod string
	}
}

func (OutputRequestQuoteTally) requestGETQuote(bookingCode string) OutputRequestQuoteTally {
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
	//result <- quotationObject
	//error <- err
	//close(result)
	//close(statusCode)
	//close(error)
	return quotationObject
}

func (OutputRequestQuoteTally) requestGETTally(bookingCode string) OutputRequestQuoteTally {
	logrus.Info("Request Tallysheet API ...")
	tallyRequest, err := http.Get("http://localhost:8081/api/tally-sheet/" + bookingCode)
	if err != nil {
		//close(result)
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
	//result <- tallysheetObject
	//error <- err
	return tallysheetObject
	//close(result)
	//close(statusCode)
	//close(error)
}

func CargoInGETQuoteTally(w http.ResponseWriter, r *http.Request) {
	// Create application context.
	//ctx, cancel := context.WithCancel(context.Background())

	runtime.GOMAXPROCS(2)

	paramurl := mux.Vars(r)
	bookingCode := paramurl["booking-code"]
	logrus.Info("BC : " + bookingCode)

	if bookingCode == "" {
		helper.ResponseError(w, http.StatusBadRequest, "Cannot Empty bookingCode")
	}

	//Create Channel
	//resultTally := make(chan OutputRequestQuoteTally)
	//resultQuote := make(chan OutputRequestQuoteTally)
	//quotationChanCode := make(chan int)
	//quotationChanError := make(chan error)
	//tallysheetChanResult := make(chan []byte)
	//tallysheetChanCode := make(chan int)
	//tallysheetChanError := make(chan error)

	Quoteresult := OutputRequestQuoteTally{}.requestGETQuote(bookingCode)
	tallyresult := OutputRequestQuoteTally{}.requestGETTally(bookingCode)
	//requestQuote, errQuote := requestGETQuote(ctx, bookingCode)
	//go requestGETTally(resultTally, w, bookingCode)
	//go requestGETQuote(resultQuote, w, bookingCode)
	//requestTally, errTally := go requestGETTally(ctx, bookingCode)

	//var quotationResult = <-resultQuote
	//var tallysheetResult = <-resultTally
	//var quotationError = <-quotationChanError
	//var tallysheetResult = <-tallysheetChanResult
	//var tallysheetCode = <-tallysheetChanCode
	//var tallysheetError = <-tallysheetChanError

	//if tallysheetResult.statusCode == 200 {
	//	helper.ResponseJSON(w, http.StatusOK, tallysheetResult)
	//} else {
	//	helper.ResponseJSON(w, http.StatusOK, quotationResult)
	//}

	switch tallyresult.statusCode {
	case 200:
		helper.ResponseJSON(w, http.StatusOK, tallyresult)
	default:
		switch Quoteresult.statusCode {
		case 200:
			helper.ResponseJSON(w, http.StatusOK, Quoteresult)
		default:
			helper.ResponseError(w, http.StatusInternalServerError, "Error : "+utils.ToString(Quoteresult.error)+"&& Error"+utils.ToString(tallyresult.error))
		}
	}
	//select {
	//case tallysheetResult := <-resultTally:
	//	if tallysheetResult.statusCode == 200 {
	//		helper.ResponseJSON(w, http.StatusOK, tallysheetResult)
	//	}
	//case quotationResult := <-resultQuote:
	//	if quotationResult.statusCode == 200 {
	//		helper.ResponseJSON(w, http.StatusOK, quotationResult)
	//	}
	//default:
	//	helper.ResponseError(w, http.StatusInternalServerError, "error request")
	//}

	//logrus.Info(quotationCode)
	//logrus.Info(quotationError)
	//logrus.Info(tallysheetError)
	//
	//close(resultTally)
	//close(resultQuote)
	//close(quotationChanError)
	//close(tallysheetChanResult)
	//close(tallysheetChanCode)
	//close(tallysheetChanError)
	//if quotationError != nil && tallysheetError != nil {
	//	helper.ResponseError(w, http.StatusInternalServerError, "ERR QUOTE :"+utils.ToString(quotationError)+" \n ERR TALLY : "+utils.ToString(tallysheetError))
	//}
	//cancel()
}
