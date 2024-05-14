package app

import (
	"bytes"
	"fmt"
	"github.com/gone-io/gone/templates"
	"github.com/urfave/cli/v2"
	"io/fs"
	"os"
	"path"
	"strings"
)

var appName, template, modName string

var flags = []cli.Flag{
	&cli.StringFlag{
		Name:        "t",
		Value:       "web",
		Usage:       "template: only support web、web+mysql, more will be supported in the future",
		Destination: &template,
	},

	&cli.StringFlag{
		Name:        "m",
		Usage:       "modName",
		Destination: &modName,
	},
}

var f = templates.F

func copyToAndReplace(source, target string, mode fs.FileMode, replace map[string]string) error {
	err := os.MkdirAll(target, mode)
	if err != nil {
		return err
	}

	dir, err := f.ReadDir(source)
	if err != nil {
		return err
	}

	for _, entry := range dir {
		if err != nil {
			return err
		}

		from := path.Join(source, entry.Name())
		to := path.Join(target, strings.TrimSuffix(entry.Name(), ".tpl"))

		if entry.IsDir() {
			err = copyToAndReplace(
				from,
				to,
				0766,
				replace,
			)
			if err != nil {
				return err
			}
		} else {
			file, err := os.OpenFile(to, os.O_RDWR|os.O_CREATE, 0644)
			if err != nil {
				return err
			}
			defer file.Close()

			data, err := f.ReadFile(from)
			if err != nil {
				return err
			}

			for k, v := range replace {
				data = bytes.ReplaceAll(data, []byte(k), []byte(v))
			}

			_, err = file.Write(data)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func action(c *cli.Context) error {
	if template != "web" && template != "web+mysql" {
		return fmt.Errorf("only support web、web+mysql, more will be supported in the future")
	}

	appName = c.Args().Get(0)
	if appName == "" {
		appName = "demo"
	}

	if modName == "" {
		modName = appName
	}

	return copyToAndReplace(
		path.Join(".", template),
		appName,
		0766,
		map[string]string{
			"${{appName}}":   appName,
			"demo-structure": modName,
		},
	)
}

var Command = &cli.Command{
	Name:        "create",
	Usage:       "[-t ${template} [-m ${modName}]] ${appName}",
	Description: "create a gone app",

	Flags:  flags,
	Action: action,
}
