package main

import (
	"flag"
	"github.com/rs/zerolog/log"
	"github.com/walterjwhite/go/lib/application"
	"github.com/walterjwhite/go/lib/application/logging"
	"github.com/walterjwhite/go/lib/application/property"
	emaill "github.com/walterjwhite/go/lib/net/email"
	"github.com/walterjwhite/go/lib/utils/web/chromedpexecutor/plugins/gateway"
	"github.com/walterjwhite/go/lib/utils/web/chromedpexecutor/plugins/gateway/email"
	"time"
)

var (
	// TODO: randomize the interval, configure minimum interval and deviation ...
	tickleInterval = flag.String("TickleInterval", "3m", "Tickle Interval")

	session       = &gateway.Session{}
	emailInstance = &email.Provider{EmailSenderAccount: &emaill.EmailSenderAccount{}}
)

func init() {
	application.Configure()

	// configure email
	property.Load(emailInstance, "")
	log.Info().Msgf("emailInstance: %v", *emailInstance)

	property.Load(emailInstance.EmailSenderAccount, "")
	log.Info().Msgf("emailInstance: %v", *emailInstance.EmailSenderAccount)

	property.Load(session, "")

	log.Info().Msgf("session: %v", *session)
	property.Load(session.Credentials, "")
	log.Info().Msgf("session: %v", *session)

	i, err := time.ParseDuration(*tickleInterval)
	logging.Panic(err)

	session.Tickle = &gateway.Tickle{TickleInterval: &i}
}

func main() {
	session.Token = emailInstance.Get()
	session.Run(application.Context)

	application.Wait()
}
