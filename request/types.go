package request

import (
	"fmt"
	"io"
	"net"

	"github.com/aaomidi/wildfire"
	"github.com/aaomidi/wildfire/authentication"
)

type Command byte

const (
	Connect = Command(0x01)
)

// A Request represents request received by a server
type Request struct {
	// Protocol version
	Version byte
	// Requested command
	Command Command
	// AuthContext provided during negotiation
	AuthContext *authentication.AuthContext
	// AddrSpec of the the network that sent the request
	RemoteAddr *AddrSpec
	// AddrSpec of the desired destination
	DestAddr *AddrSpec
	reader   *io.Reader
}

func (c *Command) toByte() byte {
	return byte(*c)
}

func NewRequest(reader *io.Reader, conn *net.Conn) (*Request, error) {
	header := []byte{0, 0, 0}

	if _, err := io.ReadAtLeast(reader, header, 3); err != nil {
		return nil, fmt.Errorf("failed to get command version: %v", err)
	}

	if header[0] != wildfire.SocksVersion {
		return nil, fmt.Errorf("unsupported version: %v", header[0])
	}

	// Find client's destination IP
	dest, err := GetAddrSpec(reader)

	// Find client's IP
	addr := (*conn).RemoteAddr()
	remote := GetAddrFromAddr(&addr)

	if err != nil {
		return nil, err
	}

	request := &Request{
		Version:    header[0],
		Command:    Command(header[1]),
		RemoteAddr: remote,
		DestAddr:   dest,
		reader:     reader,
	}

	return request, nil
}

func HandleRequest(request *Request, conn *net.Conn) {
	switch request.Command {
	case Connect:
		handleConnect(request, conn)
	}

}
