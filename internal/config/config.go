package config

import (
	"github.com/alexedwards/scs/v2"
	"log"
)

type AppConfig struct {
	InfoLog      *log.Logger
	ErrorLog     *log.Logger
	InProduction bool
	Session      *scs.SessionManager
}
