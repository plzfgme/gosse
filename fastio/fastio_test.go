package fastio

import (
	"math/rand"
	"testing"

	"github.com/plzfgme/gosse/internal/simplestorage"
	"github.com/stretchr/testify/suite"
)

var (
	_ suite.SetupTestSuite    = (*fastioTestSuite)(nil)
	_ suite.TearDownTestSuite = (*fastioTestSuite)(nil)
)

type fastioTestSuite struct {
	suite.Suite

	crm *simplestorage.RetrieverMutator
	srm *simplestorage.RetrieverMutator

	c *Client
	s *Server
}

func (s *fastioTestSuite) SetupTest() {
	s.crm = simplestorage.NewSimpleRetrieverMutator()
	s.srm = simplestorage.NewSimpleRetrieverMutator()

	key, err := randBytes(KeySize)
	s.Require().Nil(err)
	s.Require().Len(key, KeySize)
	copt := NewClientOptions(key)
	s.c, err = NewClient(copt)
	s.Require().Nil(err)
	s.Require().NotNil(s.c)

	s.s = NewServer()
	s.Require().NotNil(s.s)
}

func (s *fastioTestSuite) TearDownTest() {
	s.c = nil
	s.s = nil
	s.crm = nil
	s.srm = nil
}

func (s *fastioTestSuite) mustAddOk(w, id []byte) {
	tkn, err := s.c.GenInsertToken(s.crm, w, id)
	s.Require().Nil(err)
	s.Require().NotNil(tkn)
	err = s.s.UpdateWithToken(s.srm, tkn)
	s.Require().Nil(err)
}

func (s *fastioTestSuite) mustDeleteOk(w, id []byte) {
	tkn, err := s.c.GenDeleteToken(s.crm, w, id)
	s.Require().Nil(err)
	s.Require().NotNil(tkn)
	err = s.s.UpdateWithToken(s.srm, tkn)
	s.Require().Nil(err)
}

func (s *fastioTestSuite) mustSearchMatched(w []byte, expectedIds [][]byte) {
	tkn, err := s.c.GenSearchToken(s.crm, w)
	s.Require().Nil(err)
	if tkn == nil && len(expectedIds) == 0 {
		return
	}
	s.Require().NotNil(tkn)
	ids, err := s.s.SearchWithToken(s.srm, tkn)
	s.Require().Nil(err)
	s.Require().ElementsMatch(expectedIds, ids)
}

func (s *fastioTestSuite) TestUpdateSearch() {
	w1 := []byte("w1")
	w2 := []byte("w2")
	s.mustSearchMatched(w1, [][]byte{})
	s.mustSearchMatched(w2, [][]byte{})

	id11 := []byte("id11")
	id12 := []byte("id12")
	id21 := []byte("id21")
	s.mustAddOk(w1, id11)
	s.mustAddOk(w1, id12)
	s.mustAddOk(w2, id21)
	s.mustSearchMatched(w1, [][]byte{id11, id12})
	s.mustSearchMatched(w2, [][]byte{id21})

	s.mustDeleteOk(w1, id11)
	s.mustDeleteOk(w2, id21)
	s.mustSearchMatched(w1, [][]byte{id12})
	s.mustSearchMatched(w2, [][]byte{})

	id13 := []byte("id13")
	s.mustAddOk(w1, id13)
	s.mustSearchMatched(w1, [][]byte{id12, id13})
	s.mustSearchMatched(w2, [][]byte{})
}

func TestFASTIO(t *testing.T) {
	suite.Run(t, &fastioTestSuite{})
}

func randBytes(l int) ([]byte, error) {
	b := make([]byte, l)
	_, err := rand.Read(b)

	return b, err
}
