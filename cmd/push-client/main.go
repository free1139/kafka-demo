package main

import (
	"fmt"

	"kafka-demo/module/push"

	"github.com/gwaylib/rpctry"
)

func main() {
	c := rpctry.NewClient("127.0.0.1:5001")

	in := &push.JobArg{}
	ret := &push.JobRet{}
	if err := c.TryCall(push.PushSvcName+".CreateJob", in, ret); err != nil {
		panic(err)
	}
	fmt.Println(*ret)
}
