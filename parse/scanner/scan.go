package scanner

import (
	"os"

	"golang.org/x/tools/go/packages"
)

func (s *Scanner) GetAllFiles() map[string][]string {
	files := make(map[string][]string)

	dir, err := os.Getwd()
	if err != nil {
		s.log.Fatal(err.Error())
	}

	cfg := &packages.Config{
		Mode: packages.NeedName | packages.NeedFiles,
	}

	pkgs, err := packages.Load(cfg, dir+"/example/...")
	if err != nil {
		s.log.Fatal(err.Error())
	}

	for _, pkg := range pkgs {
		files[pkg.Name] = append(files[pkg.Name], pkg.GoFiles...)
	}

	return files
}
