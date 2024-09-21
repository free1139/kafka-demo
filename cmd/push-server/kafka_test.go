package main

import "testing"

func TestKafkaDial(t *testing.T) {
	k := NewKafka()
	conn, err := k.dial("testing")
	if err != nil {
		t.Fatal(err)
	}
	if err := conn.Close(); err != nil {
		t.Fatal(err)
	}
}
