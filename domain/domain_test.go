package domain

import (
	"github.com/300brand/spider/page"
	"github.com/300brand/spider/samplesite"
	"launchpad.net/gocheck"
	"testing"
)

type DomainSuite struct{}

var _ = gocheck.Suite(new(DomainSuite))

func Test(t *testing.T) { gocheck.TestingT(t) }

func (s *DomainSuite) TestRobotsTxt(c *gocheck.C) {
	d := &Domain{
		URL: samplesite.URL,
	}
	tests := map[string]bool{
		"/": true,
	}

	for path, canDL := range tests {
		p := &page.Page{
			URL: samplesite.URL + path,
		}
		c.Assert(d.CanDownload(p), gocheck.Equals, canDL)
	}
}
