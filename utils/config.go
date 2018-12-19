package utils

import (
	"fmt"
	"os"

	"github.com/jinzhu/configor"
)

// Config is a struct loaded project level config
var Config = loadConfig()

type config struct {
	Port string `default:"8080"`

	GoogleCloud struct {
		BucketName      string
		CredentialsPath string
	}
}

func loadConfig() (conf config) {
	dir, _ := os.Getwd()

	err := configor.Load(&conf, dir+"/conf/config.yml")
	if err != nil {
		fmt.Println(err)
	}

	os.Setenv("GOOGLE_APPLICATION_CREDENTIALS", dir+conf.GoogleCloud.CredentialsPath)
	return conf
}
