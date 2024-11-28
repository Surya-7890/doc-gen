package cli

import (
	"fmt"
	"gen-doc/utils"
	"os"
)

type CLI struct {
	RW *utils.ReadWriter
}

func NewCLI() *CLI {
	return &CLI{
		RW: utils.NewReadWriter(os.Stdin, os.Stdout),
	}
}

func (c *CLI) Parse() {
	path, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	fmt.Println(path)
}
