package tallysheetcontroller

//func UpdateGodownLocation(w http.ResponseWriter, r *http.Request) {
//	paramurl := mux.Vars(r)
//	bookingCode := paramurl["booking-code"]
//	rackName := paramurl["rack-name"]
//
//	// mengambil id rak dari url
//
//	logrus.Info("Berhasil mendapat booking code dari url. data : " + bookingCode)
//
//	// input data ke database
//	var tallyInput models.TallySheet
//	//decoder := json.NewDecoder(r.Body)
//	//if err := decoder.Decode(&tallyInput); err != nil {
//	//	response := map[string]string{"message": err.Error()}
//	//	helper.ResponseJSON(w, http.StatusBadRequest, response)
//	//	return
//	//}
//	//defer r.Body.Close()
//	//var tallysheet models.TallySheet
//	//tallysheet = tallyInput
//	if err := r.ParseForm(); err != nil {
//		http.Error(w, err.Error(), http.StatusInternalServerError)
//		return
//	}
//	//var rackName = r.Form.Get("godownlocation")
//	tallyInput.GodownLocation = rackName
//	if models.DB.Where("booking_code = ?", bookingCode).Updates(&tallyInput).RowsAffected == 0 {
//		response := map[string]string{"message": "Tallysheet Not Found!"}
//		helper.ResponseJSON(w, http.StatusBadRequest, response)
//		return
//	}
//	//response := map[string]string{"message": "TallySheet Updated Successfully!"}
//	//helper.ResponseJSON(w, http.StatusOK, response)
//
//	logrus.Info("TallySheet Updated Successfully!")
//}
