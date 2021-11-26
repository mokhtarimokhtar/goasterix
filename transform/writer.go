package transform

import (
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
	j, err = json.Marshal(w)
	return j, err
}
func WriteModelXML(w Writer, record goasterix.Record) (x []byte, err error) {
	w.write(record)
	x, err = xml.Marshal(w)
	return x, err
}
