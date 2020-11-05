package controllers

import (
	"io/ioutil"
	"os"
	"path"
	"strings"

	"github.com/lcycug/go-scanner/utils"
)

const (
	test        = "Test"
	_test       = "_Test"
	classSuffix = ".cls"
)

func RenameTestFile(filePath string) error {
	rootPath, fileName := utils.GetRootPathAndFileName(filePath)
	if strings.Contains(strings.ToLower(fileName), test) && !strings.Contains(strings.ToLower(fileName),
		strings.ToLower(_test)) {
		var rest int
		if len(fileName)+1 > 40 {
			rest = 40 - len(test)
		} else {
			rest = len(fileName) - len(test)
		}
		newFileName := path.Join(rootPath, fileName[0:rest]+_test+classSuffix)
		err := updateTestClassName(filePath, newFileName)
		if err != nil {
			return err
		}
		return os.Rename(filePath, newFileName)
	}
	return nil
}

func updateTestClassName(filePath, newName string) error {
	lines, err := utils.File2Lines(filePath)
	if err != nil {
		return err
	}
	_, fileName := utils.GetRootPathAndFileName(filePath)
	for i, l := range lines {
		if strings.Contains(l, fileName) && strings.Contains(l, "class") {
			ss := strings.Split(l, fileName)
			lines[i] = strings.Join(ss, newName)
			break
		}
	}
	return ioutil.WriteFile(filePath, []byte(strings.Join(lines, "\n")), 0644)
}
