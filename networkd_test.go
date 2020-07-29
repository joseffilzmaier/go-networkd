package networkd

import (
	"fmt"
	"testing"
)

var n *Networkd

func init() {
	var err error
	n, err = New()
	if err != nil {
		panic(err)
	}
}

func TestListLink(t *testing.T) {
	links, err := n.ListLinks()
	if err != nil {
		t.Error(err)
	}
	fmt.Println(links)
}

func TestReconfigureWlan0(t *testing.T) {
	l, err := n.LinkFromInterfaceName("wlan0")
	if err != nil {
		t.Error(err)
	}
	l.Reconfigure()
}
