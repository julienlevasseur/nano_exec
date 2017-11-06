# nano exec
![Build Status](https://travis-ci.org/julienlevasseur/nano_exec.svg?branch=master)

Nano Exec is a nano service which present an api allowing :

* Create jobs
* List jobs
* Execute Jobs

## Job

A job is a yaml file with this structure :

```yaml
---
command: systemctl
arguments:
  - status
  - sshd

```

Where :

* `command` is the command name (eg: ls or systemctl)
* `arguments` is list of command's arguments

## Routes
