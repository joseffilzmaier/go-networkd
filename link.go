package networkd

import "github.com/godbus/dbus/v5"

var (
	linkPath = "org.freedesktop.network1.Link"
)

type Link struct {
	connObj dbus.BusObject
}

func (l *Link) Reconfigure() error {
	return l.connObj.Call(linkPath+".Reconfigure", 0).Err
}

func (l *Link) Renew() error {
	return l.connObj.Call(linkPath+".Renew", 0).Err
}

func (l *Link) RevertDNS() error {
	return l.connObj.Call(linkPath+".RevertDNS", 0).Err
}

func (l *Link) RevertNTP() error {
	return l.connObj.Call(linkPath+".RevertNTP", 0).Err
}

type LinkResponse struct {
	Index  int
	Name   string
	Object dbus.ObjectPath
}
