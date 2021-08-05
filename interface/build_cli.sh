#!/bin/bash
go build -o cbadm ./cli/cbadm/cbadm.go
go build -ldflags="-X 'github.com/cloud-barista/poc-cicd-spider/interface/cli/spider/cmd.Version=v0.4.0' -X 'github.com/cloud-barista/poc-cicd-spider/interface/cli/spider/cmd.CommitSHA=$(git rev-parse --short HEAD)' -X 'github.com/cloud-barista/poc-cicd-spider/interface/cli/spider/cmd.User=$(id -u -n)' -X 'github.com/cloud-barista/poc-cicd-spider/interface/cli/spider/cmd.Time=$(date)'" -o spctl ./cli/spider/spider.go 
