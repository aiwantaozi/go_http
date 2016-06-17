package controller

import (
	"encoding/json"
	"net/http"
	"project/configs"
)

type JsonResponse struct {
	Status  int    `json:"status"`
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func RenderResponseError(w http.ResponseWriter, responseCode int, responseStatus int, errorMessage string) {
	errorResponse := JsonResponse{
		Status:  responseStatus,
		Code:    responseCode,
		Message: errorMessage,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(responseStatus)
	json.NewEncoder(w).Encode(errorResponse)
}

func ResderResponseInstance(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(data)
}

func ResderResponseCollection(w http.ResponseWriter, r *http.Request, data interface{}, size int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	var href string
	if len(r.URL.RawQuery) != 0 {
		href = r.Host + r.URL.Path + "?" + r.URL.RawQuery
	} else {
		href = r.Host + r.URL.Path
	}
	res := map[string]interface{}{"href": href, "offset": 0, "size": size, "items": data}
	json.NewEncoder(w).Encode(res)
}

func RenderResponseAppError(w http.ResponseWriter, appError *pkgdir.AppError) {
	errorResponse := JsonResponse{
		Status:  appError.Status,
		Code:    appError.Code,
		Message: appError.Error.Error(),
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(appError.Status)
	json.NewEncoder(w).Encode(errorResponse)
}
