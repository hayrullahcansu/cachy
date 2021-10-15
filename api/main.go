package main

import (
	"flag"
	"time"

	"github.com/hayrullahcansu/cachy/api/listener"
	"github.com/hayrullahcansu/cachy/data/constants"
	"github.com/hayrullahcansu/cachy/framework/logging"
)

var ops int64

func main() {
	name := flag.String("name", constants.SoftwareName, "name to print")
	flag.Parse()
	logging.Infof("Starting service for %s", *name)

	apiListener := listener.NewApiListener()
	apiListener.ListenAndServe()

	// man.InitSystem()

}
func AppCleanup() {
	time.Sleep(time.Millisecond * time.Duration(1000))
	logging.Infof("CLEANUP APP BEFORE EXIT!!!")
}
