package fastio

import (
	"crypto/rand"
	"encoding/binary"
	"errors"

	"github.com/plzfgme/gosse/fastio/internal/pb"
	"github.com/plzfgme/gosse/storage"
	"google.golang.org/protobuf/proto"
)

// ClientOptions represents the options of Client.
type ClientOptions struct {
	key []byte
}

// NewClientOptions creates a new ClientOptions from a key.
func NewClientOptions(key []byte) *ClientOptions {
	return &ClientOptions{
		key: key,
	}
}

// Client represents a FASTIO client
type Client struct {
	f  *fPRF
	h1 *h1Hash
	h2 *h2Hash
	h  *hHash
}

// NewClient creates a new Client from options, returns ErrKeySize if the key size is wrong.
func NewClient(opt *ClientOptions) (*Client, error) {
	if len(opt.key) != KeySize {
		return nil, ErrKeySize
	}

	f, _ := newFPRF(opt.key[:16], opt.key[16:])
	h1 := newH1Hash()
	h2 := newH2Hash()
	h := newHHash()

	return &Client{
		f:  f,
		h1: h1,
		h2: h2,
		h:  h,
	}, nil
}

// GenInsertToken generates insert token from the storage rm, the keyword w and the id, and returns ErrIDTooLong if id is too long.
//
// All error of rm are returned directly.
func (client *Client) GenInsertToken(rm storage.RetrieverMutator, w, id []byte) ([]byte, error) {
	return client.genUpdateToken(rm, w, id, true)
}

// GenDeleteToken generates delete token from the storage rm, the keyword w and the id, and returns ErrIDTooLong if id is too long.
//
// All error of rm are returned directly.
func (client *Client) GenDeleteToken(rm storage.RetrieverMutator, w, id []byte) ([]byte, error) {
	return client.genUpdateToken(rm, w, id, false)
}

func (client *Client) genUpdateToken(rm storage.RetrieverMutator, w, id []byte, add bool) ([]byte, error) {
	if len(id) > MaxIDSize {
		return nil, ErrIDTooLong
	}
	id = padId(id)

	st, c, err := sigmaGet(rm, w)
	if err != nil {
		if errors.Is(err, storage.ErrKeyNotFound) {
			st = make([]byte, 16)
			_, err := rand.Read(st)
			if err != nil {
				return nil, err
			}
			c = 0
		} else {
			return nil, err
		}
	}

	input := binary.BigEndian.AppendUint64(st, c+1)
	u, err := client.h1.Eval(input)
	if err != nil {
		return nil, err
	}
	var flag byte
	if add {
		flag = flagAdd
	} else {
		flag = byte(flagDelete)
	}
	ePart1 := buildEPart1(flag, id)
	ePart2, err := client.h2.Eval(input)
	if err != nil {
		return nil, err
	}
	e := xor32Bytes(ePart1, ePart2)

	err = sigmaSet(rm, w, st, c+1)
	if err != nil {
		return nil, err
	}

	tkn := &pb.FASTIOUpdateToken{
		U: u,
		E: e,
	}
	b, err := proto.Marshal(tkn)
	if err != nil {
		return nil, err
	}

	return b, nil
}

// GenSearchToken generates search token from the storage rm and the keyword w.
//
// All error of rm are returned directly.
func (client *Client) GenSearchToken(rm storage.RetrieverMutator, w []byte) ([]byte, error) {
	st, c, err := sigmaGet(rm, w)
	if err != nil {
		if errors.Is(err, storage.ErrKeyNotFound) {
			return nil, nil
		} else {
			return nil, err
		}
	}

	hw, err := client.h.Eval(w)
	if err != nil {
		return nil, err
	}
	tw := client.f.Eval(hw)

	var kw []byte
	if c != 0 {
		kw = st
		st = make([]byte, 16)
		_, err := rand.Read(st)
		if err != nil {
			return nil, err
		}

		err = sigmaSet(rm, w, st, 0)
		if err != nil {
			return nil, err
		}
	} else {
		kw = nil
	}

	tkn := &pb.FASTIOSearchToken{
		TW: tw,
		KW: kw,
		C:  c,
	}
	b, err := proto.Marshal(tkn)
	if err != nil {
		return nil, err
	}

	return b, nil
}
