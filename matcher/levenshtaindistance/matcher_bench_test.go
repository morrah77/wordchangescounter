package levenshtaindistance

import (
	"errors"
	"os"
	"testing"
)

func BenchmarkNewMatcher(b *testing.B) {
	f, err := os.Open(`../../resources/vocabulary.txt`)
	if err != nil {
		b.Error(err)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = NewMatcher(f)
	}
}

func BenchmarkMatcher_levenshteinDistance(b *testing.B) {
	var (
		m              *Matcher
		word           []byte
		vocabularyItem []byte
		err            error
	)
	word = []byte(`AAHED`)
	vocabularyItem = []byte(`AAHING`)
	m, err = createPrefilledMatcherForBenchmark()
	if err != nil {
		b.Error(err)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = m.levenshteinDistance(word, vocabularyItem)
	}
}

func BenchmarkMatcher_levenshteinDistance_1(b *testing.B) {
	var (
		m              *Matcher
		word           []byte
		vocabularyItem []byte
		err            error
	)
	word = []byte(`AAHED`)
	vocabularyItem = []byte(`AAHING`)
	m, err = createRealMatcherForBenchmark()
	if err != nil {
		b.Error(err)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = m.levenshteinDistance(word, vocabularyItem)
	}
}

func BenchmarkMatcher_Count(b *testing.B) {
	var (
		m    *Matcher
		word []byte
		err  error
	)
	word = []byte(`AAHED`)
	m, err = createPrefilledMatcherForBenchmark()
	if err != nil {
		b.Error(err)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = m.Count(word)
	}
}

func BenchmarkMatcher_Count_1(b *testing.B) {
	var (
		m    *Matcher
		word []byte
		err  error
	)
	word = []byte(`AAHED`)
	m, err = createRealMatcherForBenchmark()
	if err != nil {
		b.Error(err)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = m.Count(word)
	}
}

func BenchmarkMatcher_Count_2(b *testing.B) {
	var (
		m    *Matcher
		word []byte
		err  error
	)
	word = []byte(`AAHDE`)
	m, err = createRealMatcherForBenchmark()
	if err != nil {
		b.Error(err)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = m.Count(word)
	}
}

func createPrefilledMatcherForBenchmark() (*Matcher, error) {
	f, err := os.Open(`../testdata/vocabulary.txt`)
	if err != nil {
		return nil, err
	}
	m, err := NewMatcher(f)
	if err != nil {
		return nil, err
	}
	if m.MaxExtraChanges != 2 {
		return nil, errors.New(`Wrong maxExtraChanges`)
	}
	if m.MaxSeqLength != 6 {
		return nil, errors.New(`Wrong maxSeqLength`)
	}
	if len(m.Alphabet) != 26 {
		return nil, errors.New(`Wrong alphabet length`)
	}
	if len(m.Vocabulary) != 5 {
		return nil, errors.New(`Wrong vocabulary length`)
	}
	return m, nil
}

func createRealMatcherForBenchmark() (*Matcher, error) {
	f, err := os.Open(`../../resources/vocabulary.txt`)
	if err != nil {
		return nil, err
	}
	m, err := NewMatcher(f)
	if err != nil {
		return nil, err
	}
	if m.MaxExtraChanges != 2 {
		return nil, errors.New(`Wrong maxExtraChanges`)
	}
	if m.MaxSeqLength != 15 {
		return nil, errors.New(`Wrong maxSeqLength`)
	}
	if len(m.Alphabet) != 26 {
		return nil, errors.New(`Wrong alphabet length`)
	}
	if len(m.Vocabulary) != 178691 {
		return nil, errors.New(`Wrong vocabulary length`)
	}
	return m, nil
}
