package core

import (
	"bufio"
	"io"
)

var ALPHABET = []byte(`ABCDEFGHIJKLMNOPQRSTUVWXYZ`)

const MAX_EXTRA_CHANGES = 2

type Matcher struct {
	Alphabet            []byte
	Vocabulary          map[string]struct{}
	VocabularyByLengths map[int]map[string]struct{}
	MaxSeqLength        int
	MaxExtraChanges     int
}

func NewMatcher(reader io.Reader) (m *Matcher, err error) {
	var (
		scanner *bufio.Scanner
		word    []byte
	)
	m = &Matcher{
		Alphabet:            ALPHABET,
		MaxExtraChanges:     MAX_EXTRA_CHANGES,
		Vocabulary:          make(map[string]struct{}, 0),
		VocabularyByLengths: make(map[int]map[string]struct{}, 0),
	}
	scanner = bufio.NewScanner(reader)
	for scanner.Scan() {
		word = scanner.Bytes()
		vKey := string(word)
		wLen := len(word)
		if wLen == 0 {
			continue
		}
		m.Vocabulary[vKey] = struct{}{}
		if wLen > m.MaxSeqLength {
			m.MaxSeqLength = wLen
		}
		if _, ok := m.VocabularyByLengths[wLen]; !ok {
			m.VocabularyByLengths[wLen] = make(map[string]struct{}, 0)
		}
		m.VocabularyByLengths[wLen][vKey] = struct{}{}
	}
	if scanner.Err() != nil {
		return nil, err
	}
	return m, nil
}

func (m *Matcher) Match(word []byte) bool {
	_, ok := m.Vocabulary[string(word)]
	return ok
}
