package archive

import (
	"compress/gzip"
	"io"
)

type TargzIterator struct {
	reader   *gzip.Reader
	iterator *TarIterator
}

func NewTargzIterator(reader io.Reader) (*TargzIterator, error) {
	// init gzip stream reader
	greader, err := gzip.NewReader(reader)
	if err != nil {
		return nil, err
	}

	// init tar iterator
	iterator, err := NewTarIterator(greader)
	if err != nil {
		return nil, err
	}

	return &TargzIterator{
		reader:   greader,
		iterator: iterator,
	}, nil
}

func (t *TargzIterator) Iterate(fn IteratorFunc) error {
	return t.iterator.Iterate(fn)
}

func (t *TargzIterator) Close() error {
	return t.reader.Close()
}
