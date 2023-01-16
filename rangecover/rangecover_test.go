package rangecover

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func mustMatched(t *testing.T, ikws, skws [][]byte) {
	m := make(map[string]struct{})

	for _, kw := range ikws {
		m[string(kw)] = struct{}{}
	}

	for _, kw := range skws {
		if _, ok := m[string(kw)]; ok {
			return
		}
	}

	t.Fatal("No insert keyword contains in search keywords")
}

func mustNotMatched(t *testing.T, ikws, skws [][]byte) {
	m := make(map[string]struct{})

	for _, kw := range ikws {
		m[string(kw)] = struct{}{}
	}

	for _, kw := range skws {
		if _, ok := m[string(kw)]; ok {
			t.Fatalf("There is a insert keyword %v contains in search keywords", kw)
		}
	}
}

func TestInsertKeywords(t *testing.T) {
	ikws, err := InsertKeywords(4455)
	require.Nil(t, err)
	require.Len(t, ikws, 64)
}

func TestBRC(t *testing.T) {
	ikws, err := InsertKeywords(2333)
	require.Nil(t, err)

	skws1, err := BRCSearchKeywords(123, 49999)
	require.Nil(t, err)
	mustMatched(t, ikws, skws1)

	skws2, err := BRCSearchKeywords(2590, 2777)
	require.Nil(t, err)
	mustNotMatched(t, ikws, skws2)
}
