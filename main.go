package main

import (
	_ "github.com/breathbath/erplyapi/log"
	"github.com/breathbath/erplyapi/cli"
	log "github.com/sirupsen/logrus"
)

func main() {
	err := cli.Execute()
	if err != nil {
		log.Fatal(err)
	}
}
