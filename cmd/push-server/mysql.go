package main

import (
	_ "github.com/go-sql-driver/mysql"

	"github.com/gwaylib/qsql"
)

func init() {
	db, err := qsql.Open(qsql.DRV_NAME_MYSQL, "root:123456@tcp(127.0.0.1:3307)/push?timeout=30s&loc=Local&parseTime=true&allowOldPasswords=1")
	if err != nil {
		panic(err)
	}
	qsql.RegCache("main", db)
}

func NewMysql() *qsql.DB {
	return qsql.GetCache("main")
}
