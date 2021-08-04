package main

import (
	"log"
	"net/http"

	"gopkg.in/oauth2.v3/errors"
	"gopkg.in/oauth2.v3/manage"
	"gopkg.in/oauth2.v3/models"
	"gopkg.in/oauth2.v3/server"
	"gopkg.in/oauth2.v3/store"
)

func main() {
	manager := manage.NewDefaultManager()

	//token memory store
	manager.MustTokenStorage(store.NewMemoryTokenStore())
	clientStore := store.NewClientStore()
	clientStore.Set("232323", &models.Client{
		ID:     "000000",
		Secret: "999999",
		Domain: "http://localhost",
	})

	manager.MapClientStorage(clientStore)

	srv := server.NewDefaultServer(manager)
	srv.SetAllowGetAccessRequest(true)
	srv.SetClientInfoHandler(server.ClientFormHandler)

	srv.SetInternalErrorHandler(func(err error) (res *errors.Response) {
		log.Println("Internal error: ", err.Error())
		return
	})

	srv.SetResponseErrorHandler(func(res *errors.Response) {
		log.Println("Response Error:", res.Error.Error())
	})

	http.HandleFunc("/authorize", func(w http.ResponseWriter, r *http.Request) {
		err := srv.HandleAuthorizeRequest(w, r)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
	})

	http.HandleFunc("/token", func(w http.ResponseWriter, r *http.Request) {
		log.Println(r)
		srv.HandleTokenRequest(w, r)
	})

	log.Fatal(http.ListenAndServe(":9000", nil))
}
