package handlers

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"sendx-project/models"
	"time"
)

func worker(jobs <-chan models.CrawlJob) {
	for job := range jobs {
		if !job.IsPaying {
			time.Sleep(models.NON_PAYING_DELAY)
		}
		job.Result <- crawlPage(job.Url)
	}
}

func CrawlHandler(w http.ResponseWriter, r *http.Request) {
	models.State.Mu.Lock()
	if time.Since(models.State.LastCrawlReset) > time.Hour {
		models.State.PagesCrawledThisHour = 0
		models.State.LastCrawlReset = time.Now()
	}
	if models.State.PagesCrawledThisHour >= models.State.MaxCrawlsPerHour {
		models.State.Mu.Unlock()
		http.Error(w, "Hourly crawl limit exceeded", http.StatusTooManyRequests)
		return
	}
	models.State.PagesCrawledThisHour++
	models.State.Mu.Unlock()

	urls, ok := r.URL.Query()["url"]
	if !ok || len(urls[0]) < 1 {
		http.Error(w, "Url Param 'url' is missing", http.StatusBadRequest)
		return
	}
	url := urls[0]
	isPaying := r.URL.Query().Get("isPaying") == "true"

	// Check if we have a cache for this URL from the last 60 minutes.
	if page, ok := models.CrawledPages[url]; ok {
		if time.Now().Unix()-page.Timestamp < 3600 { // 3600 seconds = 60 minutes
			w.Write([]byte(page.Data))
			fmt.Println("Retrieved data from cache for URL:", url)
			return
		}
	}

	// If not retrieved from cache, then it's a fresh crawl.
	fmt.Println("Crawling data for URL:", url)

	workersCount := models.NON_PAYING_WORKERS
	if isPaying {
		workersCount = models.PAYING_WORKERS
	}

	jobs := make(chan models.CrawlJob, workersCount)
	for i := 0; i < workersCount; i++ {
		go worker(jobs)
	}

	result := make(chan string)
	jobs <- models.CrawlJob{Url: url, IsPaying: isPaying, Result: result}
	pageData := <-result

	currentTime := time.Now().Unix()
	models.CrawledPages[url] = models.PageData{
		Timestamp: currentTime,
		Data:      pageData,
	}
	w.Write([]byte(pageData))
	close(jobs)
}

// func crawlPage(url string) string {
// 	resp, err := http.Get(url)
// 	if err != nil {
// 		return err.Error()
// 	}
// 	defer resp.Body.Close()
// 	body, _ := ioutil.ReadAll(resp.Body)
// 	return string(body)
// }

const MAX_RETRIES = 3
const RETRY_DELAY = 5 * time.Second

func crawlPage(url string) string {
	var lastError error

	for i := 0; i < MAX_RETRIES; i++ {
		resp, err := http.Get(url)
		if err == nil {
			defer resp.Body.Close()
			body, err := ioutil.ReadAll(resp.Body)
			if err == nil {
				return string(body)
			}
			// If there's an error reading the body, set lastError and continue.
			lastError = err
		} else {
			// If there's an error making the request, set lastError and continue.
			lastError = err
		}

		// If there was an error, wait for the retry delay before the next retry.
		time.Sleep(RETRY_DELAY)
	}

	// If all retries failed, return the last error encountered.
	return lastError.Error()
}
