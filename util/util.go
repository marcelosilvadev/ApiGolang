package util

import (
	"encoding/json"
	"erncliente/model"
	"net/http"
	"reflect"
)

//App representa o aplicativo
type App struct {
	Env string
}

//ResponseWithError corresponde
func (a *App) ResponseWithError(w http.ResponseWriter, code int, message string, moreInfo string) {
	var m model.ResponseError
	m.DeveloperMessage = message
	m.UserMessage = "Erro"
	m.ErrorCode = code
	m.MoreInfo = moreInfo
	responseWithError(w, code, m)
}

//ResponseWithJSON corresponde
func (a *App) ResponseWithJSON(w http.ResponseWriter, code int, payload interface{}, limit int, recordCount int) {
	var r model.ResponseSuccess
	r.Records = payload
	lenPayload := reflect.ValueOf(payload)
	r.Meta.RecordCount = 1
	r.Meta.Limit = 1
	if lenPayload.Kind() == reflect.Slice {
		r.Meta.Limit = limit
		r.Meta.Offset = lenPayload.Len()
		r.Meta.RecordCount = recordCount
	}

	response, _ := json.Marshal(r)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func responseWithError(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
