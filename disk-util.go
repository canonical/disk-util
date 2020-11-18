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
	name       string
	offset     uint64
	size       uint64
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
	case "create":
		err = dBus.CreatePartition(device, formatType, name, offset, size)
	default:
		err = fmt.Errorf("unsupported command: %s", command)
	}

	if err != nil {
		log.Fatal(err)
	}
}

func parseArgs() {
	flag.StringVar(&command, "command", "list", "The udisks command e.g. list, create, format")
	flag.StringVar(&device, "device", "", "The device to operate on e.g. sdd, sdf3")
	flag.StringVar(&formatType, "type", "", "The format type e.g. ext4")
	flag.StringVar(&name, "name", "", "Name of a partition to be created")
	flag.Uint64Var(&offset, "offset", 0, "Offset to create a partition (Mb)")
	flag.Uint64Var(&size, "size", 0, "Size of a partition (Mb)")

	flag.Parse()
}
