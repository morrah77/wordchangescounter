package straightforward

import (
	"errors"
	"os"
	"testing"
)

func BenchmarkMatcher_Count(b *testing.B) {
	var (
		m    *Matcher
		word []byte
		err  error
	)
	m, err = createPrefilledMatcherForBenchmark()
	if err != nil {
		b.Error(err)
	}
	word = []byte(`A`)
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
	m, err = createPrefilledMatcherForBenchmark()
	if err != nil {
		b.Error(err)
	}
	word = []byte(`AAHAB`)
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
	m, err = createPrefilledMatcherForBenchmark()
	if err != nil {
		b.Error(err)
	}
	word = []byte(`AAHINGSS`)
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
