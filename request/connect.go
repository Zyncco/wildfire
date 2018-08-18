package request

import (
	"io"
	"net"
)

func handleConnect(request *Request, conn *net.Conn) error {
	target, err := net.Dial("tcp", request.DestAddr.String())

	if err != nil {
		return err
	}

	errChan := make(chan error, 2)

	go proxy(conn, target, errChan)
	go proxy(request.reader, target, errChan)

	// Wait
	for i := 0; i < 2; i++ {
		e := <-errChan
		if e != nil {
			// return from this function closes target (and conn).
			return e
		}
	}
	return nil
}

func sendConnectReply() {

}

type closeWriter interface {
	CloseWrite() error
}

func proxy(src io.Reader, dst io.Writer, errChan chan error) {
	// This is where the actual proxying happens, and where we plug in our plugins.
	_, err := io.Copy(dst, src)

	if tcpConn, ok := dst.(closeWriter); ok {
		tcpConn.CloseWrite()
	}
	errChan <- err
}
