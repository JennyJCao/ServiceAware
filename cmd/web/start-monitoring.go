package main

import (
	"fmt"
	"log"
)

// job is the unit of work to be performed
type job struct {
	HostServiceID int
}

// Run runs the scheduler job
func (j job) Run() {

	repo.ScheduledCheck(j.HostServiceID)
}

// startMonitoring starts the monitoring process
func startMonitoring() {
	if preferenceMap["monitoring_live"] == "1" {
		// trigger a message to broadcast to all clients that app is starting to monitor
		data := make(map[string]string)
		data["message"] = "Monitoring is starting..."
		// send this message to all users
		err := app.WsClient.Trigger("public-channel", "app-starting", data)
		if err != nil {
			log.Println(err)
		}

		// get all of the services that we want to monitor
		servicesToMonitor, err := repo.DB.GetServicesToMonitor()
		if err != nil {
			log.Println(err)
		}

		// range through the services
		for _, x := range servicesToMonitor {
			log.Println("*** Service to monitor on", x.HostName, "is", x.Service.ServiceName)

			// get the schedule unit and number
			var sch string // scheduler
			// "@every 3m": is the cron expression
			if x.ScheduleUnit == "d" {
				sch = fmt.Sprintf("@every %d%s", x.ScheduleNumber*24, "h")
			} else {
				sch = fmt.Sprintf("@every %d%s", x.ScheduleNumber, x.ScheduleUnit)
			}

			// create a job
			var j job
			j.HostServiceID = x.ID
			scheduleID, err := app.Scheduler.AddJob(sch, j)
			if err != nil {
				log.Println(err)
			}

			// save the id of the job so we can start/stop it

			// broadcast over websockets the fact that the service is scheduled

		}
	}
}