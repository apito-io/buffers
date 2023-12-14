package shared

import (
	"encoding/json"
)

// Marshal & Unmarshal []byte result

type Result[T any] struct {
	Result  []*T `json:"result"`
	Count   int  `json:"count"`
	Cached  bool `json:"cached"`
	HasMore bool `json:"hasMore"`
	Error   bool `json:"error"`
	Code    int  `json:"code"`
}

type DataloaderResult[T any] struct {
	Result  []T  `json:"result"`
	Count   int  `json:"count"`
	Cached  bool `json:"cached"`
	HasMore bool `json:"hasMore"`
	Error   bool `json:"error"`
	Code    int  `json:"code"`
}

func UnmarshalDynamicCount[T any](body []byte) (*T, error) {
	var res Result[T]
	err := json.Unmarshal(body, &res)
	if err != nil {
		return nil, err
	}
	if res.Count > 0 {
		return res.Result[0], nil
	}
	return nil, nil
}

func UnmarshalDynamicResults[T any](body []byte) ([]*T, error) {
	var res Result[T]
	err := json.Unmarshal(body, &res)
	if err != nil {
		return nil, err
	}
	if res.Count > 0 {
		return res.Result, nil
	}
	return nil, nil
}

func UnmarshalDynamicResult[T any](body []byte) (*T, error) {
	var res Result[T]
	err := json.Unmarshal(body, &res)
	if err != nil {
		return nil, err
	}
	if res.Count > 0 {
		return res.Result[0], nil
	}
	return nil, nil
}
