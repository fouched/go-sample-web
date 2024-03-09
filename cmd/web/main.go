package main

import (
	"fmt"
	"github.com/alexedwards/scs/v2"
	"github.com/fouched/go-sample-web/internal/config"
	"github.com/fouched/go-sample-web/internal/handlers"
	"github.com/fouched/go-sample-web/internal/render"
	"log"
	"net/http"
	"time"
)

const port = ":8000"

var app config.AppConfig
var session *scs.SessionManager

func main() {

	run()

	srv := &http.Server{
		Addr:    port,
		Handler: routes(),
	}
	fmt.Println(fmt.Sprintf("Starting application on %s", port))

	err := srv.ListenAndServe()
	if err != nil {
		log.Fatalln(err)
	}
}

func run() {
	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session
	app.InProduction = false

	hc := handlers.NewConfig(&app)
	handlers.NewHandlers(hc)
	render.NewRenderer(&app)
}
