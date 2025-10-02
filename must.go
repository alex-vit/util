package util

import (
	"fmt"
	"os"
)

func Must[V any](v V, err error) V {
	if err != nil {
		panic(err)
	}
	return v
}

func OrExit[T any](t T, err error) T {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return t
}
