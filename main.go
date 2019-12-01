package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/golang/glog"
	notify "github.com/verchol/notify/pkg"
)

func usage() {
	fmt.Fprintf(os.Stderr, "usage: example -stderrthreshold=[INFO|WARN|FATAL] -log_dir=[string]\n")
	flag.PrintDefaults()
	os.Exit(2)
}

func main() {

	msg := flag.String("msg", "Hello, from slackjob", "message that will be posted to slack")
	flag.Parse()
	var channel string
	channel = os.Getenv("SLACK_CHANNEL")
	if channel == "" {
		glog.Error("no slack channel was provided")
		os.Exit(1)
	}

	glog.V(2).Info("channel is %s", channel)
	glog.V(2).Info("test")
	date := time.Now().Format("2 Jan 2006 15:04:05")
	msgToSend := fmt.Sprintf("[%s]%s", date, *msg)

	_, err := notify.SendHttpMessage(channel, msgToSend)
	if err != nil {
		panic(err)
	}
}
