package levenshtaindistance

import (
	"bufio"
	"fmt"
	"os"
	"testing"
)

func TestNewMatcher(t *testing.T) {
	var m *Matcher
	m = createPrefilledMatcher(t)
	if len(m.Vocabulary) != 5 {
		t.Fatal(`Wrong vocabulary length`, 5, `expected but`, len(m.Vocabulary), `returned`)
	}
}

func TestMatcher_levenshteinDistance(t *testing.T) {
	var (
		m              *Matcher
		word           []byte
		vocabularyItem []byte
		result         int
		expected       int
	)
	word = []byte(`AAHED`)
	vocabularyItem = []byte(`AAHING`)
	expected = 3
	m = createPrefilledMatcher(t)
	result = m.levenshteinDistance(word, vocabularyItem)
	if result != expected {
		t.Error(`levenshteinDistance() method should calculate levenshtein distance: word`, word, `and vocabularyItem`, vocabularyItem, `given,`, expected, `expected but`, result, `returned`)
	}
}

func TestMatcher_levenshteinDistance_1(t *testing.T) {
	var (
		m              *Matcher
		word           []byte
		vocabularyItem []byte
		result         int
		expected       int
	)
	word = []byte(``)
	vocabularyItem = []byte(`AAHING`)
	expected = 6
	m = createPrefilledMatcher(t)
	result = m.levenshteinDistance(word, vocabularyItem)
	if result != expected {
		t.Error(`levenshteinDistance() method should calculate levenshtein distance: word`, word, `and vocabularyItem`, vocabularyItem, `given,`, expected, `expected but`, result, `returned`)
	}
}

func TestMatcher_levenshteinDistance_2(t *testing.T) {
	var (
		m              *Matcher
		word           []byte
		vocabularyItem []byte
		result         int
		expected       int
	)
	word = []byte(`AAHED`)
	vocabularyItem = []byte(``)
	expected = 5
	m = createPrefilledMatcher(t)
	result = m.levenshteinDistance(word, vocabularyItem)
	if result != expected {
		t.Error(`levenshteinDistance() method should calculate levenshtein distance: word`, word, `and vocabularyItem`, vocabularyItem, `given,`, expected, `expected but`, result, `returned`)
	}
}

func TestMatcher_levenshteinDistance_3(t *testing.T) {
	var (
		m              *Matcher
		word           []byte
		vocabularyItem []byte
		result         int
		expected       int
	)
	word = []byte(`B`)
	vocabularyItem = []byte(`AAHING`)
	expected = 6
	m = createPrefilledMatcher(t)
	result = m.levenshteinDistance(word, vocabularyItem)
	if result != expected {
		t.Error(`levenshteinDistance() method should calculate levenshtein distance: word`, word, `and vocabularyItem`, vocabularyItem, `given,`, expected, `expected but`, result, `returned`)
	}
}

func TestMatcher_Count(t *testing.T) {
	var (
		m        *Matcher
		word     []byte
		result   int
		err      error
		expected int
	)
	word = []byte(`AAHED`)
	expected = 0
	m = createPrefilledMatcher(t)
	result, err = m.Count(word)
	if err != nil {
		fmt.Printf("word: %#s(%v), count: %#v, error: %#s\n\n", word, len(word), result, err)
		t.Error(`Count() method should not return error for changeable word`, word, `Error:`, err)
	}
	if result != expected {
		t.Error(`Count() method should count minimum changes for word:`, word, `given,`, expected, `expected but`, result, `returned`)
	}
}

func TestMatcher_Count_1(t *testing.T) {
	var (
		m        *Matcher
		word     []byte
		result   int
		err      error
		expected int
	)
	word = []byte(`AABIH`)
	expected = 2
	m = createPrefilledMatcher(t)
	result, err = m.Count(word)
	if err != nil {
		fmt.Printf("word: %#s(%v), count: %#v, error: %#s\n\n", word, len(word), result, err)
		t.Error(`Count() method should not return error for changeable word`, word, `Error:`, err)
	}
	if result != expected {
		t.Error(`Count() method should count minimum changes for word:`, fmt.Sprintf(`%#s`, word), `given,`, expected, `expected but`, result, `returned`)
	}
}

