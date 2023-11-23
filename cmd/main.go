package main

import (
	"log"
	"time"

	"github.com/go-co-op/gocron"
	"github.com/yannickkirschen/prepper"
)

const PREPPER_D_PATH = "/etc/prepper/prepper.d"

func main() {
	download, err := prepper.ReadDownloadFile(PREPPER_D_PATH + "/" + "01-download.json")
	if err != nil {
		panic(err)
	}

	scheduler := gocron.NewScheduler(time.UTC)
	scheduler.Cron(download.Cron).Do(func() {
		log.Println("Downloading files")
		err = download.Download()
		if err != nil {
			panic(err)
		}
	})

	log.Println("Scheduled download of files")
	scheduler.StartBlocking()
}
