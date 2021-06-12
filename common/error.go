package common

import (
	"log"
)

func CheckError(message string, err error) {
	if err != nil {
		log.Fatal(message, err)
	}
}

func LogFatal(message string, err error) {
	log.Fatal(message, err)
}
