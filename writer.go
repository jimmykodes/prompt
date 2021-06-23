package prompt

import (
	"io"
	"os"

	"github.com/jimmykodes/cursor"
)

type writer struct {
	io.Writer
	Cursor *cursor.Cursor
}

var Writer = writer{
	Writer: os.Stderr,
	Cursor: cursor.New(os.Stderr),
}
