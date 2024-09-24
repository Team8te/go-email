package main

import (
	"github.com/Team8te/go-email/app"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/pflag"
)

type cmdArgs struct {
	help bool
}

var args cmdArgs

func init() {
	pflag.BoolVarP(&args.help, "help", "h", false, "Show help message")
	pflag.Parse()
}

func main() {
	defer func() {
		if err := recover(); err != nil {
			log.Error(err)
		}
	}()

	if args.help {
		pflag.Usage()
		return
	}

	a := app.NewApp()
	err := a.Run()
	if err != nil {
		log.Errorf("App complite with error: %v", err)
	}
}
