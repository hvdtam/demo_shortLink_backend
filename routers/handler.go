package routers

import (
	"encoding/json"
	"net/http"
	"shortlink/helper"
)

func ForbiddenHandler(w http.ResponseWriter, _ *http.Request) {
	ErrorHandler(w, http.StatusForbidden, "Forbidden")
}

func NotFoundHandler(w http.ResponseWriter, _ *http.Request) {
	ErrorHandler(w, http.StatusNotFound, "Not Found")
}

func InternalServerErrorHandler(w http.ResponseWriter, _ *http.Request) {
	ErrorHandler(w, http.StatusInternalServerError, "Internal Server Error")
}

func ErrorHandler(w http.ResponseWriter, code int, message string) {
	w.WriteHeader(code)
	resp := helper.JsonResponse(code, message)
	jsonResp, _ := json.Marshal(resp)
	w.Write(jsonResp)
}
