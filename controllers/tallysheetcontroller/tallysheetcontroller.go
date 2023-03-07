package tallysheetcontroller

import (
	"github.com/Gateway-Container-Line/tallysheet-service/helper"
	"github.com/Gateway-Container-Line/tallysheet-service/models"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
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
	tallysheet.DateTally = truncateText(tallysheet.DateTally, 16)
	helper.ResponseJSON(w, http.StatusOK, tallysheet)
}

func truncateText(s string, max int) string {
	return s[:max]
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
