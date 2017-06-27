package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli"
)

var (
	person int
	speed  int
	pitch  int
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
		fmt.Println(words)

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
