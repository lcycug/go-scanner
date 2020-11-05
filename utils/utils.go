package utils

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"github.com/lcycug/go-scanner/models"
)

func LogFatal(e models.Errors) {
	switch e.GetType() {
	case models.ERROR:
		if e.GetError() != nil {
			log.Fatal(e.GetMessage())
		}
	case models.OK:
		if !e.GetOk() {
			log.Fatal(e.GetMessage())
		}
	}
}

func Log(e models.Errors) {
	switch e.GetType() {
	case models.ERROR:
		if e.GetError() != nil {
			fmt.Println(e.GetMessage())
		}
	case models.OK:
		if !e.GetOk() {
			fmt.Println(e.GetMessage())
		}
	}
}

func linesFromReader(r io.Reader) ([]string, error) {
	var lines []string
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return lines, nil
}

func File2Lines(filepath string) ([]string, error) {
	f, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	return linesFromReader(f)
}

func GetRootPathAndFileName(filePath string) (string, string) {
	paths := strings.Split(filePath, "/")
	rootPath := strings.Join(paths[:len(paths)-1], "/")
	fileName := strings.Split(paths[len(paths)-1], ".")[0]
	return rootPath, fileName
}
