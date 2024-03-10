package main

import (
	"database/sql"
	"fmt"
	"github.com/alexedwards/scs/v2"
	"github.com/fouched/go-sample-web/internal/config"
	"github.com/fouched/go-sample-web/internal/handlers"
	"github.com/fouched/go-sample-web/internal/render"
	"github.com/fouched/go-sample-web/internal/repository"
	_ "github.com/jackc/pgx/v5"
	_ "github.com/jackc/pgx/v5/pgconn"
	_ "github.com/jackc/pgx/v5/stdlib"
	"log"
	"net/http"
	"os"
	"time"
)

const port = ":8000"

var app config.AppConfig
var session *scs.SessionManager
var infoLog *log.Logger
var errorLog *log.Logger

const dbString = "host=localhost port=5432 dbname=go_sample_web user=go password=gopher"

func main() {
	dbPool, err := run()
	if err != nil {
		log.Fatalln(err)
	}
	// we have database connectivity, close it after app stops
	defer dbPool.Close()

	srv := &http.Server{
		Addr:    port,
		Handler: routes(),
	}

	infoLog.Println(fmt.Sprintf("Starting application on %s", port))

	err = srv.ListenAndServe()
	if err != nil {
		log.Fatalln(err)
	}

}

func run() (*sql.DB, error) {
	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session
	app.InProduction = false

	// set up loggers
	infoLog = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	app.InfoLog = infoLog

	errorLog = log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	app.ErrorLog = errorLog

	dbPool, err := repository.CreateDbPool(dbString)
	if err != nil {
		errorLog.Println(err)
		log.Fatal("Cannot connect to database! Dying argh...")
	}

	hc := handlers.NewConfig(&app, dbPool)
	handlers.NewHandlers(hc)
	render.NewRenderer(&app)

	return dbPool, nil
}
