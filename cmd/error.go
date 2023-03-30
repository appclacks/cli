package cmd

import (
	"fmt"
	"os"
)

func exitIfError(err error) {
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
