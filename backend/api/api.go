package api

import (
	"fmt"
	"net/http"

	"github.com/rg-km/final-project-engineering-45/backend/repository"
)

type API struct {
	usersRepo    repository.UserRepository
	fakultasRepo repository.FakultasRepository
	prodiRepo    repository.ProdiRepository
	mux          *http.ServeMux
}

func NewAPI(usersRepo repository.UserRepository, fakultasRepo repository.FakultasRepository, prodiRepo repository.ProdiRepository) API {
	mux := http.NewServeMux()
	api := API{
		usersRepo,
		fakultasRepo,
		prodiRepo,
		mux,
	}

	mux.Handle("/api/user/login", api.POST(http.HandlerFunc(api.login)))
	mux.Handle("/api/user/logout", api.POST(http.HandlerFunc(api.logout)))
	mux.Handle("/api/user/signup", api.POST(http.HandlerFunc(api.signup)))

	mux.Handle("/api/fakultas/list", api.GET(api.AuthMiddleWare(http.HandlerFunc(api.fakultasList))))
	mux.Handle("/api/prodi/list", api.GET(api.AuthMiddleWare(http.HandlerFunc(api.prodiList))))

	//mux.Handle("/api/prodi/add", api.POST(api.AuthMiddleWare(api.AdminMiddleware(http.HandlerFunc(api.addProdi)))))
	return api
}

func (api *API) Handler() *http.ServeMux {
	return api.mux
}

func (api *API) Start() {
	fmt.Println("Starting web server at http://localhost:8080/")
	http.ListenAndServe(":8080", api.Handler())
}
