package cronjob

import (
	"database/sql"
	"log"

	"github.com/robfig/cron"
)


type CronJob struct {
	db *sql.DB
}

func New(db *sql.DB) *CronJob {
	return &CronJob{
		db: db,
	}
}

func (c *CronJob) Start() {
	cron := cron.New()

	err := cron.AddFunc("@every 10m", func() {
		err := DeleteExpiredOtp(c.db)
		if err != nil {
			log.Println(err)
		}
	})

	if err != nil {
		log.Println(err)
	}

	cron.Start()

	// c.AddFunc("0 10 * * * *", func() {
	// })
}
