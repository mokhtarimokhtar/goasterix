package goasterix

import (
	"github.com/mokhtarimokhtar/goasterix/uap"
	"testing"
)

func benchmarkRecordDecode(input string, items uap.StandardUAP, b *testing.B) {
	data := HexStringToByte(input)
	for n := 0; n < b.N; n++ {
		rec := new(Record)
		unRead, err := rec.Decode(data, items)
		if err != nil {
			b.Errorf("FAIL: error = %v; Expected: %v", err, nil)
		}
		if unRead != 0 {
			b.Errorf("FAIL: unRead = %v; Expected: %v", unRead, 0)
		}
	}
}

// benchmark some records
func BenchmarkRecordDecode_Len7(b *testing.B) {
	benchmarkRecordDecode(
		"e008837dfd9c58",
		uap.Cat255StrV51,
		b)
}
func BenchmarkRecordDecode_Len9(b *testing.B) {
	benchmarkRecordDecode(
		"f4083902105fb35b02",
		uap.Cat002V10,
		b)
}
func BenchmarkRecordDecode_Len16(b *testing.B) {
	benchmarkRecordDecode(
		"f50208319801bf0a1ebb43022538e200",
		uap.Cat001V12,
		b)
}
func BenchmarkRecordDecode_Len17(b *testing.B) {
	benchmarkRecordDecode(
		"f6083602429b7110940028200094008000",
		uap.Cat034V127,
		b)
}
func BenchmarkRecordDecode_Len21(b *testing.B) {
	benchmarkRecordDecode(
		"d008843b549400130000008f002f008948006a007c",
		uap.Cat032StrV70,
		b)
}
func BenchmarkRecordDecode_Len55(b *testing.B) {
	benchmarkRecordDecode(
		"fff702 0836 429b52 a0 94c70181 0913 02d0 6002b7 490d01 38a178cf4220 02e79a5d27a00c0060a3280030a4000040 063a 0743ce5b 40 20f5",
		uap.Cat048V127,
		b)
}
func BenchmarkRecordDecode_Len73(b *testing.B) {
	benchmarkRecordDecode(
		"afbbf317f1300883040070a8bcf3ff07070723f0a8800713feb7022b0389038b140704012c080811580000001e7004f04aa004b0012400544e49413531313206c84c45424c48454c58",
		uap.Cat030ArtasV70,
		b)
}

func benchmarkDataBlockDecode(input string, b *testing.B) {
	data := HexStringToByte(input)
	for n := 0; n < b.N; n++ {
		dataB, _ := NewDataBlock()
		unRead, err := dataB.Decode(data)

		if err != nil {
			b.Errorf("FAIL: error = %v; Expected: %v", err, nil)
		}
		if unRead != 0 {
			b.Errorf("FAIL: unRead = %v; Expected: %v", unRead, 0)
		}
	}
}

// benchmark one cat048 datablock
func BenchmarkDataBlock_Len280(b *testing.B) {
	benchmarkDataBlockDecode(
		"300118fff7020836429b52a094c70181091302d06002b7490d0138a178cf422002e79a5d27a00c0060a3280030a4000040063a0743ce5b4020f5fff7020836429b54e000bc020901a2005c7802e800263946e50464b1cb6ca0029ea9491062a4546093880032d4000040059602f639590220f5fff7020836429b58a0909703ff026405a26002bb4066740815f6e795e002e56a0530ffdff860b0d80032fc00004003cf0810c9ef4020fdfff7020836429b56a0775d03700ec205786002be4060910815f9c363a002a49a0f30bfffff60c4600030a4000040057207674a004020fdfff7020836429b55a0468c029804b105786002c57101124d6070d3282002adfa3333a0140060c4600030a4000040026e07d75fc04020f5",
		b)
}
