package scanner

import "log"

type IScanner interface {
	GetAllFiles() // returns all files from cwd
}

type Scanner struct {
	IScanner
	log      *log.Logger
	dir_name string
}

func NewScanner(logger *log.Logger, dir_name string) *Scanner {
	return &Scanner{
		log:      logger,
		dir_name: dir_name,
	}
}
