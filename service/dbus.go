package service

import (
	"github.com/godbus/dbus"
	"log"
)

// DBusService is the interface to abstract dbus
type DBusService interface {
	FormatPartition(device, format string) error
}

// DBus implements a wrapper for the dbus service
type DBus struct {
	systemBus *dbus.Conn
}

// NewDBus creates a dbus wrapper service
func NewDBus() (*DBus, error) {
	bus, err := dbus.SystemBus()
	if err != nil {
		log.Printf("Failed to access system dbus: %v", err)
		return nil, err
	}
	return &DBus{bus}, nil
}

func (db *DBus) getBusObject(dest, path string) dbus.BusObject {
	return busObject(db.systemBus, dest, path)
}

// busObject returns a dbus bus object interface (mockable for tests)
var busObject = func(systemBus interface{}, dest, path string) dbus.BusObject {
	return systemBus.(*dbus.Conn).Object(dest, dbus.ObjectPath(path))
}
