package archive

import (
	"archive/tar"
	"io"
)

type TarIterator struct {
	reader *tar.Reader
}

func NewTarIterator(reader io.Reader) (*TarIterator, error) {
	return &TarIterator{
		tar.NewReader(reader),
	}, nil
}

func (t *TarIterator) Iterate(fn IteratorFunc) error {
	for {
		header, err := t.reader.Next()
		if err == io.EOF {
			// end of tar archive
			break
		} else if err != nil {
			// Other error
			return err
		}
		// Skip pax_global_header
		if header.FileInfo().Name() == "pax_global_header" {
			continue
		}
		// Handle this specifc file
		if err := t.iterateFile(header, fn); err != nil {
			return err
		}
	}

	return nil
}

func (t *TarIterator) iterateFile(header *tar.Header, fn IteratorFunc) error {
	return fn(
		header.Name,
		header.FileInfo(),
		t.reader,
	)
}
