package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/cameronjpr/fplbuddy/handlers"
)

var dev = true

func disableCacheInDevMode(next http.Handler) http.Handler {
	if !dev {
		return next
	}
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Cache-Control", "no-store")
		next.ServeHTTP(w, r)
	})
}

func main() {
	r := http.NewServeMux()

	r.Handle("/static/",
		disableCacheInDevMode(
			http.StripPrefix("/static",
				http.FileServer(http.Dir("static")))))
	r.Handle("/assets/",
		disableCacheInDevMode(
			http.StripPrefix("/assets",
				http.FileServer(http.Dir("assets")))))
	r.Handle("/favicon.ico", disableCacheInDevMode(http.FileServer(http.Dir("static"))))
	r.HandleFunc("/", withLogging(handlers.HandleHomePage))
	r.HandleFunc("POST /search", withLogging(handlers.HandleSearchPlayers))
	r.HandleFunc("GET /players", withLogging(handlers.HandlePlayersPage))
	r.HandleFunc("GET /players/{playerId}", handleGetPlayer)
	r.HandleFunc("POST /players", handleSearchPlayers)
	r.HandleFunc("POST /team", handlePostTeam)
	r.HandleFunc("GET /team/{teamId}", withLogging(handlers.HandleGetEntryPage))
	r.HandleFunc("GET /team/{teamId}/refresh", withLogging(handlers.HandleGetEntryView))
	r.HandleFunc("GET /comparison", withLogging(handlers.HandleComparisonPage))
	r.HandleFunc("POST /comparison", withLogging(handlePostComparison))

	port := readPort()

	fmt.Printf("Listening on 0.0.0.0:%s\n", port)
	http.ListenAndServe(fmt.Sprintf("0.0.0.0:%s", port), r)
}

func readPort() string {
	port, ok := os.LookupEnv("PORT")
	if !ok {
		return "8080"
	}
	return port
}

func withLogging(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Request received: %s %s %s", r.Method, r.URL.Path, r.URL.RawQuery)
		next(w, r)
	})
}
