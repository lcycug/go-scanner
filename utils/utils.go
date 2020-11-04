package utils

import (
	"bufio"
	"io"
	"io/ioutil"
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

func InsertSharing(path string) error {
	lines, err := File2Lines(path)
	LogFatal(models.NewError(err, "Failed to read lines from file "+path))

	fileParts := strings.Split(path, "/")
	filename := strings.Split(fileParts[len(fileParts)-1], ".")[0]
	for i, l := range lines {
		if strings.Contains(l, filename) && strings.Contains(l,
			"class") && !strings.Contains(l, "with sharing") && !strings.
			Contains(l, "without sharing") {
			ss := strings.Split(l, "class")
			lines[i] = strings.Join(ss, "with sharing class")
			break
		}
	}
	return ioutil.WriteFile(path, []byte(strings.Join(lines, "\n")), 0644)
}
