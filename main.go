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

	log "github.com/Sirupsen/logrus"
	"github.com/urfave/cli"
)

const (
	Mp3FileName = "say-it.mp3"
)

var (
	person int
	speed  int
	pitch  int
	token  string = "24.85efc4bbd8b8315255dcb59cd82e1ac4.2592000.1501048297.282335-9739014"
)

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
		fetchVoiceAndSpeak(words)
		fmt.Println(person, speed, pitch)
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
	queries.Add("tok", token)
	queries.Add("per", strconv.Itoa(person))
	queries.Add("spd", strconv.Itoa(speed))
	queries.Add("pit", strconv.Itoa(pitch))
	urlObject.RawQuery = queries.Encode()
	return urlObject
}
