package service

import (
	"github.com/godbus/dbus"
	"log"
	"path"
)

const (
	udisksPartitionDelete       = "org.freedesktop.UDisks2.Partition.Delete"
	udisksPartitionCreate       = "org.freedesktop.UDisks2.PartitionTable.CreatePartition"
	udisksPartitionCreateFormat = "org.freedesktop.UDisks2.PartitionTable.CreatePartitionAndFormat"
)

// CreatePartition creates a new partition on a block device
func (db *DBus) CreatePartition(device, partType, name string, offset, size uint64) error {
	devicePath := path.Join("/org/freedesktop/UDisks2/block_devices", device)
	obj := db.getBusObject("org.freedesktop.UDisks2", devicePath)

	var resp string
	err := obj.Call(udisksPartitionCreate, 0, offset, size, partType, name, map[string]dbus.Variant{}).Store(&resp)
	log.Println(resp)
	return err
}

// DeletePartition deletes an existing partition
func (db *DBus) DeletePartition(device string) error {
	devicePath := path.Join("/org/freedesktop/UDisks2/block_devices", device)
	obj := db.getBusObject("org.freedesktop.UDisks2", devicePath)

	call := obj.Call(udisksPartitionDelete, 0, map[string]dbus.Variant{})
	return call.Err
}

// CreateAndFormatPartition creates a new partition on a block device and formats it
func (db *DBus) CreateAndFormatPartition(device, partType, name, formatType string, offset, size uint64) error {
	devicePath := path.Join("/org/freedesktop/UDisks2/block_devices", device)
	obj := db.getBusObject("org.freedesktop.UDisks2", devicePath)

	var resp string
	err := obj.Call(udisksPartitionCreateFormat, 0, offset, size, partType, name, map[string]dbus.Variant{}, formatType, map[string]dbus.Variant{}).Store(&resp)
	log.Println(resp)
	return err
}
