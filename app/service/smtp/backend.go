package smtp

import "github.com/emersion/go-smtp"

type Backend struct{}

func NewBackend() *Backend {
	return &Backend{}
}

func (bkd *Backend) NewSession(c *smtp.Conn) (smtp.Session, error) {
	return &session{
		c: c,
	}, nil
}
