package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"strconv"

	"time"

	"net/http"
	"net/url"

	"errors"

	"runtime"

	log "github.com/Sirupsen/logrus"
	"github.com/urfave/cli"
	t "github.com/wahyd4/say-it/token"
	"github.com/wahyd4/say-it/utils"
)

const (
	MP3FileName        = "/tmp/say-it.mp3"
	WindowsMP3FileName = "\\say-it.mp3"
)

var (
	person int
	speed  int
	pitch  int
	token  *t.Token
)

type ErrorResponse struct {
	Message string `json:"err_msg"`
	Code    int    `json:"err_no"`
}

func init() {
	if runtime.GOOS != "darwin" && runtime.GOOS != "windows" {
		log.Fatal("Sorry, currently we don't support OS: " + runtime.GOOS)
	}

	log.SetLevel(log.WarnLevel)
	//try to load token
	//if no token, then fetch and write to local
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
	app.Usage = "TTS in command line -- Pronounce the Chinese and English words you typed in."
	app.Version = "0.2.0"
	setFlags(app)
	app.Action = func(c *cli.Context) error {
		if len(c.Args()) == 0 {
			fmt.Println("Please type some words. e.g. say-it '你好, 世界'")
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
Fetch:
	urlObject := buildURL(text)
	response, err := http.Get(urlObject.String())

	if err != nil {
		log.Fatal("Fetch voice failed:" + err.Error())
		return
	}

	defer response.Body.Close()
	if utils.CheckContentType(response.Header["Content-Type"], "application/json") {

		bodyString, err := ioutil.ReadAll(response.Body)
		if err != nil {
			log.Error(err.Error())
		}
		var errorResp ErrorResponse
		json.Unmarshal(bodyString, &errorResp)

		if errorResp.Code == 502 {
			log.Warn("Access code is not valid, trying to fetch a new one")
			token = t.FetchToken()
			t.WriteToFile(token)
			goto Fetch
		}
		log.Fatalf("Get voice failed, error code is: %d, error message is %s", errorResp.Code, errorResp.Message)
	}

	out, err := os.Create(getAudioFilePath())
	if err != nil {
		log.Fatal("Create file failed:" + err.Error())
	}
	defer out.Close()
	io.Copy(out, response.Body)

	speak()
}

func speak() {
	if runtime.GOOS == "darwin" {
		command := exec.Command("afplay", getAudioFilePath())
		if err := command.Run(); err != nil {
			log.Error("Failed to say the words: " + err.Error())
		}
		return
	}
	command := exec.Command(getAudioFilePath())
	// command := exec.Command("cmdmp3.exe", getAudioFilePath())
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

func getAudioFilePath() string {
	if runtime.GOOS == "darwin" {
		return utils.HomeDir() + MP3FileName
	}
	fmt.Println(utils.HomeDir() + WindowsMP3FileName)
	return utils.HomeDir() + WindowsMP3FileName
}
