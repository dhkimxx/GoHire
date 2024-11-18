package crawler

type Job struct {
	Title       string
	Company     string
	Location    string
	PostedDate  string
	Description string
	URL         string
}

type Crawler interface {
	Crawl() ([]Job, error)
}
