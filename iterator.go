package archive

import (
	"io"
	"os"
)

type IteratorFunc func(path string, info os.FileInfo, reader io.Reader) error

type Iterator interface {
	Iterate(IteratorFunc) error
}
