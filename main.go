package main

import (
	"datastar-reactive/internal/data"
	"datastar-reactive/internal/service"
	"encoding/gob"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/alexedwards/scs/v2"
	_ "github.com/mattn/go-sqlite3"
)

// register the SessionMap type on initilization so the SessionManager can acesse it.
func init() {
	gob.Register(&data.SessionMap{})
}

// config defines a configuration struct for the application
type config struct {
	srv struct {
		port         int
		env          string
		readTimeout  string
		writeTimeout string
		idleTimeout  string
	}
	db struct {
		dsn string
	}
}

// application holds the configuration struct as wall as the logs, database connection poll and the session manager instance.
type application struct {
	config
	infoLog        *log.Logger
	errLog         *log.Logger
	services       *service.Services
	sessionManager *scs.SessionManager
}

func main() {

	// ------------------------------------------------------------------------------------------------------
	// APPLICATION INITIALIZATION
	// ------------------------------------------------------------------------------------------------------
	// start the appliction logs
	app := application{
		infoLog: log.New(os.Stdout, "💬 INFO ", log.Ldate|log.Ltime|log.Lmicroseconds),
		errLog:  log.New(os.Stderr, "🛑 ERROR  ", log.Ldate|log.Ltime|log.Lmicroseconds|log.Lshortfile),
	}

	// ------------------------------------------------------------------------------------------------------
	// COMMAND LINE FLAGS DEFINITION
	// ------------------------------------------------------------------------------------------------------
	// command line definition for various configurations
	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "%s - Datastar Reactive Examples\n", os.Args[0])
		flag.PrintDefaults()
	}
	flag.IntVar(&app.srv.port, "srv-port", 5000, "Sets the application's server port")
	flag.StringVar(&app.srv.env, "srv-env", "development", "Environment (development|production)")
	flag.StringVar(&app.srv.readTimeout, "srv-read-timeout", "30s", "Server's read-timeout limit ")
	flag.StringVar(&app.srv.writeTimeout, "srv-write-timeout", "60s", "Server's write-timeout limit")
	flag.StringVar(&app.srv.idleTimeout, "srv-idle-timeout", "90s", "Server's  idle-timeout limit")
	flag.StringVar(&app.db.dsn, "db-dsn", "", "Server's  idle-timeout limit")
	flag.Parse()

	// ------------------------------------------------------------------------------------------------------
	// DATABASE INITIALIZATION
	// ------------------------------------------------------------------------------------------------------
	// start the database connection pool
	db, err := service.OpenDB("sqlite3", app.db.dsn)
	if err != nil {
		app.errLog.Println("🛑 error initializing DB: ", "err", err.Error())
		os.Exit(1)
	}
	app.services = service.NewServices(db)
	app.infoLog.Println("🗃️ Database initialized successfuly")

	// ------------------------------------------------------------------------------------------------------
	// SESSION MANAGER INITIALIZATION
	// ------------------------------------------------------------------------------------------------------
	// starts the session manager
	sessionManager := scs.New()
	sessionManager.Lifetime = 24 * time.Hour
	sessionManager.Cookie.Secure = false
	app.sessionManager = sessionManager

	// ------------------------------------------------------------------------------------------------------
	// SERVER INITIALIZATION
	// ------------------------------------------------------------------------------------------------------
	server := http.Server{
		Addr:         fmt.Sprintf(":%d", app.srv.port),
		ReadTimeout:  app.parseStringDuration(app.srv.readTimeout),
		WriteTimeout: app.parseStringDuration(app.srv.writeTimeout),
		IdleTimeout:  app.parseStringDuration(app.srv.idleTimeout),
		ErrorLog:     app.errLog,
		Handler:      app.router(),
	}

	app.infoLog.Printf("🚀 server starting on %s mode in port %d", app.srv.env, app.srv.port)
	log.Fatal(server.ListenAndServe())
}
