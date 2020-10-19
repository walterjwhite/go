package main

import (
	"errors"
	"flag"
	"github.com/walterjwhite/go-application/libraries/application"
	"github.com/walterjwhite/go-application/libraries/application/logging"

	"github.com/rs/zerolog/log"
	"github.com/walterjwhite/go-application/libraries/utils/web/chromedpexecutor/plugins/walgreens"
)

var (
	session = &walgreens.Session{}
)

func init() {
	application.ConfigureWithProperties(session)

	log.Info().Msgf("session: %v", session)
}

func main() {
	defer application.OnEnd()

	if len(flag.Args()) < 1 {
		logging.Panic(errors.New("Command is required (login, logout, upload)"))
	}

	switch flag.Args()[0] {
	case "login":
		log.Info().Msgf("username: %v", session.Credentials.Username)
		session.Authenticate(application.Context)
	case "logout":
		session.Logout()
	case "upload":
		session.Upload(flag.Args()[1:]...)
	}
}
