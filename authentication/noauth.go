package authentication

import (
	"github.com/aaomidi/wildfire"
	"io"
)

type NoAuthAuthenticator struct{}

func (a *NoAuthAuthenticator) Authenticate(reader io.Reader, writer io.Writer) (*AuthContext, error) {
	_, err := writer.Write([]byte{wildfire.SocksVersion, NoAuth.ToByte()})

	if err != nil {
		return nil, err
	}

	return &AuthContext{NoAuth}, nil
}

func (a *NoAuthAuthenticator) GetMethodCode() Method {
	return NoAuth
}
