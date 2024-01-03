package main

import (
	"fmt"
	"os"

	"github.com/go-errors/errors"
)

func Fatal(err error) {
	if err == nil {
		return
	}

	// if errors.Is(err, errors.Error) {
	// 	fmt.Println(err.(*errors.Error).ErrorStack())
	// } else {
	// 	log.Fatal(err)
	// }

	fmt.Println(err.(*errors.Error).ErrorStack())
	os.Exit(1)
}
