package main

import (
	"encoding/json"
	"github.com/Gateway-Container-Line/auth-service/middlewares"
	_ "github.com/Gateway-Container-Line/auth-service/models"
	"github.com/Gateway-Container-Line/tallysheet-service/controllers/admincontroller"
	"github.com/Gateway-Container-Line/tallysheet-service/controllers/documentconfirmationcontroller"
	"github.com/Gateway-Container-Line/tallysheet-service/controllers/tallysheetcontroller"
	"github.com/Gateway-Container-Line/tallysheet-service/models"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"github.com/sirupsen/logrus"
	"log"
	"net/http"
	"os"
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
	//modelTally.ConnectDatabase()
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
	//r.HandleFunc("/api/quotation-data", bookingconfirmationcontroller.GetBookingConfirmationData).Methods("GET")
	//r.HandleFunc("/api/scan/in/{booking-code}", bookingconfirmationcontroller.GetBookingConfirmationData).Methods("GET")

	tallySheetRoute := r.PathPrefix("/api").Subrouter()
	//List all tally
	//r.HandleFunc("/api/tally-sheet", tallysheetcontroller.TallySheet).Methods("GET")
	tallySheetRoute.HandleFunc("/tally-sheet", tallysheetcontroller.TallySheet).Methods("GET")

	//tally sheet detail
	//r.HandleFunc("/api/tally-sheet",tallysheetcontroller.TallySheetDetail).Methods("GET")
	//router := mux.NewRouter().StrictSlash(true).UseEncodedPath()
	//r.HandleFunc("/api/tally-sheet/{booking-code}", tallysheetcontroller.TallySheetDetail).Methods("GET")
	tallySheetRoute.HandleFunc("/tally-sheet/{booking-code}", tallysheetcontroller.TallySheetDetail).Methods("GET")

	//tallysheet get quantity
	//r.HandleFunc("/api/quantity-tally/{booking-code}", tallysheetcontroller.QuantityTally).Methods("GET")

	//input tally
	tallySheetRoute.HandleFunc("/tally-sheet", tallysheetcontroller.InputTallyForm).Methods("POST")
	//r.HandleFunc("/api/tally-sheet", tallysheetcontroller.InputTallyForm).Methods("POST")

	//update tally
	tallySheetRoute.HandleFunc("/tally-sheet/{booking-code}", tallysheetcontroller.UpdateTallyForm).Methods("PUT")
	//r.HandleFunc("/api/tally-sheet/{booking-code}", tallysheetcontroller.UpdateTallyForm).Methods("PUT")

	//delete tally
	tallySheetRoute.HandleFunc("/tally-sheet/{booking-code}", tallysheetcontroller.DeleteTallySheet).Methods("DELETE")
	//r.HandleFunc("/api/tally-sheet/{booking-code}", tallysheetcontroller.DeleteTallySheet).Methods("DELETE")

	//delete marking
	tallySheetRoute.HandleFunc("/tally-sheet/marking/delete", tallysheetcontroller.DeleteMarking).Methods("DELETE")
	//r.HandleFunc("/api/tally-sheet/marking/delete", tallysheetcontroller.DeleteMarking).Methods("DELETE")
	//r.HandleFunc("/api/tally-sheet/marking/{booking-code}/{marking}", tallysheetcontroller.DeleteMarking).Methods("DELETE")

	//append Marking
	tallySheetRoute.HandleFunc("/tally-sheet/marking/append", tallysheetcontroller.AppendMarking).Methods("POST")
	//r.HandleFunc("/api/tally-sheet/marking/append", tallysheetcontroller.AppendMarking).Methods("POST")
	//r.HandleFunc("/api/tally-sheet/marking/{booking-code}", tallysheetcontroller.AppendMarking).Methods("PUT")

	//update Marking
	tallySheetRoute.HandleFunc("/tally-sheet/marking/update", tallysheetcontroller.UpdateMarking).Methods("PUT")
	//r.HandleFunc("/api/tally-sheet/marking/update", tallysheetcontroller.UpdateMarking).Methods("PUT")

	tallySheetRoute.Use(middlewares.JWTMiddleware)

	//tallysheet not in rack
	//r.HandleFunc("/api/tally-sheet-not-in-rack", tallysheetcontroller.TallyNotInRack).Methods("GET")

	//update tally surat jalan
	r.HandleFunc("/api/surat-jalan/{booking-code}", documentconfirmationcontroller.ConfirmationSuratJalan).Methods("PUT")

	//update tally doc export
	r.HandleFunc("/api/document-export/{booking-code}", documentconfirmationcontroller.ConfirmationDocumentExport).Methods("PUT")

	//CargoInGETQuoteTally
	//r.HandleFunc("/api/scan/in/{booking-code}", tallysheetcontroller.CargoInGETQuoteTally).Methods("GET")

	//Count All Tally
	r.HandleFunc("/count", admincontroller.CountAllContent).Methods("GET")

	//r.HandleFunc("/api/count/tally-sheet", admincontroller.CountTallySheet).Methods("GET")
	//Count All Cargo In
	//r.HandleFunc("/api/count/cargo-in", admincontroller.CountCargoIn).Methods("GET")
	//Count All Cargo Out
	//r.HandleFunc("/api/count/cargo-out", admincontroller.CountCargoOut).Methods("GET")
	//Count All Cargo Coloaded
	//r.HandleFunc("/api/count/cargo-coloaded", admincontroller.CargoCoload).Methods("GET")
	//Count Tally In Rack
	r.HandleFunc("/api/count/cargoinrack", admincontroller.CountCargoInRack).Methods("GET")
	//Count Cargo Loaded in Container
	//r.HandleFunc("/api/count/cargoloaded", admincontroller.CountCargoLoadedInContainer).Methods("GET")

	testMid := r.PathPrefix("/test").Subrouter()
	routeFunc := middlewares.Chain(tallysheetcontroller.JWTTest, middlewares.AuthFunc())
	//testMid.HandleFunc("/get-tally-sheet", jwtauthcontroller.JWTTest).Methods("GET")
	//testMid.HandleFunc("/get-tally-sheet", tallysheetcontroller.JWTTest).Methods("GET")
	testMid.HandleFunc("/get-tally-sheet", routeFunc).Methods("GET")
	//testMid.Use(middlewares.JWTMiddleware)

	testmid2 := r.PathPrefix("/test2").Subrouter()
	testmid2.HandleFunc("/get-tally-sheet", tallysheetcontroller.TallySheet).Methods(http.MethodGet)
	//testmid2.HandleFunc("/get-tally-sheet", tallysheetcontroller.JWTTest).Methods(http.MethodGet)
	testmid2.Use(middlewares.JWTMiddleware)

	//handler := cors.AllowAll().Handler(r)

	handler := cors.New(cors.Options{
		AllowedOrigins: []string{os.Getenv("ORIGIN_PROXY")},
		//AllowedOrigins: []string{os.Getenv("ORIGIN_FRONTEND")},
		AllowedMethods: []string{
			http.MethodHead,
			http.MethodGet,
			http.MethodPost,
			http.MethodPut,
			http.MethodPatch,
			http.MethodDelete,
			http.MethodOptions,
		},
		AllowedHeaders:   []string{"Content-Type", "Set-Cookie", "Authorization", " Origin", "X-Requested-With", "Accept"},
		AllowCredentials: true,
	}).Handler(r)

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
