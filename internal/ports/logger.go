package ports

import log "github.com/sirupsen/logrus"

// Logger creates a PORT for the logger
type Logger interface {
	MakeLogger(filename string, display bool) *log.Logger
	SetFormater()
	Hook() *log.Logger
}
