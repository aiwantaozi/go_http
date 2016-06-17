package controller

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func TestPingHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	okRes := map[string]interface{}{"status": http.StatusOK, "code": 1000, "message": "ok"}
	ResderResponseInstance(w, okRes)
}
