package poker

import (
	"os"
)

type tape struct {
	file *os.File
}

func (t *tape) Write(p []byte) (n int, err error) {
	t.file.Truncate(0) // effectively delete all content from file
	t.file.Seek(0, 0)
	return t.file.Write(p)
}
