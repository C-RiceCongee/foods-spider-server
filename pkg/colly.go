package pkg

import (
	"github.com/gocolly/colly/v2"
	"github.com/gocolly/colly/v2/extensions"
)

type CLDColly struct {
	url string
	C   *colly.Collector
}

func (c *CLDColly) Do() error {
	return c.C.Visit(c.url)
}

func NewCLDColly(url string) *CLDColly {
	c := colly.NewCollector()
	extensions.Referer(c)
	extensions.RandomUserAgent(c)
	return &CLDColly{
		url: url,
		C:   c,
	}
}
