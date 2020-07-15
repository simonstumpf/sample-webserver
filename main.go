package main

import (
	"time"

	"github.com/sirupsen/logrus"
)

func main() {
	for {
		logrus.Info("1")
		time.Sleep(time.Second)
	}

}
