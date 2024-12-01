package parser

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
)

func (p *Parser) ParseDir(dir_name string) ([]string, error) {
	// returns a list of all files and folders in cwd
	// parse all files in the directory
	files, err := os.ReadDir(dir_name)
	if err != nil {
		err_message := fmt.Sprintf("error while parsing directory (%s): %s", dir_name, err.Error())
		p.Log.Println(err_message)
		return nil, err
	}

	result := []string{}

	for _, val := range files {
		if val == nil {
			continue
		}
		name := path.Join(dir_name, val.Name())

		if val.IsDir() {
			res, err := p.ParseDir(name)
			if err != nil {
				panic(err.Error())
			}
			result = append(result, res...)
		}

		if filepath.Ext(name) != ".go" {
			continue
		}

		result = append(result, name)
	}

	return result, nil
}
