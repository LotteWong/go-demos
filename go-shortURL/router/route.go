package router

import (
	"fmt"
	"go-shortURL/configs"
	"go-shortURL/controller"
	"go-shortURL/middlewares"
	"go-shortURL/models"
	"go-shortURL/utils"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/justinas/alice"
)

type App struct {
	Router     *mux.Router
	Middleware *middlewares.Middleware
}

func NewApp() *App {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	return &App{
		Router:     mux.NewRouter(),
		Middleware: &middlewares.Middleware{},
	}
}

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	utils.RespondWithJson(w, http.StatusOK, &models.HealthCheckDto{Msg: "pong"})
}

func (a *App) Register() {
	m := alice.New(a.Middleware.LogHandler, a.Middleware.RecoverHandler)
	c := controller.GetUrlController()

	a.Router.Handle("/ping", m.ThenFunc(HealthCheck)).Methods("GET")

	a.Router.Handle("/url/shortlink", m.ThenFunc(c.CreateShortLink)).Methods("POST")
	a.Router.Handle("/{shortlink:[a-zA-Z0-9]{1,11}}", m.ThenFunc(c.RedirectOriginalUrl)).Methods("GET")
	a.Router.Handle("/url/detail", m.ThenFunc(c.GetShortLinkDetail)).Methods("GET")
}

func (a *App) Run() {
	a.Register()

	config := &configs.Config{}
	config.ParseConfig()
	addr := fmt.Sprintf("%s:%d", config.Ip, config.Port)

	log.Fatal(http.ListenAndServe(addr, a.Router))
}
