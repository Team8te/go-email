package app

import (
	ssmtp "github.com/Team8te/go-email/app/service/smtp"
	"github.com/emersion/go-smtp"
)

type CmdArgs struct {
	Help bool
	Addr string
}

type App struct {
	s *smtp.Server
}

func NewApp(args *CmdArgs) *App {
	s := smtp.NewServer(ssmtp.NewBackend())
	s.Addr = args.Addr
	s.Domain = "localhost"
	return &App{
		s: s,
	}
}

func (a *App) Run() error {
	return a.s.ListenAndServe()
}
