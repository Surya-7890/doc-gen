package scanner

import (
	"gen-doc/types"
	"os"

	"golang.org/x/tools/go/packages"
)

func (s *Scanner) GetAllFiles() map[*types.MapKey]*packages.Package {
	files := make(map[*types.MapKey]*packages.Package)

	dir, err := os.Getwd()
	if err != nil {
		s.log.Fatal(err.Error())
	}

	cfg := &packages.Config{
		Mode: packages.NeedName | packages.NeedFiles | packages.NeedSyntax | packages.NeedTypes | packages.NeedTypesInfo,
	}

	pkgs, err := packages.Load(cfg, dir+s.dir_name+"/...")
	if err != nil {
		s.log.Fatal(err.Error())
	}

	for _, pkg := range pkgs {
		name := &types.MapKey{
			Name: pkg.Name,
			Path: pkg.PkgPath,
		}
		files[name] = pkg
	}

	return files
}
