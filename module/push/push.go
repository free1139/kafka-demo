package push

import "time"

const PushSvcName = "push_svc"

type JobArg struct {
	CronTime time.Time
	Records  [][]string
}
type JobRet struct {
	Time time.Time
}

type PushSvc interface {
	CreateJob(arg *JobArg, ret *JobRet) error
}
