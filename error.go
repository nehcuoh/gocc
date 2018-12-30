package main

import (
	"fmt"
	"os"
)

var errorCount = 0
var warnCount = 0

func Error(coord *Coord, format string, v ...interface{}) {
	errorCount++
	if coord != nil {
		fmt.Fprintf(os.Stderr, "(%s,%d)", coord.filename, coord.ppline)
	}
	fmt.Fprint(os.Stderr, "error:")
	fmt.Fprintf(os.Stderr, format, v...)
	fmt.Fprintln(os.Stderr, "")
}
func Warning(coord *Coord, format string, v ...interface{}) {
	warnCount++
	if coord != nil {
		fmt.Fprintf(os.Stderr, "(%s,%d)", coord.filename, coord.ppline)
	}
	fmt.Fprint(os.Stderr, "warning:")
	fmt.Fprintf(os.Stderr, format, v...)
	fmt.Fprintln(os.Stderr, "")
}

func Fatal(format string, v ...interface{}) {
	fmt.Fprintf(os.Stderr, "fatal:")
	fmt.Fprintf(os.Stderr, format, v...)
	os.Exit(1)
}
