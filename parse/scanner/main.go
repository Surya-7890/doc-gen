package scanner

import "log"

type IScanner interface {
	GetAllFiles() // returns all files from cwd
}

type Scanner struct {
	IScanner
	log       *log.Logger
	filenames chan string
}

func NewScanner(logger *log.Logger, channel chan string) *Scanner {
	return &Scanner{
		log:       logger,
		filenames: channel,
	}
}
