package main

import (
	"flag"

	"github.com/mgajewskik/x/internal"
)

var flgCleanDownloads bool

func init() {
	flag.BoolVar(&flgCleanDownloads, "cleanDownloads", false, "Clean downloads folder")
	flag.Parse()
}

func main() {
	if flgCleanDownloads {
		internal.CleanDownloads()
		return
	}

	flag.Usage()
}
