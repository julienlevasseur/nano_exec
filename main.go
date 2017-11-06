package main

import (
	"os"
	"fmt"
	"log"
	"os/exec"
	"io/ioutil"
	"gopkg.in/yaml.v2"
)

type Job struct {
	Name 		string 		`json:"name"`
	Command		string 		`json:"command"`
	Arguments	[]string 	`json:arguments`
}

func main() {
	job := Job{}

	filename, f_err := ioutil.ReadFile(string(os.Args[1]))
	if f_err != nil {
		log.Fatal("error: %v", f_err)
	}

	err := yaml.Unmarshal(filename, &job)
	if err != nil {
		log.Fatal("error: %v", err)
	}
	fmt.Printf("job: \n%v\n\n", job)

	cmd := job.Command
	args := job.Arguments
	var out []byte

	if out, err = exec.Command(cmd, args...).Output(); err != nil {
		fmt.Println(os.Stderr, err)
		os.Exit(1)
	}

	fmt.Printf(string(out))
}
