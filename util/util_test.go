package util

import (
	"strings"
	"testing"

	. "gopkg.in/check.v1"
)

func Test(t *testing.T) { TestingT(t) }

type NgramsSuite struct{}

var _ = Suite(&NgramsSuite{})

func (s *NgramsSuite) Test2gram(c *C) {
	sent := "gophers love to go to denver"
	res := Ngrams(strings.Split(sent, " "), 2)
	c.Assert(res, DeepEquals, [][]string{
		{"gophers", "love"},
		{"love", "to"},
		{"to", "go"},
		{"go", "to"},
		{"to", "denver"},
	})
}

func (s *NgramsSuite) Test3gram(c *C) {
	sent := "gophers love to go to denver"
	res := Ngrams(strings.Split(sent, " "), 3)
	c.Assert(res, DeepEquals, [][]string{
		{"gophers", "love", "to"},
		{"love", "to", "go"},
		{"to", "go", "to"},
		{"go", "to", "denver"},
	})
}
