package main

import (
	_ "errors"
	"flag"
	_ "fmt"
	q "gmq/queue"
	"os"
)

var configfile string

func main() {

	flag.StringVar(&configfile, "f", "gmq.conf", "Location of configuation file")
	flag.Parse()

	//	if err := parseConfigFile(configfile); err != nil {
	//		panic(errors.New("Unable to parse configuration file %s"))
	//	}
}

func parseConfigFile(f string) error {
	//TODO
	_, err := os.Open(f)
	return err
}
