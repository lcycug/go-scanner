package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strings"

	"github.com/joho/godotenv"
	"github.com/lcycug/go-scanner/models"
	"github.com/lcycug/go-scanner/utils"
)

func init() {
	err := godotenv.Load()
	utils.LogFatal(models.NewError(err, "Error loading .env file"))
}

func main() {
	rootPath, ok := os.LookupEnv("ROOT_PATH")
	utils.LogFatal(models.NewOk(ok, "Failed to find ROOT_PATH in .env file."))

	fileInfos, err := ioutil.ReadDir(rootPath)
	utils.LogFatal(models.NewError(err, "Failed to read directory:"))

	for _, fi := range fileInfos {
		if strings.Contains(fi.Name(), ".cls-meta.xml") {
			continue
		}
		filepath := path.Join(rootPath, fi.Name())
		data, err := ioutil.ReadFile(filepath)
		utils.LogFatal(models.NewError(err, "Failed to read a file."))

		dataString := string(data)
		if !strings.Contains(dataString, "with sharing") && !strings.Contains(dataString, "without sharing") {
			fmt.Println("This class has no explicit sharing setting", fi.Name())
		}

		err = utils.InsertSharing(filepath)
		utils.LogFatal(models.NewError(err,
			"Failed to insert sharing setting to a file "+filepath))
		os.Exit(0)
	}
}
