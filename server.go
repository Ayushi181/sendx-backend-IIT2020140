package main

// import (
// 	"encoding/json"
// 	"fmt"
// 	"io/ioutil"
// 	"net/http"
// 	"strconv"
// 	"sync"
// 	"time"
// )

// const (
// 	PORT               = 3000
// 	NON_PAYING_DELAY   = 2 * time.Second
// 	DEFAULT_WORKERS    = 5
// 	DEFAULT_MAX_CRAWLS = 100
// )

// var crawledPages = make(map[string]PageData)

// type PageData struct {
// 	Timestamp int64
// 	Data      string
// }

// type CrawlJob struct {
// 	url      string
// 	isPaying bool
// 	result   chan string
// }

// type ServerState struct {
// 	numWorkers           int
// 	maxCrawlsPerHour     int
// 	pagesCrawledThisHour int
// 	lastCrawlReset       time.Time
// 	mu                   sync.Mutex
// }

// var state = ServerState{
// 	numWorkers:       DEFAULT_WORKERS,
// 	maxCrawlsPerHour: DEFAULT_MAX_CRAWLS,
// 	lastCrawlReset:   time.Now(),
// }

// func worker(jobs <-chan CrawlJob) {
// 	for job := range jobs {
// 		if !job.isPaying {
// 			time.Sleep(NON_PAYING_DELAY)
// 		}
// 		job.result <- crawlPage(job.url)
// 	}
// }

// func main() {
// 	http.HandleFunc("/", indexHandler)
// 	http.HandleFunc("/user-login", userLoginHandler)
// 	http.HandleFunc("/admin-login", adminLoginHandler)
// 	http.HandleFunc("/crawl", crawlHandler)
// 	http.HandleFunc("/config/numWorkers", setNumWorkers)
// 	http.HandleFunc("/config/maxCrawlsPerHour", setMaxCrawlsPerHour)
// 	http.HandleFunc("/config", getConfig)
// 	http.HandleFunc("/get-config", getConfigJSON)
// 	http.Handle("/html/", http.StripPrefix("/html/", http.FileServer(http.Dir("./html"))))
// 	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("./css"))))
// 	http.Handle("/javascript/", http.StripPrefix("/javascript/", http.FileServer(http.Dir("./javascript"))))
// 	fmt.Println("Server Running on ", PORT)
// 	http.ListenAndServe(fmt.Sprintf(":%d", PORT), nil)
// }

// func indexHandler(w http.ResponseWriter, r *http.Request) {
// 	http.ServeFile(w, r, "./html/login.html")
// }

// func userLoginHandler(w http.ResponseWriter, r *http.Request) {
// 	http.ServeFile(w, r, "./html/user.html")
// }

// func adminLoginHandler(w http.ResponseWriter, r *http.Request) {
// 	http.ServeFile(w, r, "./html/admin.html")
// }

// func crawlHandler(w http.ResponseWriter, r *http.Request) {
// 	state.mu.Lock()
// 	if time.Since(state.lastCrawlReset) > time.Hour {
// 		state.pagesCrawledThisHour = 0
// 		state.lastCrawlReset = time.Now()
// 	}
// 	if state.pagesCrawledThisHour >= state.maxCrawlsPerHour {
// 		state.mu.Unlock()
// 		http.Error(w, "Hourly crawl limit exceeded", http.StatusTooManyRequests)
// 		return
// 	}
// 	state.pagesCrawledThisHour++
// 	state.mu.Unlock()

// 	urls, ok := r.URL.Query()["url"]
// 	if !ok || len(urls[0]) < 1 {
// 		http.Error(w, "Url Param 'url' is missing", http.StatusBadRequest)
// 		return
// 	}
// 	url := urls[0]
// 	isPaying := r.URL.Query().Get("isPaying") == "true"

// 	workersCount := state.numWorkers
// 	if isPaying {
// 		workersCount = DEFAULT_WORKERS
// 	}

// 	jobs := make(chan CrawlJob, workersCount)
// 	for i := 0; i < workersCount; i++ {
// 		go worker(jobs)
// 	}

// 	result := make(chan string)
// 	jobs <- CrawlJob{url: url, isPaying: isPaying, result: result}
// 	pageData := <-result

// 	currentTime := time.Now().Unix()
// 	crawledPages[url] = PageData{
// 		Timestamp: currentTime,
// 		Data:      pageData,
// 	}
// 	w.Write([]byte(pageData))
// 	close(jobs)
// }

// func crawlPage(url string) string {
// 	resp, err := http.Get(url)
// 	if err != nil {
// 		return err.Error()
// 	}
// 	defer resp.Body.Close()
// 	body, _ := ioutil.ReadAll(resp.Body)
// 	return string(body)
// }

// func setNumWorkers(w http.ResponseWriter, r *http.Request) {
// 	if r.Method != http.MethodPost {
// 		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
// 		return
// 	}

// 	body, _ := ioutil.ReadAll(r.Body)
// 	newNumWorkers, err := strconv.Atoi(string(body))
// 	fmt.Println("number of workers assigned = ", newNumWorkers)
// 	if err != nil {
// 		http.Error(w, "Invalid number provided", http.StatusBadRequest)
// 		return
// 	}

// 	state.mu.Lock()
// 	state.numWorkers = newNumWorkers
// 	state.mu.Unlock()

// 	w.Write([]byte("Number of workers updated successfully"))
// }

// func setMaxCrawlsPerHour(w http.ResponseWriter, r *http.Request) {
// 	if r.Method != http.MethodPost {
// 		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
// 		return
// 	}

// 	body, _ := ioutil.ReadAll(r.Body)
// 	newMaxCrawls, err := strconv.Atoi(string(body))
// 	fmt.Println("new max crawls assigned = ", newMaxCrawls)
// 	if err != nil {
// 		http.Error(w, "Invalid number provided", http.StatusBadRequest)
// 		return
// 	}

// 	state.mu.Lock()
// 	state.maxCrawlsPerHour = newMaxCrawls
// 	state.mu.Unlock()

// 	w.Write([]byte("Max crawls per hour updated successfully"))
// 	fmt.Println("Max crawls per hour updated successfully")
// }

// func getConfig(w http.ResponseWriter, r *http.Request) {
// 	state.mu.Lock()
// 	defer state.mu.Unlock()

// 	config := fmt.Sprintf("Number of Workers: %d, Max Crawls per Hour: %d", state.numWorkers, state.maxCrawlsPerHour)
// 	w.Write([]byte(config))
// }

// func getConfigJSON(w http.ResponseWriter, r *http.Request) {
// 	state.mu.Lock()
// 	defer state.mu.Unlock()

// 	// Create a map to hold the configuration
// 	configData := map[string]int{
// 		"numWorkers":       state.numWorkers,
// 		"maxCrawlsPerHour": state.maxCrawlsPerHour,
// 	}

// 	// Convert the map to JSON
// 	configJSON, err := json.Marshal(configData)
// 	if err != nil {
// 		http.Error(w, "Error creating JSON", http.StatusInternalServerError)
// 		return
// 	}

// 	// Set the response header for JSON output
// 	w.Header().Set("Content-Type", "application/json")
// 	w.Write(configJSON)
// }
