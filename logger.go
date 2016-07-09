package main

import (
	"os"
	"time"

	"github.com/op/go-logging"
)

var log = logging.MustGetLogger("nut")

func init() {
	lfd, err := os.OpenFile(
		time.Now().Format("2006-01-02.log"),
		os.O_WRONLY|os.O_APPEND|os.O_CREATE,
		0600,
	)
	if err != nil {
		return
	}
	bkd1 := logging.AddModuleLevel(
		logging.NewBackendFormatter(
			logging.NewLogBackend(os.Stderr, "", 0),
			logging.MustStringFormatter(`%{color}%{time:2006-01-02 15:04:05.000} ▶ %{level:.4s} %{color:reset} %{message}`)),
	)

	bkd2 := logging.AddModuleLevel(
		logging.NewBackendFormatter(
			logging.NewLogBackend(lfd, "", 0),
			logging.MustStringFormatter(`%{time:15:04:05.000} %{level:.4s} %{message}`)),
	)

	logging.SetBackend(
		bkd1,
		bkd2,
	)
}
