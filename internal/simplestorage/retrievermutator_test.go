package simplestorage

import (
	"testing"

	"github.com/plzfgme/gosse/storage"
	"github.com/stretchr/testify/suite"
)

var (
	_ suite.SetupTestSuite    = (*retrieverMutatorTestSuite)(nil)
	_ suite.TearDownTestSuite = (*retrieverMutatorTestSuite)(nil)
)

type retrieverMutatorTestSuite struct {
	suite.Suite

	rm *RetrieverMutator
}

func (s *retrieverMutatorTestSuite) SetupTest() {
	s.rm = NewSimpleRetrieverMutator()
}

func (s *retrieverMutatorTestSuite) TearDownTest() {
	s.rm = nil
}

func (s *retrieverMutatorTestSuite) mustGetMatched(k, expectedV []byte) {
	v, err := s.rm.Get(k)
	s.Require().Nil(err)
	s.Require().Equal(expectedV, v)
}

func (s *retrieverMutatorTestSuite) mustGetNotFound(k []byte) {
	_, err := s.rm.Get(k)
	s.Require().ErrorIs(err, storage.ErrKeyNotFound)
}

func (s *retrieverMutatorTestSuite) mustSetOk(k, v []byte) {
	err := s.rm.Set(k, v)
	s.Require().Nil(err)
}

func (s *retrieverMutatorTestSuite) mustDeleteOk(k []byte) {
	err := s.rm.Delete(k)
	s.Require().Nil(err)
}

func (s *retrieverMutatorTestSuite) TestGetSetDelete() {
	key1 := []byte("key1")
	s.mustGetNotFound(key1)

	val1 := []byte("val1")
	s.mustSetOk(key1, val1)
	s.mustGetMatched(key1, val1)

	val2 := []byte("val2")
	s.mustSetOk(key1, val2)
	s.mustGetMatched(key1, val2)

	s.mustDeleteOk(key1)
	s.mustGetNotFound(key1)

	val3 := []byte("val3")
	s.mustSetOk(key1, val3)
	s.mustGetMatched(key1, val3)

	key2 := []byte("key2")
	s.mustGetNotFound(key2)
}

func TestRetrieverMutator(t *testing.T) {
	suite.Run(t, &retrieverMutatorTestSuite{})
}
