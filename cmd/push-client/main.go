package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"time"

	"kafka-demo/module/push"

	"github.com/gwaylib/errors"
	"github.com/gwaylib/rpctry"
)

var (
	cronTime = flag.String("c", "", "send now when cron time not set, the time format is RFC3339")
)

func main() {
	// read file
	flag.Parse()
	args := flag.Args()
	if len(args) == 0 {
		panic("Need input file path")
	}
	fileData, err := ioutil.ReadFile(args[0])
	if err != nil {
		panic(errors.As(err))
	}
	records, err := ParseCsv(fileData)
	if err != nil {
		panic(errors.As(err))
	}

	cron, err := time.Parse(*cronTime, time.RFC3339)
	if err != nil {
		if len(*cronTime) != 0 {
			panic(errors.As(err, *cronTime))
		}
		cron = time.Now()
	}

	c := rpctry.NewClient("127.0.0.1:5001")

	in := &push.JobArg{
		CronTime: cron,
		Records:  records,
	}
	ret := &push.JobRet{}
	if err := c.TryCall(push.PushSvcName+".CreateJob", in, ret); err != nil {
		panic(err)
	}
	fmt.Println(ret.Time)
}
