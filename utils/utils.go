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
