package main

import (
	"net/http"

	"github.com/justinas/alice"
	"github.com/makeict/MESSforMakers/controllers"

	"github.com/aaronarduino/MESSforMakers/session"
	"github.com/gorilla/mux"
)

// database connection, cookie store, etc..
type application struct {
	cookieStore    *session.CookieStore
	commonHandlers alice.Chain
}

func newApplication() *application {
	return &application{
		cookieStore:    session.NewCookieStore("mess-data"),
		commonHandlers: alice.New(loggingHandler),
	}
}

func (a *application) appRouter() http.Handler {
	router := mux.NewRouter()
	userC := controllers.User

	router.HandleFunc("/", RootHandler)
	router.HandleFunc("/user", userC.Index)

	return a.commonHandlers.Then(router)
}
