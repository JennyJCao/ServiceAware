package handlers

import (
	"fmt"
	"github.com/tsawler/vigilate/internal/helpers"
	"github.com/tsawler/vigilate/internal/models"
	"log"
	"net/http"
)

// ListEntries lists schedule entries
func (repo *DBRepo) ListEntries(w http.ResponseWriter, r *http.Request) {
	var items []models.Schedule

	for k, v := range repo.App.MonitorMap {
		var item models.Schedule
		item.ID = k
		item.EntryID = v
		item.Entry = app.Scheduler.Entry(v)
		hs, err := repo.DB.GetHostServiceByID(k)
		if err != nil {
			log.Println(err)
			return
		}
		item.ScheduleText = fmt.Sprintf("@every %d%s", hs.ScheduleNumber, hs.ScheduleUnit)
		item.LastRunFromHS = hs.LastCheck
		item.Host = hs.HostName
		item.Service = hs.Service.ServiceName
		items = append(items, item)
	}

	err := helpers.RenderPage(w, r, "schedule", nil, nil)
	if err != nil {
		printTemplateError(w, err)
	}
}
