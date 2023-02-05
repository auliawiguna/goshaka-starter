package jobs

import (
	"fmt"
	"time"

	"github.com/go-co-op/gocron"
)

var sampleJob = gocron.NewScheduler(time.UTC)

func RunCron() {
	fmt.Println("Start jobs...")
	sampleJob.Every(1).Day().At("08:00").Do(func() {
		fmt.Println("Get up! Get to work!")
	})

	sampleJob.Every(10).Minutes().Do(func() {
		fmt.Println("Run every 10 minutes")
	})

	sampleJob.StartAsync()
}
