package pkg

import (
	"github.com/gocolly/colly/v2"
	"sync"
)

type CLDColly struct {
	url  string
	C    *colly.Collector
	HTML chan *colly.HTMLElement
	Lock sync.WaitGroup
}

func (c *CLDColly) Do() error {
	return c.C.Visit(c.url)
}
func NewCLDColly(url string) *CLDColly {
	return &CLDColly{
		url: url,
		C:   colly.NewCollector(),
	}
}
