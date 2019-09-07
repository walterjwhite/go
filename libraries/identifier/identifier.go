package identifier

import (
	"log"
)

// *MUST* be injected at compile time
var ApplicationVersion string
var BuildDate string
var GoVersion string
var OsArchitecture string

func Log() {
	if isConfigured() {
		log.Printf("Application Version: %v\n", ApplicationVersion)
		log.Printf("Built on: %v\n", BuildDate)
	}
}

func isConfigured() bool {
	return len(ApplicationVersion) > 0 && len(BuildDate) > 0
}
