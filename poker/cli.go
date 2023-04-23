package poker

import (
	"bufio"
	"io"
	"strings"
)

// Package bufio implements buffered I/O. It wraps an io.Reader or io.Writer object,
// creating another object (Reader or Writer) that also implements the interface
// but provides buffering and some help for textual I/O.

type CLI struct {
	store PlayerStore
	in    *bufio.Scanner
}

func NewCLI(store PlayerStore, in io.Reader) *CLI {
	return &CLI{
		store: store,
		in:    bufio.NewScanner(in),
	}
}

func (cli *CLI) PlayPoker() {
	userInput := cli.readLine()
	cli.store.RecordWin(extractWinner(userInput)) // scenner.Text() return the string the scanner read to
}

func extractWinner(userInput string) string {
	return strings.Replace(userInput, " wins", "", 1)
}

func (cli *CLI) readLine() string {
	cli.in.Scan()
	return cli.in.Text()
}
