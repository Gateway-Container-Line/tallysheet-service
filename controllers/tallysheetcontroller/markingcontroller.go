package tallysheetcontroller

import (
	"encoding/json"
	"github.com/Gateway-Container-Line/tallysheet-service/helper"
	"github.com/Gateway-Container-Line/tallysheet-service/models"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm/clause"
	"net/http"
)

func DeleteMarking(w http.ResponseWriter, r *http.Request) {
	//paramurl := mux.Vars(r)
	//bookingCode := paramurl["booking-code"]
	//marking := paramurl["marking"]
	//Getting data from form url encoded
	if err := r.ParseForm(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var bookingCode = r.FormValue("booking-code")
	var marking = r.Form.Get("marking")

	logrus.Info("Berhasil mendapat booking code dan marking dari url. data booking code: " + bookingCode + ", data marking: " + marking)

	var markingData models.MarkingData
	if models.DB.Where("bc_refer = ? AND marking = ?", bookingCode, marking).Select(clause.Associations).Delete(&markingData).RowsAffected == 0 {
		helper.ResponseError(w, http.StatusBadRequest, "Marking Not Found!")
		return
	}
	response := map[string]string{"message": "Marking Deleted Successfully!"}
	helper.ResponseJSON(w, http.StatusOK, response)

	logrus.Info("Success Deleting Marking From TallySheet")

}

func AppendMarking(w http.ResponseWriter, r *http.Request) {
	//paramurl := mux.Vars(r)
	//bookingCode := paramurl["booking-code"]
	//Getting data from form url encoded
	if err := r.ParseForm(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var bookingCode = r.FormValue("booking-code")

	logrus.Info("Berhasil mendapat booking code dari url. data : " + bookingCode)

	// input data to struct tallysheet
	var newMarking models.MarkingData
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&newMarking); err != nil {
		helper.ResponseError(w, http.StatusInternalServerError, err.Error())
		return
	}

	//insert fk
	newMarking.BCRefer = bookingCode
	logrus.Info("marking bc : " + newMarking.BCRefer)

	defer r.Body.Close()

	// find if there is some duplicate data
	var markingSearch models.MarkingData
	if models.DB.Where(
		"bc_refer = ? AND marking = ? AND warehouse_package = ? AND warehouse_dimension_width = ? AND warehouse_dimension_height = ? AND warehouse_dimension_length = ?",
		bookingCode, newMarking.Marking, newMarking.MarkingDetail.WarehousePackage, newMarking.MarkingDetail.WarehouseDimension.WarehouseDimensionWidth,
		newMarking.MarkingDetail.WarehouseDimension.WarehouseDimensionHeight, newMarking.MarkingDetail.WarehouseDimension.WarehouseDimensionLength).First(&markingSearch).RowsAffected == 1 {
		helper.ResponseError(w, http.StatusInternalServerError, "Cannot Insert Duplicate Data Marking")
		return
	}

	////Append
	//var tallySheet models.TallySheet
	//if err := models.DB.Model(&tallySheet).Where("booking_code = ?", bookingCode).Association("MarkingData").Append(&newMarking).Error; err != nil {
	//	helper.ResponseError(w, http.StatusBadRequest, "Tidak Dapat Menambah Marking pada Tallysheet")
	//	return
	//}

	//Create New
	if err := models.DB.Create(&newMarking).Error; err != nil {
		helper.ResponseError(w, http.StatusInternalServerError, err.Error())
		return
	}

	response := map[string]string{"message": "Marking Create Successfully!"}
	helper.ResponseJSON(w, http.StatusOK, response)

	logrus.Info("Berhasil menambahkan marking pada tallysheet dengan booking no: " + bookingCode)

}

func UpdateMarking(w http.ResponseWriter, r *http.Request) {
	//Getting data from form url encoded
	if err := r.ParseForm(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var bookingCode = r.FormValue("booking-code")
	var marking = r.Form.Get("marking")
	var page_tally = r.FormValue("page-tally")

	logrus.Info("Berhasil mendapat data dari url. data booking code : " + bookingCode + ", marking : " + marking + ", page_tally : " + page_tally)

	// input data to struct tallysheet
	var markingData models.MarkingData
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&markingData); err != nil {
		helper.ResponseError(w, http.StatusInternalServerError, err.Error())
		return
	}

	defer r.Body.Close()

	//update
	if models.DB.Where("bc_refer = ? AND Marking = ? AND page_tally = ?", bookingCode, marking, page_tally).Updates(&markingData).RowsAffected == 0 {
		helper.ResponseError(w, http.StatusBadRequest, "Marking Not Found !!")
		return
	}

	response := map[string]string{"message": "Marking Updated Successfully!"}
	helper.ResponseJSON(w, http.StatusOK, response)

	logrus.Info("Berhasil mengupdate data marking")
}
