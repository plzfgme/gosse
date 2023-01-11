package pb

import "google.golang.org/protobuf/proto"

func MarshalSigmaMapValue(st []byte, c uint64) ([]byte, error) {
	stc := &FASTIOSigmaMapValue{
		St: st,
		C:  c,
	}
	b, err := proto.Marshal(stc)
	if err != nil {
		return nil, err
	}

	return b, nil
}

func UnmarshalSigmaMapValue(b []byte) ([]byte, uint64, error) {
	stc := &FASTIOSigmaMapValue{}
	err := proto.Unmarshal(b, stc)
	if err != nil {
		return nil, 0, err
	}

	return stc.GetSt(), stc.GetC(), nil
}

func MarshalIdList(l [][]byte) ([]byte, error) {
	s := &FASTIOIDList{
		Ids: l,
	}
	b, err := proto.Marshal(s)
	if err != nil {
		return nil, err
	}

	return b, nil
}

func UnmarshalIdList(b []byte) ([][]byte, error) {
	s := &FASTIOIDList{}
	err := proto.Unmarshal(b, s)
	if err != nil {
		return nil, err
	}

	return s.GetIds(), nil
}

func MarshalUpdateToken(u, e []byte) ([]byte, error) {
	tkn := &FASTIOUpdateToken{
		U: u,
		E: e,
	}
	b, err := proto.Marshal(tkn)
	if err != nil {
		return nil, err
	}

	return b, nil
}

func UnmarshalUpdateToken(b []byte) ([]byte, []byte, error) {
	tkn := &FASTIOUpdateToken{}
	err := proto.Unmarshal(b, tkn)
	if err != nil {
		return nil, nil, err
	}

	return tkn.GetU(), tkn.GetE(), nil
}

func MarshalSearchToken(tw, kw []byte, c uint64) ([]byte, error) {
	tkn := &FASTIOSearchToken{
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

func UnmarshalSearchToken(b []byte) ([]byte, []byte, uint64, error) {
	tkn := &FASTIOSearchToken{}
	err := proto.Unmarshal(b, tkn)
	if err != nil {
		return nil, nil, 0, err
	}

	return tkn.GetTW(), tkn.GetKW(), tkn.GetC(), nil
}
