package main

import (
	"fmt"
	"log"
	"os"

	"github.com/codex-team/capella.go"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()

	app.Action = func(c *cli.Context) error {

		// upload image from local path

		filepath := c.Args().Get(0)
		response, errC := capella.UploadFile(filepath)

		if errC.Message != "" {
			log.Fatal(errC.Message)
		}
		fmt.Printf(response.URL)
		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}

}
