package main

import (
	"fmt"
	"github.com/dafsic/assistant/node"
	"github.com/urfave/cli/v2"
)

var switchCmd = &cli.Command{
	Name:  "switch",
	Usage: "On/off the auto pledge",
	Subcommands: []*cli.Command{
		{
			Name: "on",
			Action: func(c *cli.Context) error {
				node.On()
				return nil
			},
		},
		{
			Name: "off",
			Action: func(c *cli.Context) error {
				node.Off()
				return nil
			},
		},
		{
			Name: "show",
			Action: func(c *cli.Context) error {
				s := node.State()
				if s == 0 {
					fmt.Println("Off")
				} else if s == 1 {
					fmt.Println("On")
				} else {
					fmt.Println("Unknown")
				}
				return nil
			},
		},
	},
}
