package routers

import (
	"net/http"

	"github.com/gorilla/mux"
)

func SetupRouter() *mux.Router {
	router := mux.NewRouter()

	// Sub-rotas registradas no mesmo router
	SetupRouterUsuario(router)
	SetupRouterVendedor(router)
	SetupRouterTime(router)

	// Rota estática
	router.PathPrefix("/").Handler(
		http.StripPrefix("/", http.FileServer(
			http.Dir("./static/"))))

	return router
}
