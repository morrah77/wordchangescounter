package straightforward

import (
	"errors"
	"io"
	"strings"

	"github.com/morrah77/wordchangescounter/matcher/core"
)

type Matcher struct {
	*core.Matcher
}

func NewMatcher(reader io.Reader) (m *Matcher, err error) {
	mt, err := core.NewMatcher(reader)
	if err != nil {
		return nil, err
	}
	return &Matcher{
		mt,
	}, nil
}

func (m *Matcher) change(word []byte) [][]byte {
	var res [][]byte
	res = append(res, m.changeRemoveChar(word)...)
	res = append(res, m.changeSubstituteChar(word)...)
	res = append(res, m.changeAddChar(word)...)
	return res
}

func (m *Matcher) changeRemoveChar(word []byte) [][]byte {
	var (
		res [][]byte
		cw  []byte
		l   int
		r   int
	)
	for i, _ := range word {
		l = r
		cw = append(cw, word[0:i]...)
		cw = append(cw, word[i+1:]...)
		r = len(cw)
		res = append(res, cw[l:])
	}
	return res
}

func (m *Matcher) changeAddChar(word []byte) [][]byte {
	var (
		res [][]byte
		cw  []byte
		p   int
		l   int
		r   int
	)
	p = len(word)
	for i := 0; i <= p; i++ {
		for _, c := range m.Alphabet {
			l = r
			cw = append(cw, word[0:i]...)
			cw = append(cw, c)
			cw = append(cw, word[i:]...)
			r = len(cw)
			res = append(res, cw[l:])
		}
	}
	return res
}

func (m *Matcher) changeSubstituteChar(word []byte) [][]byte {
	var (
		res [][]byte
		cw  []byte
		l   int
		r   int
	)
	for i, _ := range word {
		for _, c := range m.Alphabet {
			l = r
			cw = append(cw, word[0:i]...)
			cw = append(cw, c)
			cw = append(cw, word[i+1:]...)
			r = len(cw)
			res = append(res, cw[l:])
		}
	}
	return res
}

func (m *Matcher) Count(word []byte) (count int, err error) {
	var (
		changed [][]byte
	)
	changed = [][]byte{
		[]byte(strings.ToUpper(string(word))),
	}
LOOP:
	for {
		for _, c := range changed {
			if m.Matcher.Match(c) {
				break LOOP
			}
		}
		count++
		if count > m.MaxSeqLength+m.MaxExtraChanges {
			return 0, errors.New(`No appropriate word in vocabulary`)
		}
		prevChanged := changed
		changed = make([][]byte, 0)
		for _, cw := range prevChanged {
			changed = append(changed, m.change(cw)...)
		}
	}
	return count, nil
}
