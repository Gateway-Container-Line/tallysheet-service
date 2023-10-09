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
	"strconv"
)

// AllTallySheetOutput Output for get All Tallysheet data
type AllTallySheetOutput struct {
	//Code          int
	Error         bool                `json:"error,omitempty"`
	Message       string              `json:"message,omitempty"`
	InventoryData []models.TallySheet `json:"inventory_data"`
	MetaData      models.MetaData     `json:"meta_data,omitempty"`
}

func paginate(r *http.Request) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		query := r.URL.Query()
		page, _ := strconv.Atoi(query.Get("page"))
		if page <= 0 {
			page = 1
		}

		limit, _ := strconv.Atoi(query.Get("limit"))
		switch {
		case limit > 100:
			limit = 100
		case limit <= 0:
			limit = 10
		}

		offset := (page - 1) * limit

		return db.Offset(offset).Limit(limit)
	}
}

func searching(r *http.Request) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		query := r.URL.Query()
		search := query.Get("search")
		if search == "" {
			return db
		} else {
			//return db.Where("* Like %", search)
			//db.Where("booking_code LIKE %?%", search).Or("destination_city LIKE %?%", search).Or("vessel LIKE %?% ", search).
			//	Or("shipper_name LIKE %?%", search).Or("package_type LIKE %?% ", search).Or("quantity LIKE %?%", search).
			//	Or("description_of_goods LIKE %?% ", search).Or("date_tally LIKE %?%", search).Or("party_tally LIKE %?%", search).Or("marking LIKE %?% ", search)
			return db.Joins("LEFT JOIN marking_data ON tally_sheets.booking_code = marking_data.bc_refer").Where("booking_code LIKE ? OR destination_city LIKE ? OR vessel LIKE ? OR shipper_name LIKE ? OR package_type LIKE ? OR quantity LIKE ? OR description_of_goods LIKE ? OR date_tally LIKE ? OR party_tally LIKE ? OR marking_data.marking LIKE ? OR status_tally LIKE ?", "%"+search+"%", "%"+search+"%", "%"+search+"%",
				"%"+search+"%", "%"+search+"%", "%"+search+"%", "%"+search+"%", "%"+search+"%", "%"+search+"%", "%"+search+"%", "%"+search+"%")
		}
	}
}

// TallySheet Search All tallysheet
func TallySheet(w http.ResponseWriter, r *http.Request) {
	models.ConnectDatabase()
	var tallysheet []models.TallySheet
	//var meta models.MetaData

	if err := models.DB.Preload(clause.Associations).Scopes(searching(r), paginate(r)).Find(&tallysheet).Error; err != nil {
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
	//if err := models.DB.Preload(clause.Associations).Find(&tallysheet).Error; err != nil {
	//	switch err {
	//	case gorm.ErrRecordNotFound:
	//		response := map[string]any{"error": true, "message": "There is no Tallysheet! :("}
	//		helper.ResponseError(w, http.StatusNotFound, response)
	//		return
	//	default:
	//		response := map[string]any{"error": true, "message": err.Error()}
	//		helper.ResponseError(w, http.StatusInternalServerError, response)
	//		return
	//	}
	//}

	var TSOutput AllTallySheetOutput
	//TSOutput.Code = http.StatusOK
	TSOutput.Error = false
	TSOutput.InventoryData = tallysheet
	TSOutput.MetaData.Page = TSOutput.MetaData.GetPage(r)
	TSOutput.MetaData.Limit = TSOutput.MetaData.GetLimit(r)
	var ts models.TallySheet
	TSOutput.MetaData.TotalRows = TSOutput.MetaData.GetTotalRows(models.DB, &ts)
	TSOutput.MetaData.TotalPages = TSOutput.MetaData.GetTotalPages(models.DB, &ts, r)

	models.CloseConnection()
	helper.ResponseJSON(w, http.StatusOK, TSOutput)
}

type TallySheetOutput struct {
	Error          bool              `json:"error,omitempty"`
	Message        string            `json:"message,omitempty"`
	TallysheetData models.TallySheet `json:"tallysheet_data,omitempty"`
	MetaData       struct{}          `json:"meta_data,omitempty"`
}

func TallySheetDetail(w http.ResponseWriter, r *http.Request) {
	models.ConnectDatabase()
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
	models.CloseConnection()
	helper.ResponseJSON(w, http.StatusOK, TSOutput)
}

func CheckingTallySheet(bookingCode string) (TallySheetOutput, error) {

	//paramurl := mux.Vars(r)
	//bookingCode := paramurl["booking-code"]
	models.ConnectDatabase()
	bookingCode, _ = url.QueryUnescape(bookingCode)
	logrus.Info("BC : " + bookingCode)
	var TSOutput TallySheetOutput
	var tallysheet models.TallySheet
	if err := models.DB.Where("booking_code = ?", bookingCode).Preload(clause.Associations).First(&tallysheet).Error; err != nil {
		//switch err {
		//case gorm.ErrRecordNotFound:
		//	response := map[string]any{"error": true, "message": "Tallysheet Not Found! :("}
		//	helper.ResponseError(w, http.StatusNotFound, response)
		//	return
		//default:
		//	helper.ResponseError(w, http.StatusInternalServerError, err.Error())
		//	return
		//}
		TSOutput.Error = true
		TSOutput.Message = err.Error()
		return TSOutput, err
		//if err != gorm.ErrRecordNotFound {
		//	TSOutput.Error = true
		//	TSOutput.Message = err.Error()
		//	return TSOutput, err
		//}
	}
	//if models.DB.Where("booking_code = ?", bookingCode).Preload(clause.Associations).First(&tallysheet).RowsAffected == 0 {
	//	response := map[string]string{"message": "Tallysheet Not Found!"}
	//	helper.ResponseJSON(w, http.StatusNotFound, response)
	//	return
	//}
	//tallysheet.Error = false
	tallysheet.ETD = helper.TruncateDateText(tallysheet.ETD, 10)
	tallysheet.DateTally = helper.TruncateDateText(tallysheet.DateTally, 16)
	TSOutput.Error = false
	TSOutput.TallysheetData = tallysheet
	//helper.ResponseJSON(w, http.StatusOK, TSOutput)
	models.CloseConnection()
	return TSOutput, nil
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
