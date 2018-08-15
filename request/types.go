package request

import (
	"fmt"
	"io"
	"net"

	"github.com/aaomidi/wildfire"
	"github.com/aaomidi/wildfire/authentication"
)

type Command byte

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
		return nil, fmt.Errorf("Failed to get command version: %v", err)
	}

	if header[0] != wildfire.SocksVersion {
		return nil, fmt.Errorf("Unsupported version: %v", header[0])
	}

	dest, err := GetAddrSpec(reader)

	if err != nil {
		return nil, err
	}

	(*conn).RemoteAddr()

	request := &Request{
		Version:  header[0],
		Command:  Command(header[1]),
		DestAddr: &dest,
		reader:   reader,
	}

	return request, nil
}

func HandleRequest(request *Request) {

}
