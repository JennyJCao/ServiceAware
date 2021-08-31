package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/tsawler/vigilate/internal/models"
	"log"
	"net/http"
	"strconv"
	"strings"
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

// ScheduledCheck performs a scheduled check on a host service by id
func (repo *DBRepo) ScheduledCheck(hostServiceID int) {

}

// TestCheck manually tests a host service and sends JSON response
func (repo *DBRepo) TestCheck(w http.ResponseWriter, r *http.Request) {
	hostServiceID, _ := strconv.Atoi(chi.URLParam(r, "id"))
	oldStatus := chi.URLParam(r, "oldStatus")
	okay := true

	// get host service
	hs, err := repo.DB.GetHostServiceByID(hostServiceID)
	if err != nil {
		log.Println(err)
		okay = false
	}

	// get host
	h, err := repo.DB.GetHostByID(hs.HostID)
	if err != nil {
		log.Println(err)
		okay = false
	}

	// test the service
	newStatus, msg := repo.testServiceForHost(h, hs)

	// update the host service in the database with status (if changed) and last check
	hs.Status = newStatus
	hs.LastCheck = time.Now()
	hs.UpdatedAt = time.Now()


	err = repo.DB.UpdateHostService(hs)
	if err != nil {
		log.Println(err)
		okay = false
	}

	// broadcast service status changed event -using websocket

	// create json
	var resp jsonResp
	if okay {
		resp = jsonResp{
			OK:            true,
			Message:       msg,
			ServiceID:     hs.ServiceID,
			HostServiceID: hs.ID,
			HostID:        hs.HostID,
			OldStatus:     oldStatus,
			NewStatus:     newStatus,
			LastCheck:     time.Now(),
		}
	} else {
		resp.OK = false
		resp.Message = "Something went wrong"
	}

	// send json to client
	out, _ := json.MarshalIndent(resp, "", "    ")
	w.Header().Set("Content-Type", "application/json")
	w.Write(out)
}

// testServiceForHost tests a service for a host
func (repo *DBRepo) testServiceForHost(h models.Host, hs models.HostService) (string, string){
	var msg, newStatus string

	switch hs.ServiceID {
	case HTTP:
		newStatus, msg = testHTTPForHost(h.URL)
		break
	}

	return newStatus, msg
}

// testHTTPForHost tests HTTP service
func testHTTPForHost(url string) (string, string) {
	// trim the suffix of url if it's '/'
	if strings.HasSuffix(url, "/") {
		url = strings.TrimSuffix(url, "/")
	}

	// we only test 'http://' instead of 'https://'
	// n = -1: we want to replace it everywhere of url
	url = strings.Replace(url, "https://", "http://", -1)

	// send the request to test
	resp, err := http.Get(url)
	if err != nil {
		return "problem", fmt.Sprintf("%s - %s", url, "error connecting")
	}
	defer resp.Body.Close()

	// if status code is not 200, something went wrong
	if resp.StatusCode != http.StatusOK {
		return "problem", fmt.Sprintf("%s - %s", url, resp.Status)
	}

	return "healthy", fmt.Sprintf("%s - %s", url, resp.Status)
}


