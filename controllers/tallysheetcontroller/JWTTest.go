package tallysheetcontroller

import (
	"github.com/Gateway-Container-Line/auth-service/helper"
	"net/http"
)

func JWTTest(w http.ResponseWriter, r *http.Request) {
	data := []map[string]any{
		{
			"id":   1,
			"nama": "test1",
		},
		{
			"id":   2,
			"nama": "test2",
		},
		{
			"id":   3,
			"nama": "test3",
		},
	}

	helper.ResponseJSON(w, http.StatusOK, data)
}
