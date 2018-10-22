package wordsegmentation

import (
	"testing"

	c "github.com/bradbeam/wordsegmentation/corpus"
	"github.com/stretchr/testify/assert"
)

func TestSegment(t *testing.T) {
	expected := []string{"what", "is", "the", "weather", "like", "today"}
	englishCorpus := c.NewEnglishCorpus()

	assert.Equal(t, Segment(englishCorpus, "WhatIsTheWeatherliketoday? "), expected)
}
