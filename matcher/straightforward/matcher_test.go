package straightforward

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"testing"
)

func TestMatcher_changeRemoveChar(t *testing.T) {
	var (
		m        *Matcher
		word     []byte
		changed  [][]byte
		expected [][]byte
	)
	expected = [][]byte{
		[]byte(``),
	}
	m = createPrefilledMatcher(t)
	word = []byte(`A`)
	changed = m.changeRemoveChar(word)
	if len(changed) != len(expected) {
		fmt.Printf(`%#s\n%#s`, changed, expected)
		t.Fatal(`change() method should return set of new sequences with appropriate length. Expected`, len(expected), `but returned`, len(changed))
	}
	for i, c := range changed {
		if !bytes.Equal(c, expected[i]) {
			fmt.Printf(`%#s\n%#s`, changed, expected)
			t.Fatal(`changeRemoveChar() method should return appropriate set of new sequences`, i)
		}
	}
}

func TestMatcher_changeRemoveChar_1(t *testing.T) {
	var (
		m        *Matcher
		word     []byte
		changed  [][]byte
		expected [][]byte
	)
	expected = [][]byte{
		[]byte(`BC`),
		[]byte(`AC`),
		[]byte(`AB`),
	}
	m = createPrefilledMatcher(t)
	word = []byte(`ABC`)
	changed = m.changeRemoveChar(word)
	if len(changed) != len(expected) {
		fmt.Printf("%#s\n%#s\n", changed, expected)
		t.Fatal(`changeRemoveChar() method should return set of new sequences with appropriate length. Expected`, len(expected), `but returned`, len(changed))
	}
	for i, c := range changed {
		if !bytes.Equal(c, expected[i]) {
			fmt.Printf("%#s\n%#s\n", changed, expected)
			t.Fatal(`changeRemoveChar() method should return appropriate set of new sequences`, i)
		}
	}
}

func TestMatcher_changeAddChar(t *testing.T) {
	var (
		m        *Matcher
		word     []byte
		changed  [][]byte
		expected [][]byte
	)
	expected = [][]byte{
		[]byte(`AA`),
		[]byte(`BA`),
		[]byte(`CA`),
		[]byte(`DA`),
		[]byte(`EA`),
		[]byte(`AA`),
		[]byte(`AB`),
		[]byte(`AC`),
		[]byte(`AD`),
		[]byte(`AE`),
	}
	m = createPrefilledMatcher(t)
	m.Alphabet = []byte(`ABCDE`)
	word = []byte(`A`)
	changed = m.changeAddChar(word)
	if len(changed) != len(expected) {
		fmt.Printf("%#s\n%#s\n", changed, expected)
		t.Fatal(`changeAddChar() method should return set of new sequences with appropriate length. Expected`, len(expected), `but returned`, len(changed))
	}
	for i, c := range changed {
		if !bytes.Equal(c, expected[i]) {
			fmt.Printf("%#s\n%#s\n", changed, expected)
			t.Fatal(`changeAddChar() method should return appropriate set of new sequences`, i)
		}
	}
}

func TestMatcher_changeAddChar_1(t *testing.T) {
	var (
		m        *Matcher
		word     []byte
		changed  [][]byte
		expected [][]byte
	)
	expected = [][]byte{
		[]byte(`AABC`),
		[]byte(`BABC`),
		[]byte(`CABC`),
		[]byte(`DABC`),
		[]byte(`EABC`),
		[]byte(`AABC`),
		[]byte(`ABBC`),
		[]byte(`ACBC`),
		[]byte(`ADBC`),
		[]byte(`AEBC`),
		[]byte(`ABAC`),
		[]byte(`ABBC`),
		[]byte(`ABCC`),
		[]byte(`ABDC`),
		[]byte(`ABEC`),
		[]byte(`ABCA`),
		[]byte(`ABCB`),
		[]byte(`ABCC`),
		[]byte(`ABCD`),
		[]byte(`ABCE`),
	}
	m = createPrefilledMatcher(t)
	m.Alphabet = []byte(`ABCDE`)
	word = []byte(`ABC`)
	changed = m.changeAddChar(word)
	if len(changed) != len(expected) {
		fmt.Printf("%#s\n%#s\n", changed, expected)
		t.Fatal(`changeAddChar() method should return set of new sequences with appropriate length. Expected`, len(expected), `but returned`, len(changed))
	}
	for i, c := range changed {
		if !bytes.Equal(c, expected[i]) {
			fmt.Printf("%#s\n%#s\n", changed, expected)
			t.Fatal(`changeAddChar() method should return appropriate set of new sequences`, i)
		}
	}
}

