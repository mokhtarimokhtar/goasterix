package transform

import (
	"encoding/json"
	"encoding/xml"
	"github.com/mokhtarimokhtar/goasterix/uap"
)

type Writer interface {
	write([]uap.DataField)
}

func WriteModel(w Writer, items []uap.DataField) {
	w.write(items)
}
func WriteModelJSON(w Writer, items []uap.DataField) (j []byte, err error) {
	w.write(items)
	j, err = json.Marshal(w)
	return j, err
}
func WriteModelXML(w Writer, items []uap.DataField) (x []byte, err error) {
	w.write(items)
	x, err = xml.Marshal(w)
	return x, err
}
