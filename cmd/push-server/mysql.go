package main

import (
	_ "github.com/go-sql-driver/mysql"

	"github.com/gwaylib/errors"
	"github.com/gwaylib/qsql"
)

func NewMysql(dns string) (*qsql.DB, error) {
	db, err := qsql.Open(qsql.DRV_NAME_MYSQL, dns)
	return db, errors.As(err)
}
