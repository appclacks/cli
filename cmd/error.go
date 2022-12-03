package cmd

import (
	"fmt"
	"os"
)

func exitIfError(err error) {
	if err != nil {
		fmt.Printf(err.Error())
		os.Exit(1)
	}
}