func TestMatcher_changeSubstituteChar(t *testing.T) {
	var (
		m        *Matcher
		word     []byte
		changed  [][]byte
		expected [][]byte
	)
	expected = [][]byte{
		[]byte(`A`),
		[]byte(`B`),
		[]byte(`C`),
		[]byte(`D`),
		[]byte(`E`),
	}
	m = createPrefilledMatcher(t)
	m.Alphabet = []byte(`ABCDE`)
	word = []byte(`A`)
	changed = m.changeSubstituteChar(word)
	if len(changed) != len(expected) {
		fmt.Printf("%#s\n%#s\n", changed, expected)
		t.Fatal(`changeSubstituteChar() method should return set of new sequences with appropriate length. Expected`, len(expected), `but returned`, len(changed))
	}
	for i, c := range changed {
		if !bytes.Equal(c, expected[i]) {
			fmt.Printf("%#s\n%#s\n", changed, expected)
			t.Fatal(`changeSubstituteChar() method should return appropriate set of new sequences`, i)
		}
	}
}

func TestMatcher_changeSubstituteChar_1(t *testing.T) {
	var (
		m        *Matcher
		word     []byte
		changed  [][]byte
		expected [][]byte
	)
	expected = [][]byte{
		[]byte(`ABC`),
		[]byte(`BBC`),
		[]byte(`CBC`),
		[]byte(`DBC`),
		[]byte(`EBC`),
		[]byte(`AAC`),
		[]byte(`ABC`),
		[]byte(`ACC`),
		[]byte(`ADC`),
		[]byte(`AEC`),
		[]byte(`ABA`),
		[]byte(`ABB`),
		[]byte(`ABC`),
		[]byte(`ABD`),
		[]byte(`ABE`),
	}
	m = createPrefilledMatcher(t)
	m.Alphabet = []byte(`ABCDE`)
	word = []byte(`ABC`)
	changed = m.changeSubstituteChar(word)
	if len(changed) != len(expected) {
		fmt.Printf("%#s\n%#s\n", changed, expected)
		t.Fatal(`change() method should return set of new sequences with appropriate length. Expected`, len(expected), `but returned`, len(changed))
	}
	for i, c := range changed {
		if !bytes.Equal(c, expected[i]) {
			fmt.Printf("%#s\n%#s\n", changed, expected)
			t.Fatal(`change() method should return appropriate set of new sequences`, i)
		}
	}
}

func TestMatcher_change(t *testing.T) {
	var (
		m        *Matcher
		word     []byte
		changed  [][]byte
		expected [][]byte
	)
	expected = [][]byte{
		[]byte(``),
		[]byte(`A`),
		[]byte(`B`),
		[]byte(`C`),
		[]byte(`D`),
		[]byte(`E`),
		[]byte(`AA`),
		[]byte(`BA`),
		[]byte(`CA`),
		[]byte(`DA`),
		[]byte(`EA`),
		[]byte(`AA`),
		[]byte(`AB`),
		[]byte(`AC`),
		[]byte(`AD`),
		[]byte(`AE`),
	}
	m = createPrefilledMatcher(t)
	m.Alphabet = []byte(`ABCDE`)
	word = []byte(`A`)
	changed = m.change(word)
	if len(changed) != len(expected) {
		fmt.Printf("%#s\n%#s\n", changed, expected)
		t.Fatal(`change() method should return set of new sequences with appropriate length. Expected`, len(expected), `but returned`, len(changed))
	}
	for i, c := range changed {
		if !bytes.Equal(c, expected[i]) {
			fmt.Printf("%#s\n%#s\n", changed, expected)
			t.Fatal(`change() method should return appropriate set of new sequences`, i)
		}
	}
}

func TestMatcher_change_1(t *testing.T) {
	var (
		m        *Matcher
		word     []byte
		changed  [][]byte
		expected int
	)
	m = createPrefilledMatcher(t)
	word = []byte(`ABC`)
	expected = len(word) + (len(word)+1)*(len(m.Alphabet)) + len(word)*(len(m.Alphabet))
	changed = m.change(word)
	if len(changed) != expected {
		fmt.Printf("%#s\n", changed)
		t.Error(`change() method should return a set of appropriate length.`, expected, `expected but`, len(changed), `returned`)
	}
}

func TestMatcher_Count(t *testing.T) {
	var (
		m        *Matcher
		word     []byte
		changes  int
		expected int
		err      error
	)
	expected = 1
	m = createPrefilledMatcher(t)
	word = []byte(`A`)
	changes, err = m.Count(word)
	if err != nil {
		t.Error(`Count() method should not return error for changeable word`, err)
	}
	if changes != expected {
		t.Error(`Count() method should count minimum changes for word`)
	}
}

