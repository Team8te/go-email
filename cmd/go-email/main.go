package main

import (
	"github.com/Team8te/go-email/app"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/pflag"
)

var args app.CmdArgs

func init() {
	pflag.BoolVarP(&args.Help, "help", "h", false, "Show help message")
	pflag.StringVar(&args.Addr, "a", "127.0.0.1:1025", "Listen address")
	pflag.Parse()
}

func main() {
	defer func() {
		if err := recover(); err != nil {
			log.Error(err)
		}
	}()

	if args.Help {
		pflag.Usage()
		return
	}

	a := app.NewApp(&args)
	err := a.Run()
	if err != nil {
		log.Errorf("App complite with error: %v", err)
	}
}
