package networkd

import (
	"errors"
	"io"

	"github.com/godbus/dbus/v5"
	"github.com/godbus/dbus/v5/introspect"
)

var (
	managerPath = "org.freedesktop.network1.Manager"
	likPath     = "org.freedesktop.network1.Link"
)

// New instantiates a new Networkd object allowing to manage systemd-networkd
func New() (*Networkd, error) {
	conn, err := dbus.SystemBus()
	if err != nil {
		return nil, err
	}
	return &Networkd{
		conn:   conn,
		sysObj: conn.Object("org.freedesktop.network1", dbus.ObjectPath("/org/freedesktop/network1")),
	}, nil
}

// Networkd Central type for invoking methods or reading properties from systemd-networkd
type Networkd struct {
	conn   *dbus.Conn
	sysObj dbus.BusObject
	io.Closer
}

// Close closes the underlying dbus connection
func (n *Networkd) Close() error {
	return n.conn.Close()
}

// Introspect list available methods
func (n *Networkd) Introspect() (*introspect.Node, error) {
	node, err := introspect.Call(n.conn.Object("org.freedesktop.network1", "/org/freedesktop/network1"))
	if err != nil {
		return nil, err
	}
	return node, nil
}

func (n *Networkd) ListLinks() ([]LinkResponse, error) {
	links := make([]LinkResponse, 0)
	err := n.sysObj.Call(managerPath+".ListLinks", 0).Store(&links)
	if err != nil {
		return nil, err
	}
	return links, nil
}

func (n *Networkd) LinkFromInterfaceName(name string) (*Link, error) {
	links, err := n.ListLinks()
	if err != nil {
		return nil, err
	}
	for _, l := range links {
		if l.Name != name {
			continue
		}
		return &Link{
			n.conn.Object("org.freedesktop.network1", l.Object),
		}, nil
	}
	return nil, errors.New("not found")
}

func (n *Networkd) ReconfigureLink(index int) error {
	return n.sysObj.Call(managerPath+".ReconfigureLink", 0, index).Err
}

func (n *Networkd) Reload() error {
	return n.sysObj.Call(managerPath+".Reload", 0).Err
}

func (n *Networkd) RenewLink(index int) error {
	return n.sysObj.Call(managerPath+".RenewLink", 0, index).Err
}
