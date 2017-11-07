package main

import (
	"os"
	"fmt"
	"reflect"
	"strings"
	"testing"
	"io/ioutil"
	"github.com/fatih/color"
)

func TestGetJobs(t *testing.T) {
	// create a fake job :
	fcontent := []byte("---\ncommand: echo\narguments:\n  - 'uF7yVCdZ'")
	err := ioutil.WriteFile("TestGetJobs.yml", fcontent, 0644)
	if err != nil {
		t.Fatalf("[ERROR] Unable to write file: TestGetJobs.yml")
	}

	// Ensure GetJobs return type is correct :
	joblist := GetJobs()
	if reflect.TypeOf(joblist).Kind() == reflect.Slice {
		fmt.Printf("%s\n", color.GreenString("[OK] GetJobs return type is: []string"))
	} else {
		t.Fatalf("[FAIL] GetJobs return type is: %v, instead of: []string", reflect.TypeOf(joblist))
	}

	// Check the previously created job is found in the return value :
	telltale := 0
	for _, job := range joblist {
		if strings.Contains(job, "TestGetJobs") {
			telltale = 1
		}
	}

	if telltale == 1 {
		fmt.Printf("%s\n", color.GreenString("[OK] GetJobs returned TestGetJobs"))
	} else {
		t.Fatalf("[FAIL] GetJobs doesn't return TestGetJobs")
	}

	// Remove the fake job file :
	err = os.Remove("./TestGetJobs.yml")
	if err != nil {
		fmt.Printf("%s\n", color.YellowString("[Warning] Unable to write file: TestGetJobs.yml"))
	}
}

/*func TestExecuteJob(t *testing.T) {

}*/
