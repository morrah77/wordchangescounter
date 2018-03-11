package levenshtaindistance

import (
	"io"

	"strings"

	"github.com/morrah77/wordchangescounter/matcher/core"
)

//TODO(h.lazar) consider to use FTS-like matcher instead of map, even structured
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

func (m *Matcher) Count(word []byte) (count int, err error) {
	normalizedWord := []byte(strings.ToUpper(string(word)))
	wl := len(normalizedWord)
	if v, ok := m.VocabularyByLengths[wl]; ok {
		if _, ok = v[string(word)]; ok {
			return 0, nil
		}
	}
	//if m.Match(normalizedWord) {
	//	return 0, nil
	//}
	count = wl
	if m.MaxSeqLength > count {
		count = m.MaxSeqLength
	}
	//	wli := wl
	//	wlii := 0
	//	even := false
	//	ampl := wl
	//	if m.MaxSeqLength > wl {
	//		ampl = m.MaxSeqLength - wl
	//		if wl > ampl {
	//			ampl = wl
	//		}
	//	}
	//LOOP:
	//	for {
	//		if v, ok := m.VocabularyByLengths[wli]; ok {
	//			for k, _ := range v {
	//				l := m.levenshteinDistance(normalizedWord, []byte(k))
	//				if l < count {
	//					count = l
	//				}
	//				if l == 0 {
	//					break LOOP
	//				}
	//			}
	//		}
	//		d := wli - wl
	//		if d < 0 {
	//			d = -d
	//		}
	//		if count < d {
	//			break LOOP
	//		}
	//		if even {
	//			wli = wl + wlii
	//		} else {
	//			wlii++
	//			wli = wl - wlii
	//		}
	//		if wlii > ampl {
	//			break LOOP
	//		}
	//		even = !even
	//	}
LOOP:
	for l, v := range m.VocabularyByLengths {
		d := l - wl
		if d < 0 {
			d = -d
		}
		if count > d {
			for k, _ := range v {
				l := m.levenshteinDistance(normalizedWord, []byte(k))
				if l < count {
					count = l
				}
				if l == 0 {
					break LOOP
				}
			}
		}
	}
	return count, nil
}

func (m *Matcher) levenshteinDistance(word []byte, vocabularyItem []byte) (d int) {
	var (
		ls   int
		lt   int
		cr   []int
		pr   []int
		cost int
	)
	ls = len(word)
	lt = len(vocabularyItem)
	if ls == 0 {
		return lt
	}
	if lt == 0 {
		return ls
	}
	pr = make([]int, ls+1)
	for i := 0; i <= lt; i++ {
		cr = make([]int, ls+1)
		if i == 0 {
			for j := 0; j <= ls; j++ {
				cr[j] = j
			}
			copy(pr, cr)
			continue
		}
		for j := 0; j <= ls; j++ {
			if j == 0 {
				cr[j] = i
				continue
			}
			cost = 0
			if word[j-1] != vocabularyItem[i-1] {
				cost = 1
			}
			// Let's don't use math.Min(float64, float64) to work with such a few int values
			cr[j] = pr[j] + 1
			distPrev2 := cr[j-1] + 1
			if distPrev2 < cr[j] {
				cr[j] = distPrev2
			}
			distPrev3 := pr[j-1] + cost
			if distPrev3 < cr[j] {
				cr[j] = distPrev3
			}
		}
		copy(pr, cr)
	}
	d = cr[ls]
	return d
}
