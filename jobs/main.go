package jobs

import (
	"fmt"
	"time"

	"github.com/go-co-op/gocron"
)

var sampleJob = gocron.NewScheduler(time.UTC)

func RunCron() {
	fmt.Println("Start jobs...")
	_, err := sampleJob.Every(1).Day().At("08:00").Do(func() {
		fmt.Println("Get up! Get to work!")
	})
	if err != nil {
		fmt.Println("error cron")
	}

	_, err2 := sampleJob.Every(10).Minutes().Do(func() {
		fmt.Println("Run every 10 minutes")
	})
	if err2 != nil {
		fmt.Println("error cron")
	}

	sampleJob.StartAsync()
}
