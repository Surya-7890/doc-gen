package scanner

import (
	"fmt"
	"go/ast"
	"os"
	"path"
	"path/filepath"
)

func (s *Scanner) GetAllFiles() []*ast.File {
	files := []*ast.File{}

	dir, err := os.Getwd()
	if err != nil {
		s.log.Fatal(err.Error())
	}

	dir = dir + "/example"
	filenames := s.getFilesFromDir(dir)

	for _, filename := range filenames {
		fmt.Println(filename)
	}

	return files
}

func (s *Scanner) getFilesFromDir(dir_name string) []string {
	filenames := []string{}
	entries, err := os.ReadDir(dir_name)
	if err != nil {
		s.log.Fatal(err.Error())
	}

	for _, entry := range entries {
		if entry.IsDir() {
			res := s.getFilesFromDir(path.Join(dir_name, entry.Name()))
			filenames = append(filenames, res...)
			continue
		}

		if filepath.Ext(entry.Name()) == ".go" {
			filenames = append(filenames, entry.Name())
		}
	}

	return filenames
}
