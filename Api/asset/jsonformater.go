package asset

import (
	"encoding/json"
	"io"
)

func FromJSON(r io.Reader, p interface{}) error {
	e := json.NewDecoder(r)
	return e.Decode(p)

}
func ToJson(w io.Writer, p interface{}) error {
	e := json.NewEncoder(w)
	return e.Encode(p)
}
