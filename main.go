package main

import (
	"encoding/json"
	"github.com/Gateway-Container-Line/tallysheet-service/controllers/bookingconfirmationcontroller"
	"github.com/Gateway-Container-Line/tallysheet-service/controllers/documentconfirmationcontroller"
	"github.com/Gateway-Container-Line/tallysheet-service/controllers/tallysheetcontroller"
	"github.com/Gateway-Container-Line/tallysheet-service/models"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"github.com/sirupsen/logrus"
	"log"
	"net/http"
)

type Result struct {
	Code    int         `json:"code"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

//func corsMiddleware(next http.Handler) http.Handler {
//	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//		w.Header().Set("Access-Control-Allow-Origin", "*")                                                            // 允许访问所有域，可以换成具体url，注意仅具体url才能带cookie信息
//		w.Header().Add("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token") //header的类型
//		w.Header().Add("Access-Control-Allow-Credentials", "true")                                                    //设置为true，允许ajax异步请求带cookie信息
//		w.Header().Add("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")                             //允许请求方法
//		w.Header().Set("content-type", "application/json;charset=UTF-8")                                              //返回数据格式是json
//		if r.Method == "OPTIONS" {
//			w.WriteHeader(http.StatusNoContent)
//			return
//		}
//		next.ServeHTTP(w, r)
//	})
//}

func main() {
	models.ConnectDatabase()
	logrus.Println("Server running up...")
	r := mux.NewRouter().StrictSlash(true).UseEncodedPath()
	//r.Use(mux.CORSMethodMiddleware(r))
	//r.Use(corsMiddleware)
	//c := cors.New(cors.Options{
	//	AllowedOrigins:   []string{"*"},
	//	AllowCredentials: true,
	//	Debug:            true,
	//})

	r.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)

		res := Result{Code: 404, Message: "Method not found"}
		response, _ := json.Marshal(res)
		w.Write(response)
	})

	r.MethodNotAllowedHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusMethodNotAllowed)

		res := Result{Code: 403, Message: "Method not allowed"}
		response, _ := json.Marshal(res)
		w.Write(response)
	})

	//h := handlers.AllowedHeaders([]string{"*"})
	//m := handlers.AllowedMethods([]string{"*"})
	//o := handlers.AllowedOrigins([]string{"*"})
	//Get Data From BookingCode
	r.HandleFunc("/api/quotation-data", bookingconfirmationcontroller.GetBookingConfirmationData).Methods("GET")

	//List all tally
	r.HandleFunc("/api/tally-sheet", tallysheetcontroller.TallySheet).Methods("GET")

	//tally sheet detail
	//r.HandleFunc("/api/tally-sheet",tallysheetcontroller.TallySheetDetail).Methods("GET")
	//router := mux.NewRouter().StrictSlash(true).UseEncodedPath()
	r.HandleFunc("/api/tally-sheet/{booking-code}", tallysheetcontroller.TallySheetDetail).Methods("GET")

	//tallysheet get quantity
	r.HandleFunc("/api/quantity-tally/{booking-code}", tallysheetcontroller.QuantityTally).Methods("GET")

	//input tally
	r.HandleFunc("/api/tally-sheet", tallysheetcontroller.InputTallyForm).Methods("POST")

	//update tally
	r.HandleFunc("/api/tally-sheet/{booking-code}", tallysheetcontroller.UpdateTallyForm).Methods("PUT")

	//delete tally
	r.HandleFunc("/api/tally-sheet/{booking-code}", tallysheetcontroller.DeleteTallySheet).Methods("DELETE")

	//tallysheet not in rack
	r.HandleFunc("/api/tally-sheet-not-in-rack", tallysheetcontroller.TallyNotInRack).Methods("GET")

	//update tally surat jalan
	r.HandleFunc("/api/surat-jalan/{booking-code}", documentconfirmationcontroller.ConfirmationSuratJalan).Methods("PUT")

	//update tally doc export
	r.HandleFunc("/api/document-export/{booking-code}", documentconfirmationcontroller.ConfirmationDocumentExport).Methods("PUT")

	handler := cors.AllowAll().Handler(r)

	//New(cors.Options{
	//	AllowedOrigins:       []string{"*"},
	//	AllowedMethods:       []string{"POST, GET ,OPTIONS, PATCH, PUT, DELETE"},
	//	AllowedHeaders:       []string{"Content-Type,AccessToken,X-CSRF-Token, Authorization, Token"},
	//	OptionsPassthrough:   true,
	//	OptionsSuccessStatus: http.StatusNoContent,
	//	Debug:                true,
	//})
	log.Fatal(http.ListenAndServe(":8081", handler))
	//handlers.CORS()(r))
}
