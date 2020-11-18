package main

import (
	"flag"
	"fmt"
	"github.com/slimjim777/disk-util/service"
	"log"
)

var (
	command    string
	device     string
	formatType string
)

func main() {
	dBus, err := service.NewDBus()
	if err != nil {
		log.Fatal(err)
	}

	parseArgs()

	switch command {
	case "format":
		err = dBus.FormatPartition(device, formatType)
	default:
		err = fmt.Errorf("unsupported command: %s", command)
	}

	if err != nil {
		log.Fatal(err)
	}
}

func parseArgs() {
	flag.StringVar(&command, "command", "list", "The udisks command e.g. list, format")
	flag.StringVar(&device, "device", "", "The device to operate on e.g. sdd, sdf3")
	flag.StringVar(&formatType, "type", "", "The format type e.g. ext4")

	flag.Parse()
}
