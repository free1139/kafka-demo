package main

import (
	"flag"
	"fmt"
	"io/ioutil"

	"kafka-demo/module/push"

	"github.com/gwaylib/errors"
	"github.com/gwaylib/rpctry"
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
	fmt.Println(string(fileData))

	c := rpctry.NewClient("127.0.0.1:5001")
	in := &push.JobArg{}
	ret := &push.JobRet{}
	if err := c.TryCall(push.PushSvcName+".CreateJob", in, ret); err != nil {
		panic(err)
	}
	fmt.Println(*ret)
}
