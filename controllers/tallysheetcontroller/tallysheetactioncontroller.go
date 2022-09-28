package tallysheetcontroller

import (
	"encoding/json"
	"github.com/Gateway-Container-Line/tallysheet-service/helper"
	"github.com/Gateway-Container-Line/tallysheet-service/models"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"net/http"
)

func InputTallyForm(w http.ResponseWriter, r *http.Request) {
	//mengambil inputan json yang diterima dari frontend
	var tallyInput models.TallySheet
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&tallyInput); err != nil {
		response := map[string]string{"message": err.Error()}
		helper.ResponseJSON(w, http.StatusBadRequest, response)
		return
	}
	defer r.Body.Close()

	//input data ke database penyimpanan
	if err := models.DB.Create(&tallyInput).Error; err != nil {
		response := map[string]string{"message": err.Error()}
		helper.ResponseJSON(w, http.StatusInternalServerError, response)
		return
	}

	response := map[string]string{"message": "success"}
	helper.ResponseJSON(w, http.StatusOK, response)

	// input id tally to db redis
	//RDB := models.NewRedisClient(0)
	//ttl := time.Duration(24) * time.Hour
	//set2db := RDB.Set(context.Background(), "id_tally", tallyInput.IdTally, ttl)
	//if err := set2db.Err(); err != nil {
	//	logrus.Error("Unable To Set Data to Redis.", err)
	//	return
	//}
	logrus.Info("Berhasil mengirimkan data ke redis DB")
	//booking_code, err := RDB.Get(context.Background(), "booking_code").Result()
	//if err != nil {
	//	logrus.Error("Gagal mendapat Data")
	//}
	//logrus.Info("Berhasil mendapat data. data : ", booking_code)
}

func UpdateTallyForm(w http.ResponseWriter, r *http.Request) {
	paramurl := mux.Vars(r)
	bookingCode := paramurl["booking-code"]

	// mengambil id rak dari url

	logrus.Debug("Berhasil mendapat booking code dari url. data : " + bookingCode)

	// input data ke database
	var tallysheet models.TallySheet

	if models.DB.Where("booking_code = ?", bookingCode).Updates(&tallysheet).RowsAffected == 0 {
		response := map[string]string{"message": "Tallysheet Not Found!"}
		helper.ResponseJSON(w, http.StatusBadRequest, response)
		return
	}
	response := map[string]string{"message": "TallySheet Updated Successfully!"}
	helper.ResponseJSON(w, http.StatusOK, response)

	logrus.Info("Berhasil mengupdate data tallysheet")

}

func DeleteTallySheet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var tallysheet models.TallySheet

	paramurl := mux.Vars(r)
	bookingCode := paramurl["booking-code"]

	if models.DB.Where("booking_code = ?", bookingCode).Delete(&tallysheet).RowsAffected == 0 {
		response := map[string]string{"message": "Tallysheet Not Found!"}
		helper.ResponseJSON(w, http.StatusBadRequest, response)
		return
	}
	response := map[string]string{"message": "TallySheet Deleted Successfully!"}
	helper.ResponseJSON(w, http.StatusOK, response)

	logrus.Info("Berhasil menghapus data dari DB")
}
