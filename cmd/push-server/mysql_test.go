package main

import (
	"strconv"
	"testing"
	"time"

	"github.com/gwaylib/qsql"
)

func TestMysqlDial(t *testing.T) {
	db, err := NewMysql("root:123456@tcp(127.0.0.1:3307)/push?timeout=30s&loc=Local&parseTime=true&allowOldPasswords=1")
	if err != nil {
		t.Fatal(err)
	}
	if _, err := db.Exec(`
CREATE TABLE IF NOT EXISTS testing 
(
    username VARCHAR(32) NOT NULL,
    passwd VARCHAR(128) NOT NULL,
    created_at DATETIME NOT NULL DEFAULT NOW(),
    PRIMARY KEY(username)
) ENGINE=InnoDB;
	`); err != nil {
		t.Fatal(err)
	}

	username := strconv.FormatInt(time.Now().UnixNano(), 16)
	passwd := username
	if _, err := db.Exec("INSERT INTO testing (username,passwd) VALUES (?,?);", username, passwd); err != nil {
		t.Fatal(err)
	}
	outputPwd := ""
	if err := qsql.QueryElem(db, &outputPwd, "SELECT passwd FROM testing WHERE username=?", username); err != nil {
		t.Fatal(err)
	}
	if outputPwd != passwd {
		t.Fatalf("expect: %s, but:%s", passwd, outputPwd)
	}
}
