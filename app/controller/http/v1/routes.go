package v1

import (
	"encoding/json"
	"io"
	"net/http"
	"unhashService/app/usecase/hasher"
	"unhashService/pkg/logger"
)

type HashController struct {
	uc     hasher.HasherUC
	logger *logger.Logger
}

func NewHashController(uc hasher.HasherUC, logger *logger.Logger) *HashController {
	return &HashController{
		uc:     uc,
		logger: logger,
	}
}

func (hc *HashController) HashPhoneNumber(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		hc.logger.Error("Bad method")
		return
	}

	w.Header().Add("Content-Type", "application/json")

	body, err := io.ReadAll(r.Body)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(badRequest))
		hc.logger.Error(err.Error())
		return
	}

	var req HashRequest
	err = json.Unmarshal(body, &req)
	if err != nil {
		makeInternalServerError(w, err)
		hc.logger.Error(err.Error())
		return
	}

	phones, err := hc.uc.HashPhoneNumber(req.Hash, req.Domain)
	if err != nil {
		makeBadRequestError(w, err)
		hc.logger.Error(err.Error())
		return
	}
	response := HashResponse{Hashes: phones}
	jsonedResp, err := json.Marshal(response)
	if err != nil {
		makeInternalServerError(w, err)
		hc.logger.Error(err.Error())
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(jsonedResp)

}

func (hc *HashController) UnhashPhoneNumber(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		hc.logger.Error("Bad method")
		return
	}

	w.Header().Add("Content-Type", "application/json")
	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(badRequest))
		hc.logger.Error(err.Error())
		return
	}

	var req UnhashRequest
	err = json.Unmarshal(body, &req)
	if err != nil {
		makeInternalServerError(w, err)
		hc.logger.Error(err.Error())
		return
	}

	phones, err := hc.uc.UnhashPhoneNumber(req.Hash, req.Domain)
	if err != nil {
		makeBadRequestError(w, err)
		hc.logger.Error(err.Error())
		return
	}
	response := UnhashResponse{PhoneNumbers: phones}
	jsonedResp, err := json.Marshal(response)
	if err != nil {
		makeInternalServerError(w, err)
		hc.logger.Error(err.Error())
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(jsonedResp)

}

func (hc *HashController) SetupRoutes() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/hash", hc.HashPhoneNumber)
	mux.HandleFunc("/unhash", hc.UnhashPhoneNumber)
	return mux
}

func (hc *HashController) ListenAndServe(endpoint string, mux *http.ServeMux) error {
	hc.logger.Info("Сервер запущен на " + endpoint)
	return http.ListenAndServe(endpoint, mux)
}
