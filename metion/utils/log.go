package utils

import (
	"log"
	"os"
)

func SetupLogger() {
	logFileLocation, _ := os.OpenFile("./logInfo.log", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0744)
	log.SetOutput(logFileLocation)
}
