package push

const PushSvcName = "push_svc"

type JobArg struct {
}
type JobRet struct {
}

type PushSvc interface {
	CreateJob(arg *JobArg, ret *JobRet) error
}
