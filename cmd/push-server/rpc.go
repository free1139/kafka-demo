package main

import (
	"net"
	"net/rpc"
	"time"

	"kafka-demo/module/push"

	"github.com/gwaylib/errors"
	"github.com/gwaylib/log"
	"github.com/gwaylib/qsql"
)

type pushSvc struct {
	db *qsql.DB
	kf *Kafka
}

func NewService() (push.PushSvc, error) {
	db, err := NewMysql("root:123456@tcp(127.0.0.1:3307)/push?timeout=30s&loc=Local&parseTime=true&allowOldPasswords=1")
	if err != nil {
		return nil, errors.As(err)
	}
	kf := NewKafka("localhost:9092")
	svc := &pushSvc{
		db: db,
		kf: kf,
	}
	if err := rpc.RegisterName(push.PushSvcName, svc); err != nil {
		return nil, errors.As(err)
	}
	return svc, nil
}

func (svc *pushSvc) close() error {
	// TODO: close resource
	return nil
}

func (svc *pushSvc) listen(addr string) {
	log.Infof("Rpc listen: %s", addr)
	conn, err := net.Listen("tcp", addr)
	if err != nil {
		log.Exit(2, errors.As(err))
		return
	}
	defer func() {
		qsql.Close(conn)
		log.Exit(0, "rpc conn exit")
	}()
	rpc.Accept(conn)
}

func (svc *pushSvc) CreateJob(arg *push.JobArg, ret *push.JobRet) error {
	if len(arg.Records) == 0 {
		return errors.New("no records")
	}
	ret.Time = time.Now()

	// TODO: deal data

	return nil
}
