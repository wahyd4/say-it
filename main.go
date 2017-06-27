package main

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"strconv"

	"time"

	"net/http"
	"net/url"

	"errors"

	log "github.com/Sirupsen/logrus"
	"github.com/urfave/cli"
	t "github.com/wahyd4/say-it/token"
	"github.com/wahyd4/say-it/utils"
)

const (
	Mp3FileName = "say-it.mp3"
)

var (
	person int
	speed  int
	pitch  int
	token  *t.Token
)

func init() {
	//try to load token
	//if no token, then fetch
	//fetch token and write token
	token = t.LoadToken()
	if !t.TokenValid(token) {
		log.Info("No valid token found or token expires, will try to fetch one")
		token = t.FetchToken()
		t.WriteToFile(token)
	}
}

func main() {

	app := cli.NewApp()
	app.Name = "say-it"
	app.Usage = "TTS in command line -- Pronounce the Chinese or English words you typed in."
	app.Version = "0.2.0"
	setFlags(app)
	app.Action = func(c *cli.Context) error {
		if len(c.Args()) == 0 {
			fmt.Println("Please type some words. e.g: say-it '你好, 世界'")
			return nil
		}
		words := c.Args().Get(0)
		if err := inputCheck(); err != nil {
			log.Fatal("Please type the correct option values and retry : " + err.Error())
		}
		fetchVoiceAndSpeak(words)
		return nil
	}
	app.Run(os.Args)
}

func setFlags(app *cli.App) {
	app.Flags = []cli.Flag{
		cli.IntFlag{
			Name:        "person, p",
			Value:       0,
			Destination: &person,
			Usage:       "set different voice. 0: female, 1 and 2: male, 3: male with emotion, 4: female with emotion",
		},
		cli.IntFlag{
			Name:        "speed, s",
			Value:       5,
			Destination: &speed,
			Usage:       "set speed of voice. 0 - 9",
		},
		cli.IntFlag{
			Name:        "pitch, t",
			Value:       5,
			Destination: &pitch,
			Usage:       "set the voice pitch. 0 - 9",
		},
	}
}

func inputCheck() error {
	if !utils.Contains(person, []int{0, 1, 2, 3, 4}) {
		return errors.New("Person value is not valid")
	}

	if !utils.Contains(speed, []int{1, 2, 3, 4, 5, 6, 7, 8, 9}) {
		return errors.New("Speed value is not valid")
	}

	if !utils.Contains(pitch, []int{1, 2, 3, 4, 5, 6, 7, 8, 9}) {
		return errors.New("Pitch value is not valid")
	}
	return nil
}

func fetchVoiceAndSpeak(text string) {
	urlObject := buildURL(text)
	response, err := http.Get(urlObject.String())

	if err != nil {
		log.Error("Fetch voice failed:" + err.Error())
	}
	// fmt.Println(response.StatusCode)

	defer response.Body.Close()
	out, err := os.Create(Mp3FileName)
	if err != nil {
		log.Fatal("Create file failed:" + err.Error())
	}
	defer out.Close()
	io.Copy(out, response.Body)

	speak()
}

func speak() {
	command := exec.Command("afplay", Mp3FileName)
	if err := command.Run(); err != nil {
		log.Error("Failed to say the words: " + err.Error())
	}
}

func buildURL(text string) *url.URL {
	baseURL := "http://tsn.baidu.com/text2audio"
	urlObject, _ := url.Parse(baseURL)

	queries := url.Values{}
	queries.Add("tex", text)
	queries.Add("lan", "zh")
	queries.Add("cuid", strconv.FormatInt(time.Now().Unix(), 10))
	queries.Add("ctp", "1")
	queries.Add("tok", token.Value)
	queries.Add("per", strconv.Itoa(person))
	queries.Add("spd", strconv.Itoa(speed))
	queries.Add("pit", strconv.Itoa(pitch))
	urlObject.RawQuery = queries.Encode()
	return urlObject
}
