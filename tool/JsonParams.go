package tool

import (
	"encoding/json"
	"io"
)

type JsonParams struct {
}

func Decode(io io.ReadCloser, v interface{}) error {
	return json.NewDecoder(io).Decode(v)
}
