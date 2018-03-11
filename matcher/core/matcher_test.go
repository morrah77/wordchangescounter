package core

import (
	"os"
	"testing"
)

func TestMatcher_match(t *testing.T) {
	var (
		m        *Matcher
		word     []byte
		result   bool
		expected bool
	)
	expected = true
	m = createPrefilledMatcher(t)
	word = []byte(`AAH`)
	result = m.Match(word)
	if result != expected {
		t.Error(`Match() method should Match existing word`)
	}
}

func TestMatcher_match_1(t *testing.T) {
	var (
		m        *Matcher
		word     []byte
		result   bool
		expected bool
	)
	expected = false
	m = createPrefilledMatcher(t)
	word = []byte(`AAHZ`)
	result = m.Match(word)
	if result != expected {
		t.Error(`Match() method should not Match inexisting word`)
	}
}

func createPrefilledMatcher(t *testing.T) *Matcher {
	f, err := os.Open(`../testdata/vocabulary.txt`)
	if err != nil {
		t.Error(err)
	}
	m, err := NewMatcher(f)
	if err != nil {
		t.Error(err)
	}
	if m.MaxExtraChanges != 2 {
		t.Fatal(`Wrong maxExtraChanges`)
	}
	if m.MaxSeqLength != 6 {
		t.Fatal(`Wrong maxSeqLength`)
	}
	if len(m.Alphabet) != 26 {
		t.Fatal(`Wrong alphabet length`)
	}
	if len(m.Vocabulary) != 5 {
		t.Fatal(`Wrong vocabulary length`)
	}
	if len(m.VocabularyByLengths) != 5 {
		t.Fatal(`Wrong vocabulary length`)
	}
	if _, ok := m.VocabularyByLengths[1]; ok {
		t.Fatal(`Unexpected VocabularyByLengths element exists!`)
	}
	if v, ok := m.VocabularyByLengths[2]; ok {
		if len(v) != 1 {
			t.Fatal(`Wrong VocabularyByLengths element length`)
		}
	} else {
		t.Fatal(`Expected VocabularyByLengths element does not exist!`)
	}
	if v, ok := m.VocabularyByLengths[4]; ok {
		if len(v) != 1 {
			t.Fatal(`Wrong VocabularyByLengths element length`)
		}
	} else {
		t.Fatal(`Expected VocabularyByLengths element does not exist!`)
	}
	return m
}
