package routers

import (
	"http/example/factory"
	"http/example/helper"

	"net/http"
)

func RouterSetup(presenter factory.Presenter) {
	http.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "POST":
			presenter.UserPresenter.Register(w, r)
			return
		case "GET":
			presenter.UserPresenter.GetUser(w, r)
		default:
			helper.FailedMethodNotAllowed(w)
			return
		}
	})
	http.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "POST":
			presenter.UserPresenter.Login(w, r)
			return
		default:
			helper.FailedMethodNotAllowed(w)
			return
		}
	})
}
