package simplestorage

import "github.com/plzfgme/gosse/storage"

var _ storage.RetrieverMutator = (*RetrieverMutator)(nil)

type RetrieverMutator map[string][]byte

func NewSimpleRetrieverMutator() *RetrieverMutator {
	rm := make(RetrieverMutator)

	return &rm
}

func (rm *RetrieverMutator) Get(k []byte) ([]byte, error) {
	if v, ok := (*rm)[string(k)]; ok {
		return v, nil
	} else {
		return nil, storage.ErrKeyNotFound
	}
}

func (rm *RetrieverMutator) Set(k []byte, v []byte) error {
	(*rm)[string(k)] = v

	return nil
}

func (rm *RetrieverMutator) Delete(k []byte) error {
	delete(*rm, string(k))

	return nil
}
