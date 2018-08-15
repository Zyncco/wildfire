package request

import (
	"errors"
	"fmt"
	"io"
	"net"
)

const (
	ipv4AddressType = byte(0x01)
	fqdnAddressType = byte(0x03)
	ipv6AddressType = byte(0x04)
)

type AddrSpec struct {
	FQDN string
	IP   net.IP
	Port uint8
}

func (a *AddrSpec) String() string {
	if a.FQDN != "" {
		return fmt.Sprintf("%s:%d", a.FQDN, a.Port)
	}
	return fmt.Sprintf("%s:%d", a.IP, a.Port)
}

func GetAddrSpec(reader *io.Reader) (*AddrSpec, error) {
	var address *AddrSpec

	addressType := make([]byte, 1)

	if _, err := (*reader).Read(addressType); err != nil {
		return nil, err
	}

	switch addressType[0] {
	case ipv4AddressType:
		fallthrough
	case ipv6AddressType:
		length := 4

		if addressType[0] == ipv6AddressType {
			length = 16
		}

		ipAddress := make([]byte, length)

		if _, err := io.ReadAtLeast(*reader, ipAddress, length); err != nil {
			return nil, err
		}

		address = &AddrSpec{
			IP: net.IP(ipAddress),
		}
	case fqdnAddressType:
		addressLength := make([]byte, 1)

		if _, err := (*reader).Read(addressLength); err != nil {
			return nil, err
		}

		fqdn := make([]byte, addressLength[0])

		if _, err := io.ReadAtLeast(*reader, fqdn, int(addressLength[0])); err != nil {
			return nil, err
		}

		address = &AddrSpec{
			FQDN: string(fqdn),
		}
	}

	if address == nil {
		return nil, errors.New("could not read address host")
	}

	port := make([]byte, 2)

	if _, err := io.ReadAtLeast(*reader, port, 2); err != nil {
		return nil, err
	}

	address.Port = uint8((int(port[0]) << 8) | int(port[1]))

	return address, nil
}

func GetAddrFromAddr(p *net.Addr) *AddrSpec {
	addr := *p
	spec := AddrSpec{}

	switch addr := addr.(type) {
	case *net.UDPAddr:
		spec.IP = addr.IP
		spec.Port = uint8(addr.Port)
	case *net.TCPAddr:
		spec.IP = addr.IP
		spec.Port = uint8(addr.Port)
	}
	return &spec
}
