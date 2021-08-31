package handlers

import (
	"encoding/json"
	"github.com/go-chi/chi"
	"github.com/tsawler/vigilate/internal/models"
	"log"
	"net/http"
	"strconv"
	"time"
)

const (
	HTTP           = 1
	HTTPS          = 2
	SSLCertificate = 3
)


type jsonResp struct {
	OK            bool      `json:"ok"`
	Message       string    `json:"message"`
	ServiceID     int       `json:"service_id"`
	HostServiceID int       `json:"host_service_id"`
	HostID        int       `json:"host_id"`
	OldStatus     string    `json:"old_status"`
	NewStatus     string    `json:"new_status"`
	LastCheck     time.Time `json:"last_check"`
}

// TestCheck manually tests a host service and sends JSON response
func (repo *DBRepo) TestCheck(w http.ResponseWriter, r *http.Request) {
	hostServiceID, _ := strconv.Atoi(chi.URLParam(r, "id"))
	oldStatus := chi.URLParam(r, "oldStatus")

	log.Println(hostServiceID, oldStatus)

	// get host service
	hs, err := repo.DB.GetHostServiceByID(hostServiceID)
	if err != nil {
		log.Println(err)
		return
	}

	log.Println("Service name is", hs.Service.ServiceName)

	// get host
	h, err := repo.DB.GetHostByID(hs.HostID)
	if err != nil {
		log.Println(err)
		return
	}

	// test the service
	repo.testServiceForHost(h, hs)

	// create json
	resp := jsonResp{
		OK:      true,
		Message: "test message",
	}

	// send json to client


	out, _ := json.MarshalIndent(resp, "", "    ")
	w.Header().Set("Content-Type", "application/json")
	w.Write(out)
}

// testServiceForHost tests a service for a host
func (repo *DBRepo) testServiceForHost(h models.Host, hs models.HostService) {

}