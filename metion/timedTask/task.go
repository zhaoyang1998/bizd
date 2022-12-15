package timedTask

import (
	"bizd/metion/service"
	"github.com/go-co-op/gocron"
	"time"
)

func StartTimedTask() {
	timezone := time.FixedZone("CST", 8*3600) //替换上海时区

	s := gocron.NewScheduler(timezone)

	_, err := s.Every(1).Monday().At("17:00").Do(func() {
		go service.GetCurrentConductor()
	})
	if err != nil {
		return
	}
	s.StartBlocking()
}
