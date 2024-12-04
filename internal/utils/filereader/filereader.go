package filereader

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

type FileReader struct {
	Path        string
	line        int
	hasMore     bool
	readFile    *os.File
	fileScanner *bufio.Scanner
}

func New(path string) FileReader {
	return FileReader{
		Path:        path,
		line:        -1,
		hasMore:     true,
		readFile:    nil,
		fileScanner: nil,
	}
}

func NewFromDayInput(day int, input int) FileReader {
	wd, err := os.Getwd()
	if err != nil {
		log.Fatalln(err)
	}
	filepath := filepath.Join(wd, "puzzles", fmt.Sprintf("day%d", day), fmt.Sprintf("input%d.txt", input))
	return New(filepath)
}

func NewFromDayExample(day int, input int) FileReader {
	wd, err := os.Getwd()
	if err != nil {
		log.Fatalln(err)
	}
	filepath := filepath.Join(wd, "puzzles", fmt.Sprintf("day%d", day), fmt.Sprintf("example%d.txt", input))
	return New(filepath)
}

func (fr *FileReader) ReadLines() ([]string, error) {
	readFile, err := os.Open(fr.Path)

	if err != nil {
		return []string{}, err
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	var fileLines []string

	for fileScanner.Scan() {
		fileLines = append(fileLines, fileScanner.Text())
	}

	readFile.Close()

	return fileLines, nil
}

func (fr *FileReader) Read() (string, bool, error) {
	if !fr.hasMore {
		return "", fr.hasMore, errors.New("End of the file. Reset the reader")
	}

	if fr.readFile == nil {
		var err error
		fr.readFile, err = os.Open(fr.Path)

		if err != nil {
			log.Fatalln(err)
		}

		fr.fileScanner = bufio.NewScanner(fr.readFile)
		fr.fileScanner.Split(bufio.ScanLines)
		if fr.fileScanner.Scan() {
			fr.line++
		} else {
			fr.hasMore = false
			return "", false, nil
		}

	}

	line := fr.fileScanner.Text()

	if fr.fileScanner.Scan() {
		fr.line++
	} else {
		fr.hasMore = false
	}

	return line, fr.hasMore, nil
}

func (fr *FileReader) HasMore() (bool) {
	return fr.hasMore
}

func (fr *FileReader) Close() {
	if fr.readFile == nil {
		return
	}
	fr.readFile.Close()
}

func (fr *FileReader) Reset() {
	if fr.readFile == nil {
		return
	}

	fr.fileScanner = bufio.NewScanner(fr.readFile)
	fr.fileScanner.Split(bufio.ScanLines)
}
