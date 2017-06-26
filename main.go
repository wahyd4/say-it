package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "say-it"
	app.Usage = "TTS in command line -- Pronounce the Chinese or English words you typed in."
	app.Version = "0.2.0"
	app.Action = func(c *cli.Context) error {
		if len(c.Args()) == 0 {
			fmt.Println("Please type some words. e.g: say-it '你好, 世界'")
			return nil
		}
		fmt.Println(c.Args().Get(0))

		return nil
	}

	app.Run(os.Args)
}
