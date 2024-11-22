package main

import (
	"SCS_project/internal/config"
	"SCS_project/internal/handlers"
	"SCS_project/internal/helpers"
	"fmt"
	"github.com/alexedwards/scs/v2"
	"log"
	"net/http"
	"os"
	"time"
)

const portNumber = ":8080"
const inProduction = false

var app config.AppConfig
var session *scs.SessionManager

var infoLog *log.Logger
var errorLog *log.Logger

// main is the main entry point in the web
func main() {
	repo := handlers.NewRepo(&app)

	run(repo)

	server := &http.Server{
		Addr:    portNumber,
		Handler: routes(),
	}

	fmt.Println(fmt.Sprintf("Starting application on port %s", portNumber))

	err := server.ListenAndServe()
	log.Fatal(err)
}

func run(repo *handlers.Repository) {
	setUpSessions()
	setUpAppConfig(repo)

	helpers.NewHelpers(&app)
	handlers.NewHandlers(repo)
}

// setUpAppConfig sets up the web configuration (app)
func setUpAppConfig(repo *handlers.Repository) {
	app.InProduction = inProduction

	// TODO() Persist logs throughout server lifetime
	infoLog = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	app.InfoLog = infoLog

	errorLog = log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	app.ErrorLog = errorLog

	app.Session = session
}

func setUpSessions() {
	// TODO() Register any models for the session
	//gob.Register(models.Reservation{})

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction
}
