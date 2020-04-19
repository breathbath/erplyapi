package main

import (
	"github.com/breathbath/erplyapi/server"
	"github.com/breathbath/go_utils/utils/env"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	logLvl, err := log.ParseLevel(env.ReadEnv("LOG_LEVEL", "info"))
	if err != nil {
		log.Fatal(err)
	}
	log.SetLevel(logLvl)

	err = server.Start()
	if err != nil {
		log.Fatal(err)
	}
}
