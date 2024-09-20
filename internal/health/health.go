package health

import (
	"backend-test/internal/config"
	"backend-test/internal/constants"
	"backend-test/internal/models"
	"encoding/json"
	"net/http"
)

type HealthController struct {
	version string
}

func NewHealthController(configuration *config.Configuration) *HealthController {
	return &HealthController{version: configuration.Version}
}

func (c *HealthController) HealthCheck(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	w.WriteHeader(http.StatusOK)
	enc := json.NewEncoder(w)
	enc.Encode(models.NewHealthCheck(c.version, constants.HealthPass))
}
