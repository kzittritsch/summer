package main

import (
	"github.com/gorilla/handlers"
	"github.com/julienschmidt/httprouter"
	"github.com/spf13/viper"
	"net/http"
	"os"
	"log"
)

func middleware(router http.Handler) http.Handler {
	// Logger
	if (viper.GetBool("logger")) == true {
		return handlers.CombinedLoggingHandler(os.Stdout, router)
	}

	return router
}

func authMiddleware(h httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		log.Println("Auth middleware called")
		path := setRoot(ps.ByName("path"))
		user, passwd, hasAuth := r.BasicAuth()

		if !hasAuth {
			errorHandler(w, r, "auth", 401, path)
			return
		}

		success := Authenticate(user, passwd)
		log.Println("Auth success: ", success)
		if success {
			h(w, r, ps)
			return
		} else {
			errorHandler(w, r, "auth", 401, path)
			return
			return
		}
	}
}
