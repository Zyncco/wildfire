package authentication

import "io"

type Method byte

const (
	NoAuth          = Method(0x00)
	UserAuthVersion = Method(0x01)
	UserPassAuth    = Method(0x02)
	NoAcceptable    = Method(0xff)
)

const (
	AuthSuccess = Method(0x00)
	AuthFailure = Method(0x01)
)

var (
	authenticators map[Method]Authenticator
)

type Authenticator interface {
	// What should it return?
	// Should we write the results directly into it?
	Authenticate(reader io.Reader, writer io.Writer) (*AuthContext, error)
	GetMethodCode() Method
}

type AuthContext struct {
	Method Method
}

func GetAuthenticator(method Method) Authenticator {
	prepareAuthenticatorMap()

	return authenticators[method]
}

func prepareAuthenticatorMap() {
	if authenticators != nil {
		return
	}

	authenticators = make(map[Method]Authenticator)
	authenticators[NoAuth] = &NoAuthAuthenticator{}
}

func (m *Method) ToByte() byte {
	return byte(*m)
}

func (m *Method) GetAuthenticator() Authenticator {
	return GetAuthenticator(*m)
}
