package rangecover

// InsertKeywords converts a uint64 value to keywords for insert, and then it can be searched by SearchKeywords methods.
func InsertKeywords(v uint64) ([][]byte, error) {
	res := make([][]byte, 64)
	var err error
	for i := 0; i < 64; i++ {
		res[i], err = makeKeyword(v>>i, i)
		if err != nil {
			return nil, err
		}
	}

	return res, nil
}
