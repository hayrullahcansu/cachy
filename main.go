package main

import (
	"flag"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/hayrullahcansu/cachy/api"
	"github.com/hayrullahcansu/cachy/data/constants"
	"github.com/hayrullahcansu/cachy/framework/logging"
)

var ops int64

func main() {
	name := flag.String("name", constants.SoftwareName, "name to print")
	flag.Parse()
	logging.Infof("Starting service for %s", *name)
	// setup signal catching
	sigs := make(chan os.Signal, 1)
	// catch all signals since not explicitly listing
	signal.Notify(sigs)
	signal.Notify(sigs, syscall.SIGQUIT)
	// method invoked upon seeing signal

	go func() {
		s := <-sigs
		logging.Infof("RECEIVED SIGNAL: %s", s)
		AppCleanup()
		os.Exit(1)
	}()

	apiWorker := api.NewApiWorker()
	apiWorker.ListenAndServce()

	// man.InitSystem()

}
func AppCleanup() {
	time.Sleep(time.Millisecond * time.Duration(1000))
	logging.Infof("CLEANUP APP BEFORE EXIT!!!")
}
