package util

import (
	"log"
	"os"
)

var _, IsDebug = os.LookupEnv("DEBUG")

func Debug(s string) {
	if IsDebug {
		log.Print(s)
	}
}
