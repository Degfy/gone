package mock

import (
	"github.com/urfave/cli/v2"
	"os"
	"path/filepath"
)

func CreateCommand() *cli.Command {
	return &cli.Command{
		Name:        "mock",
		Usage:       "-f ${fromGoFile} -o ${outGoFile}",
		Description: "patch code which is generated by github.com/golang/mock/mockgen, add `gone.flag` to MockStructure",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "f",
				Usage: "FromGoFile",
			},
			&cli.StringFlag{
				Name:     "o",
				Usage:    "outGoFile",
				Required: true,
			},
		},
		Action: action,
	}
}

func action(ctx *cli.Context) error {
	return doAction(ctx.String("f"), ctx.String("o"))
}

func doAction(fromfile, outfile string) error {
	err := os.MkdirAll(filepath.Dir(outfile), os.ModePerm)
	if err != nil {
		return err
	}

	outFile, err := os.Create(outfile)
	if err != nil {
		return err
	}

	defer outFile.Close()

	if isInputFromPipe() {
		return patchMock(os.Stdin, outFile)
	} else {
		file, e := getFile(fromfile)
		if e != nil {
			return e
		}
		defer file.Close()
		return patchMock(file, outFile)
	}
}