func TestMatcher_Count_2(t *testing.T) {
	var (
		m        *Matcher
		c        int
		changes  int
		expected int
		words    [][]byte
	)
	expected = 8
	m = createRealMatcher(t)
	inputFile, err := os.Open(`../../docs/example_input`)
	if err != nil {
		t.Fatal(err)
	}
	s := bufio.NewScanner(inputFile)
	s.Split(bufio.ScanWords)
	for s.Scan() {
		word := s.Bytes()

		if len(word) <= 0 {
			continue
		}
		words = append(words, word)
		c, err = m.Count(word)
		if err != nil {
			fmt.Printf("word: %#s(%v), count: %#v, error: %#s\n\n", word, len(word), c, err)
			t.Error(`Count() method should not return error for changeable word %#s`, word, `Error:`, err)
			continue
		}
		changes += c
	}
	if changes != expected {
		ws := fmt.Sprintf(`%#s`, words)
		t.Error(`Count() method should count minimum changes for word:`, ws, `given,`, expected, `expected but`, changes, `returned`)
	}
}

func TestMatcher_Count_3(t *testing.T) {
	var (
		m        *Matcher
		c        int
		changes  int
		expected int
		words    [][]byte
	)
	expected = 8
	m = createRealMatcher(t)
	inputFile, err := os.Open(`../testdata/example_input_without_newline`)
	if err != nil {
		t.Fatal(err)
	}
	s := bufio.NewScanner(inputFile)
	s.Split(bufio.ScanWords)
	for s.Scan() {
		word := s.Bytes()

		if len(word) <= 0 {
			continue
		}
		words = append(words, word)
		c, err = m.Count(word)
		if err != nil {
			fmt.Printf("word: %#s(%v), count: %#v, error: %#s\n\n", word, len(word), c, err)
			t.Error(`Count() method should not return error for changeable word %#s`, word, `Error:`, err)
			continue
		}
		changes += c
	}
	if changes != expected {
		ws := fmt.Sprintf(`%#s`, words)
		t.Error(`Count() method should count minimum changes for word:`, ws, `given,`, expected, `expected but`, changes, `returned`)
	}
}

func TestMatcher_Count_187(t *testing.T) {
	var (
		m        *Matcher
		c        int
		changes  int
		expected int
		words    [][]byte
	)
	expected = 187
	m = createRealMatcher(t)
	inputFile, err := os.Open(`../../docs/187`)
	if err != nil {
		t.Fatal(err)
	}
	s := bufio.NewScanner(inputFile)
	s.Split(bufio.ScanWords)
	for s.Scan() {
		word := s.Bytes()

		if len(word) <= 0 {
			continue
		}
		words = append(words, word)
		c, err = m.Count(word)
		if err != nil {
			fmt.Printf("word: %#s(%v), count: %#v, error: %#s\n\n", word, len(word), c, err)
			t.Error(`Count() method should not return error for changeable word %#s`, word, `Error:`, err)
			continue
		}
		changes += c
	}
	if changes != expected {
		ws := fmt.Sprintf(`%#s`, words)
		t.Error(`Count() method should count minimum changes for word:`, ws, `given,`, expected, `expected but`, changes, `returned`)
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
	return m
}

func createRealMatcher(t *testing.T) *Matcher {
	f, err := os.Open(`../../resources/vocabulary.txt`)
	if err != nil {
		t.Fatal(err)
	}
	m, err := NewMatcher(f)
	if err != nil {
		t.Fatal(err)
	}
	if m.MaxExtraChanges != 2 {
		t.Fatal(`Wrong maxExtraChanges`, 2, `expected but`, m.MaxExtraChanges, `returned`)
	}
	if m.MaxSeqLength != 15 {
		t.Fatal(`Wrong maxSeqLength`, 15, `expected but`, m.MaxSeqLength, `returned`)
	}
	if len(m.Alphabet) != 26 {
		t.Fatal(`Wrong alphabet length`, 26, `expected but`, len(m.Alphabet), `returned`)
	}
	if len(m.Vocabulary) != 178691 {
		t.Fatal(`Wrong vocabulary length`, 178691, `expected but`, len(m.Vocabulary), `returned`)
	}
	return m
}
