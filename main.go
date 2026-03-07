package main

import (
	"fmt"
	"os"
	"path/filepath"

	flag "github.com/spf13/pflag"
)

// binary metadata
var Version = "NOTSET"

// flags
var (
	PrintVersion = flag.BoolP("version", "V", false, "print version and exit")
	ConfigPath   = flag.StringP("config", "c", "./cyborg.toml", "path to configuration")
	ShowHelp     = flag.BoolP("help", "h", false, "show help")
)

func init() {
	flag.Usage = func() {
		name := filepath.Base(os.Args[0])
		fmt.Fprintf(os.Stderr, "%v - CY_BORG game interface\n\n", name)
		fmt.Fprintf(os.Stderr, "Usage:\n  %v [options]\n\n", name)
		fmt.Fprintf(os.Stderr, "Flags:\n")
		flag.PrintDefaults()
	}
}

func main() {
	flag.Parse()

	if *PrintVersion {
		fmt.Println(Version)
		return
	}
	if *ShowHelp {
		flag.Usage()
		os.Exit(0)
	}
}
