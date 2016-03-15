package archive

import (
	"io"
	"os"
)

type IteratorFunc func(path string, info os.FileInfo, reader io.Reader)

type Iterator interface {
	Iterate(IteratorFunc) error
}
