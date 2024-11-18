package crawler

import (
	"fmt"
	"log"

	"github.com/gocolly/colly/v2"
)

const saraminDomain = "https://www.saramin.co.kr"

type SaraminCrawler struct{}

func NewSaraminCrawler() Crawler {
	return &SaraminCrawler{}
}

func (s *SaraminCrawler) Crawl() ([]Job, error) {
	var jobs []Job

	c := colly.NewCollector()

	c.OnHTML("div.content", func(e *colly.HTMLElement) {
		// log.Println("Found div:", e.Text)

		e.ForEach("div.area_job", func(_ int, e *colly.HTMLElement) {

			// e.ForEach("h2.job_tit", func(_ int, span *colly.HTMLElement) {
			// 	log.Println("title: ", strings.TrimSpace(span.Text))
			// })

			job := Job{
				Title:       e.ChildText("h2.job_tit"),
				Company:     "",
				Location:    "",
				PostedDate:  "",
				Description: "",
				URL:         fmt.Sprintf("%s%s", saraminDomain, e.ChildAttr("a.data_layer", "href")),
			}
			jobs = append(jobs, job)
			log.Println(job)
		})

	})

	err := c.Visit(fmt.Sprintf("%s/zf_user/search?search_area=main&search_done=y&search_optional_item=n&searchType=search&searchword=개발", saraminDomain))
	if err != nil {
		log.Println("Error visiting Saramin:", err)
		return nil, err
	}

	return jobs, nil
}
