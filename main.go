package main

import (
	"fmt"
	"net/http"
	"sendx-project/handlers"
)

const (
	PORT = 3000
)

func main() {
	http.HandleFunc("/", handlers.IndexHandler)
	http.HandleFunc("/user-login", handlers.UserLoginHandler)
	http.HandleFunc("/admin-login", handlers.AdminLoginHandler)
	http.HandleFunc("/crawl", handlers.CrawlHandler)
	http.HandleFunc("/config/numWorkers", handlers.SetNumWorkers)
	http.HandleFunc("/config/maxCrawlsPerHour", handlers.SetMaxCrawlsPerHour)
	http.HandleFunc("/config", handlers.GetConfig)
	http.HandleFunc("/get-config", handlers.GetConfigJSON)
	http.Handle("/html/", http.StripPrefix("/html/", http.FileServer(http.Dir("./html"))))
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("./css"))))
	http.Handle("/javascript/", http.StripPrefix("/javascript/", http.FileServer(http.Dir("./javascript"))))
	http.ListenAndServe(fmt.Sprintf(":%d", PORT), nil)
}