func TestMatcher_Count_1(t *testing.T) {
	var (
		m        *Matcher
		word     []byte
		changes  int
		expected int
		err      error
	)
	expected = 2
	m = createPrefilledMatcher(t)
	word = []byte(`AAHAB`)
	changes, err = m.Count(word)
	if err != nil {
		t.Error(`Count() method should not return error for changeable word`, err)
	}
	if changes != expected {
		t.Error(`Count() method should count minimum changes for word`)
	}
}

func TestMatcher_Count_2(t *testing.T) {
	var (
		m        *Matcher
		word     []byte
		changes  int
		expected int
		err      error
	)
	expected = 1
	m = createPrefilledMatcher(t)
	word = []byte(`AAHIN`)
	changes, err = m.Count(word)
	if err != nil {
		t.Error(`Count() method should not return error for changeable word`, err)
	}
	if changes != expected {
		t.Error(`Count() method should count minimum changes for word`)
	}
}

func TestMatcher_Count_3(t *testing.T) {
	var (
		m        *Matcher
		word     []byte
		changes  int
		expected int
		err      error
	)
	expected = 2
	m = createPrefilledMatcher(t)
	word = []byte(`AAHINGSS`)
	changes, err = m.Count(word)
	if err != nil {
		t.Error(`Count() method should not return error for changeable word`, err)
	}
	if changes != expected {
		t.Error(`Count() method should count minimum changes for word`)
	}
}

//Uncomment if you want to hangout your machine
//func TestMatcher_Count_4(t *testing.T) {
//	var (
//		m                    *Matcher
//		word                 []byte
//		changes              int
//		expected             int
//		err                  error
//		expectedErrorMessage string
//	)
//	expected = 0
//	expectedErrorMessage = `No appropriate word in vocabulary`
//	m = createPrefilledMatcher(t)
//	word = []byte(`AAHINGSSQ`)
//	changes, err = m.Count(word)
//	if err == nil {
//		t.Error(`Count() method should return error for word required too many`, err)
//	}
//	if err.Error() != expectedErrorMessage {
//		t.Error(`Count() method should return appropriate error message`, expectedErrorMessage, `expected but`, err.Error(), `returned`)
//	}
//	if changes != expected {
//		t.Error(`Count() method should count minimum changes for word`)
//	}
//}

func TestMatcher_Count_4(t *testing.T) {
	var (
		m        *Matcher
		words    [][]byte
		c        int
		changes  int
		expected int
		err      error
	)
	expected = 4
	m = createPrefilledMatcher(t)
	words = [][]byte{
		[]byte(`AAHINGSS`),
		[]byte(`abc`),
	}
	for _, word := range words {
		c, err = m.Count(word)
		if err != nil {
			t.Error(`Count() method should not return error for changeable word`, err)
			continue
		}
		changes += c
	}
	if changes != expected {
		t.Error(`Count() method should count minimum changes for word`)
	}
}

func TestMatcher_Count_6(t *testing.T) {
	var (
		m        *Matcher
		c        int
		changes  int
		expected int
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
		c, err = m.Count(word)
		if err != nil {
			fmt.Printf("word: %#s(%v), count: %#v, error: %#s\n\n", word, len(word), c, err)
			t.Error(`Count() method should not return error for changeable word %#s`, word, `Error:`, err)
			continue
		}
		changes += c
	}
	if changes != expected {
		t.Error(`Count() method should count minimum changes for word`)
	}
}

//Uncomment if you want to hangout your machine
//func TestMatcher_Count_187(t *testing.T) {
//	var (
//		m        *Matcher
//		c        int
//		changes  int
//		expected int
//	)
//	expected = 187
//	m = createRealMatcher(t)
//	inputFile, err := os.Open(`../../docs/187`)
//	if err != nil {
//		t.Fatal(err)
//	}
//	s := bufio.NewScanner(inputFile)
//	s.Split(bufio.ScanWords)
//	for s.Scan() {
//		word := s.Bytes()
//
//		if len(word) <= 0 {
//			continue
//		}
//		c, err = m.Count(word)
//		if err != nil {
//			fmt.Printf("word: %#s(%v), count: %#v, error: %#s\n\n", word, len(word), c, err)
//			t.Error(`Count() method should not return error for changeable word %#s`, word, `Error:`, err)
//			continue
//		}
//		changes += c
//	}
//	if changes != expected {
//		t.Error(`Count() method should count minimum changes for word`)
//	}
//}

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
