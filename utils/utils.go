package utils

import (
	"os/user"

	log "github.com/Sirupsen/logrus"
)

func HomeDir() string {
	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	return usr.HomeDir
}

func CheckContentType(contentTypes []string, target string) bool {
	for _, item := range contentTypes {
		if item == target {
			return true
		}
	}
	return false
}
