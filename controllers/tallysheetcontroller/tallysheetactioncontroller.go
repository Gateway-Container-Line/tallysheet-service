package tallysheetcontroller

import (
	"encoding/json"
	"github.com/Gateway-Container-Line/tallysheet-service/helper"
	"github.com/Gateway-Container-Line/tallysheet-service/models"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"net/http"
)

func InputTallyForm(w http.ResponseWriter, r *http.Request) {
	//mengambil inputan json yang diterima dari frontend
	var tallyInput models.TallySheet
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&tallyInput); err != nil {
		helper.ResponseError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	//input data ke database penyimpanan
	if err := models.DB.Create(&tallyInput).Error; err != nil {
		helper.ResponseError(w, http.StatusInternalServerError, err.Error())
		return
	}

	response := map[string]string{"message": "Success Create Tallysheet"}
	helper.ResponseJSON(w, http.StatusOK, response)

	// input id tally to db redis
	//RDB := models.NewRedisClient(0)
	//ttl := time.Duration(24) * time.Hour
	//set2db := RDB.Set(context.Background(), "id_tally", tallyInput.IdTally, ttl)
	//if err := set2db.Err(); err != nil {
	//	logrus.Error("Unable To Set Data to Redis.", err)
	//	return
	//}
	logrus.Info("Success Create Tallysheet")
}

func UpdateTallyForm(w http.ResponseWriter, r *http.Request) {
	//if r.Method == "OPTIONS" {
	//	w.Header().Set("Access-Control-Allow-Origin", "*")
	//	w.Header().Set("Access-Control-Allow-Methods", "*")
	//	w.Header().Set("Access-Control-Allow-Headers", "*")
	//	w.Write([]byte("allowed"))
	//	return
	//}

	paramurl := mux.Vars(r)
	bookingCode := paramurl["booking-code"]

	//req, _ := http.NewRequest(http.MethodOptions, "http://host.docker.internal:8081/api/tally-sheet/"+bookingCode, nil)
	//resp, err := http.DefaultClient.Do(req)
	////resp, err := client.Do(rackingin)
	//if err != nil {
	//	log.Fatalln(err)
	//}
	//defer resp.Body.Close()
	// mengambil id rak dari url

	logrus.Info("Berhasil mendapat booking code dari url. data : " + bookingCode)

	// input data ke database
	var tallysheet models.TallySheet
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&tallysheet); err != nil {
		helper.ResponseError(w, http.StatusInternalServerError, err.Error())
		return
	}
	defer r.Body.Close()

	if models.DB.Where("booking_code = ?", bookingCode).Session(&gorm.Session{FullSaveAssociations: true}).Updates(&tallysheet).RowsAffected == 0 {
		helper.ResponseError(w, http.StatusBadRequest, "Tidak Dapat Mengupdate Tallysheet")
		return
	}

	tallysheet.BookingCode = bookingCode
	//.Session(&gorm.Session{FullSaveAssociations: true})
	//if models.DB.Model(&tallysheet.TallyTable).Where("IdTable = ?", tallysheet.TallyTableIdTable).Updates(&tallysheet).RowsAffected == 0 {
	//	response := map[string]string{"message": "Tallysheet Not Found!"}
	//	helper.ResponseJSON(w, http.StatusBadRequest, response)
	//	return
	//}

	response := map[string]string{"message": "TallySheet Updated Successfully!"}
	helper.ResponseJSON(w, http.StatusOK, response)

	logrus.Info("Berhasil mengupdate data tallysheet")

}

func DeleteTallySheet(w http.ResponseWriter, r *http.Request) {

	paramurl := mux.Vars(r)
	bookingCode := paramurl["booking-code"]

	var tallysheet models.TallySheet
	if models.DB.Where("booking_code = ?", bookingCode).Select(clause.Associations).Delete(&tallysheet).RowsAffected == 0 {
		helper.ResponseError(w, http.StatusBadRequest, "Tallysheet Not Found!")
		return
	}
	response := map[string]string{"message": "TallySheet Deleted Successfully!"}
	helper.ResponseJSON(w, http.StatusOK, response)

	logrus.Info("Berhasil menghapus data dari DB")
}
