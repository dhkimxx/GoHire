package crawler_test

import (
	"go-hire/crawler"

	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCrawlJobs(t *testing.T) {
	var crawler crawler.Crawler = crawler.NewSaraminCrawler()

	jobs, err := crawler.Crawl()
	if err != nil {
		t.Fatalf("Error crawling: %v\n", err)
	}
	t.Logf("Found %d jobs:\n", len(jobs))
	if len(jobs) == 0 {
		t.Fatal("failed to jobs")
	}
	for _, job := range jobs {
		t.Logf("Title: %s\nCompany: %s\nLocation: %s\nURL: %s\n\n", job.Title, job.Company, job.Location, job.URL)
		assert.NotEmpty(t, job.Title)
		assert.NotEmpty(t, job.Company)
		assert.NotEmpty(t, job.Location)
		assert.NotEmpty(t, job.PostedDate)
		assert.NotEmpty(t, job.Description)
		assert.NotEmpty(t, job.URL)
	}
}
