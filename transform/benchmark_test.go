package transform

import (
	"github.com/mokhtarimokhtar/goasterix"
	"github.com/mokhtarimokhtar/goasterix/uap"
	"github.com/mokhtarimokhtar/goasterix/util"
	"testing"
)

func benchmarkWriteModel(input string, items uap.StandardUAP, b *testing.B) {
	data, _ := util.HexStringToByte(input)
	//rec := goasterix.NewRecord()
	//rec := new(goasterix.Record)
	//unRead, err := rec.Decode(data, items)
	//if err != nil {
	//	b.Errorf("FAIL: error = %v; Expected: %v", err, nil)
	//}
	//if unRead != 0 {
	//	b.Errorf("FAIL: unRead = %v; Expected: %v", unRead, 0)
	//}

	for n := 0; n < b.N; n++ {
		rec := new(goasterix.Record)
		_, _ = rec.Decode(data, items)
		cat048Model := new(Cat048Model)
		WriteModel(cat048Model, *rec)
	}
}

func BenchmarkWriteModel_Len59(b *testing.B) {
	benchmarkWriteModel(
		"ffff020836429b52a094c70181091302d06002b7490d0138a178cf422002e79a5d27a00c0060a3280030a4000040063a008000800743ce5b4020f5",
		uap.Cat048V127,
		b)
}
