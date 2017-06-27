package token

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"time"

	log "github.com/Sirupsen/logrus"
	"github.com/wahyd4/say-it/utils"
)

const (
	TokenFile = "/.sayit"
	TokenURL  = "http://say.toozhao.com:8000/api/token"
)

type Token struct {
	Value     string `json:"Token"`
	ExpiresAt int64  `json:"ExpiresTime"`
}

func LoadToken() *Token {
	homeDir := utils.HomeDir()
	log.Info("Loading token from local file")
	tokenString, err := ioutil.ReadFile(homeDir + TokenFile)
	if err != nil {
		log.Error("Load json file failed, maybe there " + err.Error())
		return nil
	}

	var t Token
	json.Unmarshal(tokenString, &t)
	return &t
}

func WriteToFile(token *Token) {
	homeDir := utils.HomeDir()
	log.Info("Write updated token to file")
	tokenJSON, _ := json.Marshal(token)
	ioutil.WriteFile(homeDir+TokenFile, tokenJSON, 0644)
}

func FetchToken() *Token {
	log.Info("Fetching access token from remote server")
	response, err := http.Get(TokenURL)
	if err != nil {
		log.Fatal("Fetch access token failed. Please concat the Author: wahyd4@gmail.com: " + err.Error())
	}
	defer response.Body.Close()
	body, _ := ioutil.ReadAll(response.Body)
	var token Token
	if err = json.Unmarshal(body, &token); err != nil {
		log.Fatal("Unmarshal http response to token failed" + err.Error())
	}
	return &token
}

func TokenValid(token *Token) bool {
	return token != nil && token.Value != "" && time.Now().Before(time.Unix(token.ExpiresAt, 0))
}
