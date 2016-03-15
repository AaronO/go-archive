package archive

import (
	"archive/zip"
	"io"
)

type ZipIterator struct {
	reader *zip.Reader
}

func NewZipIterator(reader io.ReaderAt, size int64) (*ZipIterator, error) {
	zr, err := zip.NewReader(reader, size)
	if err != nil {
		return nil, err
	}

	return &ZipIterator{
		zr,
	}, nil
}

// Iterate walks over all the files in the zip and calls the callback for each
func (z *ZipIterator) Iterate(fn IteratorFunc) error {
	for _, file := range z.reader.File {
		if err := z.iterateFile(file, fn); err != nil {
			return err
		}
	}
	return nil
}

// iterateFile handles the operations for a given file
func (z *ZipIterator) iterateFile(file *zip.File, fn IteratorFunc) error {
	// Open reader
	rc, err := file.Open()
	if err != nil {
		return err
	}

	// Call the iterator
	if err := fn(
		file.Name,
		file.FileInfo(),
		rc,
	); err != nil {
		return err
	}

	// Close reader
	return rc.Close()
}
