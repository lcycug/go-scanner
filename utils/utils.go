package utils

import (
	"log"

	"github.com/lcycug/go-scanner/models"
)

func LogFatal(e models.Errors) {
	switch e.GetType() {
	case models.ERROR:
		if e.GetError() != nil {
			log.Fatal(e.GetMessage())
		}
	case models.OK:
		if !e.GetOk() {
			log.Fatal(e.GetMessage())
		}
	}
}
