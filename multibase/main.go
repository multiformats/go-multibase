package main

import (
	"fmt"
	"github.com/multiformats/go-multibase"
	"io/ioutil"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "multibase",
		Usage: "base encoding and transcoding tool",
		Commands: []*cli.Command{
			{
				Name:      "encode",
				ArgsUsage: "<base>",
				Usage:     "encode data in multibase",
				Flags: []cli.Flag{
					&cli.PathFlag{
						Name:    "input",
						Aliases: []string{"i"},
						Usage:   "the file that should be encoded",
					},
				},
				Action: func(context *cli.Context) error {
					if !context.Args().Present() || context.NArg() > 2 {
						return cli.ShowCommandHelp(context, "")
					}

					p := context.Path("input")
					if (p == "" && context.NArg() != 2) || (p != "" && context.NArg() != 1) {
						return cli.ShowCommandHelp(context, "")
					}

					base, err := multibase.EncoderByName(context.Args().First())
					if err != nil {
						return err
					}

					if p != "" {
						fileData, err := ioutil.ReadFile(p)
						if err != nil {
							return err
						}
						fmt.Println(base.Encode(fileData))
						return nil
					}

					fmt.Println(base.Encode([]byte(context.Args().Get(1))))
					return nil
				},
			},
			{
				Name:      "decode",
				ArgsUsage: "<data>",
				Usage:     "encode data in multibase",
				Flags: []cli.Flag{
					&cli.PathFlag{
						Name:    "output",
						Aliases: []string{"o"},
						Usage:   "output decoded data to a file",
					},
				},
				Action: func(context *cli.Context) error {
					if context.NArg() != 1 {
						return cli.ShowCommandHelp(context, "")
					}

					_, data, err := multibase.Decode(context.Args().First())
					if err != nil {
						return err
					}

					p := context.Path("output")
					if p == "" {
						fmt.Printf(string(data))
						return nil
					}
					return ioutil.WriteFile(p, data, os.ModePerm)
				},
			},
			{
				Name:      "transcode",
				ArgsUsage: "<new-base> <data>",
				Usage:     "transcode multibase data",
				Action: func(context *cli.Context) error {
					if context.NArg() != 2 {
						return cli.ShowCommandHelp(context, "")
					}

					newbase, err := multibase.EncoderByName(context.Args().Get(0))
					if err != nil {
						return err
					}

					_, data, err := multibase.Decode(context.Args().Get(1))
					if err != nil {
						return err
					}

					fmt.Println(newbase.Encode(data))
					return nil
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		panic(err)
	}
}
