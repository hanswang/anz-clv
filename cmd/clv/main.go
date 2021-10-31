package main

import (
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/hanswang/clv/internal/processor"
	log "github.com/sirupsen/logrus"
)

func main() {
	var filename string
	flag.StringVar(&filename, "f", "", "Path to a CSV file as input read for credit limit validator, e.g ./sample.csv")
	var verbose bool
	flag.BoolVar(&verbose, "v", false, "Enable verbose logging")

	flag.Parse()

	if filename == "" {
		die("clv: flag -f for input filename is required, abort\nRun 'clv -h' for usage.")
	}
	
	if verbose {
		log.SetLevel(log.DebugLevel)
	}

	exist, err := Exists(filename)
	if exist {
		csvBody, err := ioutil.ReadFile(filename)
		if err != nil {
			die("clv: read file %v with error: %w, abort", filename, err)
		}
		processor.Run(csvBody)
	} else if err != nil {
		die("clv: check file %v with error: %w, abort", filename, err)
	} else {
		die("clv: file %v doesnot exist, abort", filename)
	}

}

func die(format string, args ...interface{}) {
	if _, err := fmt.Fprintf(os.Stderr, format+"\n", args...); err != nil {
		os.Exit(1)
	}
	os.Exit(0)
}

func Exists(name string) (bool, error) {
    _, err := os.Stat(name)
    if err == nil {
        return true, nil
    }
    if errors.Is(err, os.ErrNotExist) {
        return false, nil
    }
    return false, err
}