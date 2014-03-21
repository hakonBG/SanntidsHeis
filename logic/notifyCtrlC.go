package logic

import (
	"./../driver"
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func cleanup() {
	fmt.Println("   CTRL C PUSHED")
	driver.Elev_set_speed(0)
}

func NotifyCtrlC() {

	c := make(chan os.Signal, 1)

	signal.Notify(c, os.Interrupt)
	signal.Notify(c, syscall.SIGTERM)
	<-c
	cleanup()
	os.Exit(1)
}
