package transform

import (
	"bytes"
	"encoding/json"
	"encoding/xml"

	"github.com/mokhtarimokhtar/goasterix"
)

type Writer interface {
	write(record goasterix.Record)
}

func WriteModel(w Writer, record goasterix.Record) {
	w.write(record)
}

func WriteModelJSON(w Writer, record goasterix.Record) (j []byte, err error) {
	w.write(record)
	j, err = JSONMarshal(w)
	return j, err
}
func WriteModelXML(w Writer, record goasterix.Record) (x []byte, err error) {
	w.write(record)
	x, err = xml.Marshal(w)
	return x, err
}

func JSONMarshal(v interface{}) ([]byte, error) {
	b, err := json.Marshal(v)

	b = bytes.Replace(b, []byte("\\u003c"), []byte("<"), -1)
	b = bytes.Replace(b, []byte("\\u003e"), []byte(">"), -1)
	b = bytes.Replace(b, []byte("\\u0026"), []byte("&"), -1)

	return b, err
}
