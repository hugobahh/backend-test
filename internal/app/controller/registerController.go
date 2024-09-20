package controller

import (
	"backend-test/internal/app/service"
	"backend-test/pkg/logger"
	"encoding/json"
	"net/http"
)

type RegisterController struct {
	Log     logger.Logger
	Service service.IRegisterService
}

func NewResumeController(resumeService *service.RegisterService, log *logger.Log) *RegisterController {
	return &RegisterController{Log: log,
		Service: resumeService,
	}
}

func (rc *RegisterController) RegisterEntrance(w http.ResponseWriter, r *http.Request) {
	err := rc.Service.RegisterEntrance(r.Context(), "1")
	if err != nil {
		rc.Log.Infof("RegisterServiceEntrance..." + err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	response := map[string]string{"status": "success"}
	w.WriteHeader(http.StatusOK) // 200 OK
	json.NewEncoder(w).Encode(response)
}

func (rc *RegisterController) RegisterExit(w http.ResponseWriter, r *http.Request) {
	err := rc.Service.RegisterExit(r.Context(), "1")
	if err != nil {
		rc.Log.Infof("RegisterServiceEntrance..." + err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	response := map[string]string{"status": "success"}
	w.WriteHeader(http.StatusOK) // 200 OK
	json.NewEncoder(w).Encode(response)
}
