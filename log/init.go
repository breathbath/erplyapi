package log

import (
	"github.com/breathbath/go_utils/utils/env"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	logLvl, err := log.ParseLevel(env.ReadEnv("LOG_LEVEL", "info"))
	if err != nil {
		log.Fatal(err)
	}
	log.SetLevel(logLvl)

	customFormatter := &log.TextFormatter{
		FullTimestamp: true,
	}
	customFormatter.TimestampFormat = "2006-01-02 15:04:05.000000"
	log.SetFormatter(customFormatter)

}
