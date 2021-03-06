package property

import (
	"flag"
	"github.com/octago/sflags/gen/gflag"
	"github.com/walterjwhite/go/lib/application/logging"
)

func LoadCli(config interface{}) {
	logging.Panic(gflag.ParseToDef(config))
	flag.Parse()
}
