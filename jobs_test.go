package main

import (
	"reflect"
	"testing"
)

func TestGetJobs(t *testing.T) {
	joblist := GetJobs()
	if reflect.TypeOf(joblist).Kind() == reflect.Slice {
		PrintGreen("[OK] GetJobs return type is: []string\n")
	} else {
		t.Fatalf("[FAIL] GetJobs return type is: %v, instead of: []string", reflect.TypeOf(joblist))
	}
}
