package utils

import (
	log "github.com/sirupsen/logrus"
	"io"
)

//Close closes and reports as warning a possible error
func Close(closer io.Closer) {
	err := closer.Close()
	if err != nil {
		log.Warnf("failed to close: %v", err)
	}
}
