package internal

import (
	"io/fs"
	"net/http"

	"github.com/gorilla/mux"

	wp "github.com/to-com/wp"
)

func (a *Application) routes() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/v1/retailers/{retailerId}/mfcs/{mfcId}/wp", a.handler.Getwp).Methods(http.MethodGet, http.MethodOptions)
	router.HandleFunc("/v1/retailers/{retailerId}/mfcs/{mfcId}/wp", a.handler.Createwp).Methods(http.MethodPost, http.MethodOptions)
	router.HandleFunc("/v1/triggers:generate", a.handler.generateTriggers).Methods(http.MethodPost)
	router.HandleFunc("/v1/triggers:fire", a.handler.FireTriggers).Methods(http.MethodPost)
	router.HandleFunc("/v1/triggers", a.handler.GetTriggers).Methods(http.MethodGet)

	swaggerAssetsContent, err := fs.Sub(fs.FS(wp.EmbedAPIAssets), "api")
	if err != nil {
		a.logger.Infof("failed to init FS for swagger assets, err: %+v", err)
	}
	fileServer := http.StripPrefix("/", http.FileServer(http.FS(swaggerAssetsContent)))
	router.PathPrefix("/").Handler(fileServer)

	return router
}
