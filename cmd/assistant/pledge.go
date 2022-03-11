package main

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"io/ioutil"
	"net/http"
)

var pledgeCmd = &cli.Command{
	Name:  "pledge",
	Usage: "Add a CC sector",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:  "minerapi",
			Value: "192.168.28.136:6660",
			Usage: "指定miner地址, default 192.168.28.136:6660",
		},
	},

	Action: func(cctx *cli.Context) error {
		url := fmt.Sprintf("http://%s/pledge", cctx.String("minerapi"))
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			fmt.Println(err.Error())
			return err
		}

		client := http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			fmt.Println(err.Error())
			return err
		}
		defer resp.Body.Close()

		contentBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println(err.Error())
			return err
		}

		fmt.Println(string(contentBytes))
		return nil
	},
}
