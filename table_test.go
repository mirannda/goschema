package main

import (
	"testing"
)

func TestGetTables(t *testing.T) {
	opt := options{
		user:     "admin",
		password: "abc",
		host:     "localhost",
		port:     3306,
		name:     "redivus",
	}
	db, err := dbConnect(&opt)
	if err != nil {
		t.Fatal(err)
	}
	tbls, err := getTables(db, opt.name)
	if err != nil {
		t.Fatal(err)
	}
	if len(tbls) < 10 {
		t.Errorf("Expected > 10 tables got %v\n", len(tbls))
	}
}
