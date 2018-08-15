package wildfire

import (
	"bufio"
	"fmt"
	"github.com/aaomidi/wildfire/config"
	"github.com/labstack/gommon/log"
	"io"
	"net"
)

const (
	SocksVersion = 5
)

func Serve(config config.Config) {
	listener, err := net.Listen("tcp", config.GetConnectionString())
	if err != nil {
		log.Error(err)
		return
	}

	serve(&listener)
}

func serve(listener *net.Listener) {
	for {
		conn, err := (*listener).Accept()

		if err != nil {
			log.Error(err)
		}

		go serveConnection(conn)
	}
}

func serveConnection(conn net.Conn) error {
	defer conn.Close()

	reader := bufio.NewReader(conn)

	version := []byte{0}
	nmethods := []byte{0}

	_, err := io.ReadFull(reader, version)

	if err != nil {
		return err
	}

	_, err = io.ReadFull(reader, nmethods)
	if err != nil {
		return err
	}

	if version[0] != SocksVersion {
		return fmt.Errorf("Unrecognized version number for connection.\n")
	}

}
