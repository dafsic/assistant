package main

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"io/ioutil"
	"net/http"
)

var switchCmd = &cli.Command{
	Name:  "switch",
	Usage: "On/off the auto pledge",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:  "assapi",
			Value: "192.168.28.172:6660",
			Usage: "指定miner地址, default 192.168.28.172:6660",
		},
	},
	Subcommands: []*cli.Command{
		{
			Name: "on",
			Action: func(c *cli.Context) error {
				resp, err := http.Get(fmt.Sprintf("http://%s/switch/on", c.String("assapi")))
				if err != nil {
					fmt.Println(err.Error())
					return err
				}
				defer resp.Body.Close()
				body, _ := ioutil.ReadAll(resp.Body)
				fmt.Println(string(body))
				return nil
			},
		},
		{
			Name: "off",
			Action: func(c *cli.Context) error {
				resp, err := http.Get(fmt.Sprintf("http://%s/switch/off", c.String("assapi")))
				if err != nil {
					fmt.Println(err.Error())
					return err
				}
				defer resp.Body.Close()
				body, _ := ioutil.ReadAll(resp.Body)
				fmt.Println(string(body))
				return nil
			},
		},
		{
			Name: "show",
			Action: func(c *cli.Context) error {
				resp, err := http.Get(fmt.Sprintf("http://%s/switch/state", c.String("assapi")))
				if err != nil {
					fmt.Println(err.Error())
					return err
				}
				defer resp.Body.Close()
				body, _ := ioutil.ReadAll(resp.Body)
				fmt.Println(string(body))
				return nil
			},
		},
	},
}
