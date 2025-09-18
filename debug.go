package util

import (
	"fmt"
	"log"
	"os"
)

var _, IsDebug = os.LookupEnv("DEBUG")

func SetLogTextOnly() {
	log.SetFlags(log.Flags() &^ (log.Ldate | log.Ltime))
}

func Debug(s string) {
	if IsDebug {
		fmt.Println(s)
	}
}
