package authentication

import "io"

type Authenticator interface {
	// What should it return?
	// Should we write the results directly into it?
	Authenticate(reader io.Reader, writer io.Writer)
}
