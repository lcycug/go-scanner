package controllers

import (
	"io/ioutil"
	"strings"

	"github.com/lcycug/go-scanner/models"
	"github.com/lcycug/go-scanner/utils"
)

func InsertSharing(path string) error {
	lines, err := utils.File2Lines(path)
	utils.LogFatal(models.NewError(err, "Failed to read lines from file "+path))

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
