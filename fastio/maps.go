package fastio

import (
	"bytes"

	"github.com/plzfgme/gosse/fastio/internal/pb"
	"github.com/plzfgme/gosse/storage"
)

func sigmaGet(r storage.Retriever, w []byte) ([]byte, uint64, error) {
	b, err := r.Get(sigmaKey(w))
	if err != nil {
		return nil, 0, err
	}

	st, c, err := pb.UnmarshalSigmaMapValue(b)
	if err != nil {
		return nil, 0, err
	}

	return st, c, nil
}

func sigmaSet(m storage.RetrieverMutator, w []byte, st []byte, c uint64) error {
	b, err := pb.MarshalSigmaMapValue(st, c)
	if err != nil {
		return err
	}

	err = m.Set(sigmaKey(w), b)
	if err != nil {
		return err
	}

	return nil
}

func tcGet(r storage.Retriever, tw []byte) ([][]byte, error) {
	b, err := r.Get(tcKey(tw))
	if err != nil {
		return nil, err
	}
	if b == nil {
		return nil, nil
	}

	l, err := pb.UnmarshalIdList(b)
	if err != nil {
		return nil, err
	}

	return l, nil
}

func tcSet(m storage.Mutator, tw []byte, ids [][]byte) error {
	b, err := pb.MarshalIdList(ids)
	if err != nil {
		return err
	}

	err = m.Set(tcKey(tw), b)
	if err != nil {
		return err
	}

	return nil
}

func teGet(r storage.Retriever, u []byte) ([]byte, error) {
	b, err := r.Get(teKey(u))
	if err != nil {
		return nil, err
	}

	return b, nil
}

func teSet(m storage.Mutator, u []byte, e []byte) error {
	err := m.Set(teKey(u), e)
	if err != nil {
		return err
	}

	return nil
}

func teDelete(m storage.Mutator, u []byte) error {
	return m.Delete(teKey(u))
}

func sigmaKey(w []byte) []byte {
	return bytes.Join([][]byte{[]byte("s"), w}, nil)
}

func teKey(u []byte) []byte {
	return bytes.Join([][]byte{[]byte("e"), u}, nil)
}

func tcKey(tw []byte) []byte {
	return bytes.Join([][]byte{[]byte("c"), tw}, nil)
}
