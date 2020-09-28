package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"
	"path/filepath"
	"time"

	redis "github.com/go-redis/redis/v8"
	"github.com/gorilla/mux"
)

const (
	// APIVersion is used throughout the application in responses
	// and set our current route responder for API actions
	APIVersion = "v1"
)

type object map[string]interface{}

type intrusionHandler struct{}

func (i intrusionHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	bodyBytes, bodyBytesErr := ioutil.ReadAll(r.Body)

	if bodyBytesErr != nil {
		return
	}

	resp := &object{
		"body":    string(bodyBytes),
		"url":     r.URL.Path,
		"headers": r.Header,
	}

	jsonResp, jsonRespErr := json.Marshal(resp)

	if jsonRespErr != nil {
		return
	}

	w.WriteHeader(http.StatusBadRequest)
	w.Write(jsonResp)
}

type spaHandler struct {
	staticPath string
	indexPath  string
}

func (h spaHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Get the absolute path to prevent directory traversal
	path, err := filepath.Abs(r.URL.Path)
	if err != nil {
		// If we failed to get the absolute path respond with a 400 bad request
		// and return
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Prepend the path with the path to the static directory
	path = filepath.Join(h.staticPath, path)

	// Check whether a file exists at the given path
	_, err = os.Stat(path)
	if os.IsNotExist(err) {
		// File does not exist, serve index.html
		http.ServeFile(w, r, filepath.Join(h.staticPath, h.indexPath))
		return
	} else if err != nil {
		// If we got an error (that wasn't that the file doesn't exist) stating the
		// file, return a 500 internal server error and stop
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Otherwise, use http.FileServer to serve the static dir
	http.FileServer(http.Dir(h.staticPath)).ServeHTTP(w, r)
}

type app struct {
	Router *mux.Router
	Redis  *redis.Client
}

func main() {
	env := os.Getenv("ENVIRONMENT")
	redisConnection := &redis.Options{
		Addr:     "redis:6379",
		Password: "",
		DB:       0,
	}
	redisClient := redis.NewClient(redisConnection)
	app := &app{
		Router: mux.NewRouter(),
		Redis:  redisClient,
	}

	route(app)

	// Debugger
	if env == "local" {
		go func() {
			log.Println("Tracing started on port", os.Getenv("TRACE_PORT"), ": /debug/pprof")
			http.ListenAndServe(":"+os.Getenv("TRACE_PORT"), http.DefaultServeMux)
		}()
	}

	srv := &http.Server{
		Handler: app.Router,
		Addr:    "0.0.0.0:" + os.Getenv("WEB_PORT"),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())

}

func route(app *app) *mux.Router {
	intrusion := intrusionHandler{}
	spa := spaHandler{staticPath: "client/build", indexPath: "index.html"}

	api := app.Router.PathPrefix("/api/" + APIVersion).Subrouter()
	api.Handle("/sample", sampleHandler(app)).Methods("GET")
	api.PathPrefix("/").Handler(intrusion)

	// Static Client
	app.Router.PathPrefix("/").Handler(spa)

	return app.Router
}

func sampleHandler(app *app) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Sample handler hit")
	})
}
