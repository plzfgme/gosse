package fastio

import (
	"encoding/binary"

	mapset "github.com/deckarep/golang-set/v2"
	"github.com/plzfgme/gosse/fastio/internal/pb"
	"github.com/plzfgme/gosse/storage"
)

// Server represents a FASTIO server.
type Server struct {
	h1 *h1Hash
	h2 *h2Hash
}

// NewServer creates a new server.
func NewServer() *Server {
	return &Server{
		h1: newH1Hash(),
		h2: newH2Hash(),
	}
}

// UpdateWithToken updates the storage rm with token tkn.
//
// All error of rm are returned directly.
func (server *Server) UpdateWithToken(rm storage.RetrieverMutator, tkn []byte) error {
	u, e, err := pb.UnmarshalUpdateToken(tkn)
	if err != nil {
		return err
	}
	err = teSet(rm, u, e)
	if err != nil {
		return err
	}

	return nil
}

// SearchWithToken searchs the storage rm with token tkn.
//
// All error of rm are returned directly.
func (server *Server) SearchWithToken(rm storage.RetrieverMutator, tkn []byte) ([][]byte, error) {
	tw, kw, c, err := pb.UnmarshalSearchToken(tkn)
	if err != nil {
		return nil, err
	}

	cachedIds, err := tcGet(rm, tw)
	if err != nil && err != storage.ErrKeyNotFound {
		return nil, err
	}
	if kw == nil {
		return cachedIds, nil
	}

	idSet := mapset.NewSet[string]()
	for _, id := range cachedIds {
		idSet.Add(string(id))
	}

	for i := uint64(1); i <= c; i++ {
		input := binary.BigEndian.AppendUint64(kw, i)
		u, err := server.h1.Eval(input)
		if err != nil {
			return nil, err
		}

		e, err := teGet(rm, u)
		if err != nil {
			return nil, err
		}
		ePart2, err := server.h2.Eval(input)
		if err != nil {
			return nil, err
		}
		ePart1 := xor32Bytes(e, ePart2)
		flag, id := parseEPart1(ePart1)
		id = unpadId(id)

		if flag == flagDelete {
			idSet.Remove(string(id))
		} else {
			idSet.Add(string(id))
		}

		err = teDelete(rm, u)
		if err != nil {
			return nil, err
		}
	}

	strIds := idSet.ToSlice()
	ids := make([][]byte, len(strIds))
	for i, v := range strIds {
		ids[i] = []byte(v)
	}

	err = tcSet(rm, tw, ids)
	if err != nil {
		return nil, err
	}

	return ids, nil
}
