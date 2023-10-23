package models

import (
	"sync"
	"time"
)

const (
	NON_PAYING_DELAY   = 2 * time.Second
	 PAYING_WORKERS = 5
     NON_PAYING_WORKERS = 2
	DEFAULT_MAX_CRAWLS = 100
)

type ServerState struct {
	NumWorkers           int
	MaxCrawlsPerHour     int
	PagesCrawledThisHour int
	LastCrawlReset       time.Time
	Mu                   sync.Mutex
}

var State = ServerState{
	MaxCrawlsPerHour: DEFAULT_MAX_CRAWLS,
	LastCrawlReset:   time.Now(),
}

var CrawledPages = make(map[string]PageData)
