package model

import (
	"bytes"
	"github.com/mokhtarimokhtar/goasterix"
	"github.com/mokhtarimokhtar/goasterix/item"
	"github.com/mokhtarimokhtar/goasterix/util"
	"testing"
)

func benchmarkWriteModel(input string, uap item.UAP, b *testing.B) {
	data, _ := util.HexStringToByte(input)
	rb := bytes.NewReader(data)
	//rec := goasterix.NewRecord()
	rec := new(goasterix.Record)
	_, _ = rec.Decode(rb, uap)

	for n := 0; n < b.N; n++ {
		model := new(Cat048Model)
		model.write(*rec)
	}
}

func BenchmarkWriteModel_Len59(b *testing.B) {
	benchmarkWriteModel(
		"ffff020836429b52a094c70181091302d06002b7490d0138a178cf422002e79a5d27a00c0060a3280030a4000040063a008000800743ce5b4020f5",
		goasterix.Cat048V127,
		b)
}
