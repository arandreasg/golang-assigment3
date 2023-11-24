package main

import (
	"assigmenttiga/config"
	"assigmenttiga/controller"

	"github.com/jasonlvhit/gocron"
	_ "github.com/lib/pq"
)

func main() {
	config.Connect()
	gocron.Every(15).Seconds().Do(controller.UpdateData)

	// Start scheduler
	<-gocron.Start()

	select {}
}
