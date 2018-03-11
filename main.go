package main

import (
	"bufio"
	"flag"
	"io"
	"log"
	"os"

	matcher "github.com/morrah77/wordchangescounter/matcher/levenshtaindistance"
)

type Conf struct {
	sourcePath string
}

const vocabularySourcePath = `./resources/vocabulary.txt`

const reportPrefix = `Word changes counter `

var conf *Conf

type WordChangesCounter interface {
	Count([]byte) (int, error)
}

func init() {
	conf = &Conf{}
	flag.StringVar(&conf.sourcePath, `source-path`, `./187`, `Pass source file path`)
	flag.Parse()
}

func main() {
	var (
		input          io.ReadCloser
		changesCounter WordChangesCounter
		logger         *log.Logger
		result         int
	)
	logger = log.New(os.Stdout, reportPrefix, log.Flags())
	defer func() {
		if r := recover(); r != nil {
			logger.Fatal(r)
		}
	}()
	changesCounter = initChangesCounterFromFile()
	input = initInputFromFile()
	result = calculateResult(input, changesCounter, logger)
	println(result)
}

func calculateResult(input io.ReadCloser, changesCounter WordChangesCounter, logger *log.Logger) int {
	var (
		result int
		err    error
	)
	s := bufio.NewScanner(input)
	s.Split(bufio.ScanWords)
	for s.Scan() {
		word := s.Bytes()

		if len(word) <= 0 {
			continue
		}
		if len(word) <= 0 {
			continue
		}
		if n, e := changesCounter.Count(word); e == nil {
			result += n
		} else {
			logger.Println(e)
		}
	}
	err = input.Close()
	if err != nil {
		logger.Println(err.Error())
	}
	return result
}

func initChangesCounterFromFile() WordChangesCounter {
	var (
		vocabularyFile io.ReadCloser
		err            error
		changesCounter WordChangesCounter
	)
	vocabularyFile, err = os.Open(vocabularySourcePath)
	failOnError(err)
	changesCounter, err = matcher.NewMatcher(vocabularyFile)
	failOnError(err)
	err = vocabularyFile.Close()
	failOnError(err)
	return changesCounter
}

func initInputFromFile() io.ReadCloser {
	var (
		inputFile io.ReadCloser
		err       error
	)
	inputFile, err = os.Open(conf.sourcePath)
	failOnError(err)
	return inputFile
}

func failOnError(err error) {
	if err != nil {
		panic(err)
	}
}
