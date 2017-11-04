package main
  
import (  
  "os"
  "fmt"
  "log"
  "bytes"
  "strings"
  "os/exec"
  "io/ioutil"
  "gopkg.in/yaml.v2"
)

type Job struct {
  Name  string `json:"name"`
  Action  string `json:"action"`
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

  //fmt.Printf("action: %v\n", job.Action)
  cmd_elem := strings.Split(job.Action, " ")
  var str bytes.Buffer
  for _, i := range cmd_elem {
    str.WriteString(i)
    str.WriteString(" ")
  }
  fmt.Printf("action: %v\n", str.String())


  cmd := exec.Command(str.String())
  var out bytes.Buffer
  cmd.Stdout = &out
  c_err := cmd.Run()
  if err != nil {
    log.Fatal(c_err)
  }

  fmt.Printf("%q\n", out.String())
}
