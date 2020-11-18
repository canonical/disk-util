package service

import (
	"github.com/godbus/dbus"
	"path"
)

const (
	udisksBlockFormat = "org.freedesktop.DBus.UDisks2.Block.Format"
)

// FormatPartition use udisks to format a partition
func (db *DBus) FormatPartition(device, format string) error {
	devicePath := path.Join("/org/freedesktop/UDisks2/block_devices", device)
	obj := db.getBusObject("org.freedesktop.UDisks2", devicePath)

	call := obj.Call(udisksBlockFormat, 0, format, map[string]dbus.Variant{}, false)

	return call.Err
}