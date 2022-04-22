package errcheck

import (
	"fmt"
	"os"
)

// CheckAndPanic check error and panic if error is not nil
func CheckAndPanic(err error) {
	if err != nil {
		panic(err)
	}
}

// CheckAndExit check error and exit if error is not nil
func CheckAndExit(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

// CheckAndPrintln check error and print if error is not nil
func CheckAndPrintln(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

// CheckAndPrintf check error and print if error is not nil
func CheckAndPrintf(err error, format string, args ...interface{}) {
	if err != nil {
		fmt.Printf(format, args...)
	}
}
