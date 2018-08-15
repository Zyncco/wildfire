package request

import (
	"fmt"
	"io"
	"net"
)

type AddrSpec struct {
	Host string
	IP   net.IP
	Port uint8
}

func (a *AddrSpec) String() string {
	if a.Host != "" {
		return fmt.Sprintf("%s (%s):%d", a.Host, a.IP, a.Port)
	}
	return fmt.Sprintf("%s:%d", a.IP, a.Port)
}

func GetAddrSpec(reader *io.Reader) (AddrSpec, error) {

}
