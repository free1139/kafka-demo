package main

import (
	"encoding/csv"
	"strings"
)

func ParseCsv(data []byte) ([][]string, error) {
	r := csv.NewReader(strings.NewReader(string(data)))
	r.Comma = ';'
	r.Comment = '#'
	return r.ReadAll()
}
