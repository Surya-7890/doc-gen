package scanner

import "log"

type IScanner interface {
	GetAllFiles() // returns all files from cwd
}

type Scanner struct {
	IScanner
	log *log.Logger
}

func NewScanner(logger *log.Logger) *Scanner {
	return &Scanner{
		log: logger,
	}
}
