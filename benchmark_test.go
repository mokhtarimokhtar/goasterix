package goasterix

import (
	"bytes"
	"github.com/mokhtarimokhtar/goasterix/_uap"
	"github.com/mokhtarimokhtar/goasterix/util"
	"testing"
)

func BenchmarkFspecReader(b *testing.B) {
	input := []byte{0xFF, 0x01, 0xF2}
	rb := bytes.NewReader(input)

	for n := 0; n < b.N; n++ {
		_, _ = FspecReader(rb)
	}
}

func BenchmarkFspecIndex(b *testing.B) {
	input := []byte{0xef, 0x98}

	for n := 0; n < b.N; n++ {
		_ = FspecIndex(input)
	}
}

func BenchmarkSubItemBitReader(b *testing.B) {
	input := []byte{0x10, 0x00, 0x00, 0x00}
	for n := 0; n < b.N; n++ {
		sub := new(SubItemBit)
		sub.Name = "item1"
		sub.Type = _uap.Bit
		sub.Pos = 29
		_ = sub.Reader(input)
	}
}

func BenchmarkGetBitsFromTo(b *testing.B) {
	var input = []byte{0xdd, 0x75}
	var from = uint8(13)
	var to = uint8(3)
	for n := 0; n < b.N; n++ {
		_ = GetBitsFromTo(input, from, to)
	}
}

// BitReader
func BenchmarkOneBitReader(b *testing.B) {
	var input byte = 0xd5
	var pos = uint8(6)

	for n := 0; n < b.N; n++ {
		_ = OneBitReader(input, pos)
	}
}

func BenchmarkFromToBitReader8(b *testing.B) {
	var input byte = 0xd5
	var from = uint8(7)
	var to = uint8(2)
	for n := 0; n < b.N; n++ {
		_ = FromToBitReader8(input, from, to)
	}
}

func BenchmarkFromToBitReader16(b *testing.B) {
	//var input uint16 = 0xdd75
	var input = []byte{0xdd, 0x75}
	var from = uint8(13)
	var to = uint8(3)
	for n := 0; n < b.N; n++ {
		_ = FromToBitReader16(input, from, to)
	}
}

func BenchmarkFromToBitReader32(b *testing.B) {
	var input uint32 = 0xdd75cc33
	var from = uint8(13)
	var to = uint8(3)
	for n := 0; n < b.N; n++ {
		_ = FromToBitReader32(input, from, to)
	}
}

func benchmarkRecordDecode(input string, uap StandardUAP, b *testing.B) {
	data, _ := util.HexStringToByte(input)
	for n := 0; n < b.N; n++ {
		rec := new(Record)
		_, _ = rec.Decode(data, uap)
	}
}

func BenchmarkRecordDecode_Len43(b *testing.B) {
	benchmarkRecordDecode(
		"f780 ffff 01 0302 0801020304050607 03aaaaaabbbbbbcccccc  b80101010202aaaabbbb0201 0201 04010203",
		CatForTest,
		b)
}

func BenchmarkRecordDecode_Len55(b *testing.B) {
	benchmarkRecordDecode(
		"fff702 0836 429b52 a0 94c70181 0913 02d0 6002b7 490d01 38a178cf4220 02e79a5d27a00c0060a3280030a4000040 063a 0743ce5b 40 20f5",
		Cat048V127,
		b)
}

/*
func benchmarkRecordDecode(input string, items _uap.StandardUAP, b *testing.B) {
	data, _ := util.HexStringToByte(input)
	for n := 0; n < b.N; n++ {
		rec := new(Record)
		unRead, err := rec.Decode(data, items)
		if err != nil {
			b.Errorf("MsgFailInValue: error = %v; Expected: %v", err, nil)
		}
		if unRead != 0 {
			b.Errorf("MsgFailInValue: unRead = %v; Expected: %v", unRead, 0)
		}
	}
}

// benchmark some records
func BenchmarkRecordDecode_Len7(b *testing.B) {
	benchmarkRecordDecode(
		"e008837dfd9c58",
		_uap.Cat255StrV51,
		b)
}
func BenchmarkRecordDecode_Len9(b *testing.B) {
	benchmarkRecordDecode(
		"f4083902105fb35b02",
		_uap.Cat002V10,
		b)
}
func BenchmarkRecordDecode_Len16(b *testing.B) {
	benchmarkRecordDecode(
		"f50208319801bf0a1ebb43022538e200",
		_uap.Cat001V12,
		b)
}
func BenchmarkRecordDecode_Len17(b *testing.B) {
	benchmarkRecordDecode(
		"f6083602429b7110940028200094008000",
		_uap.Cat034V127,
		b)
}
func BenchmarkRecordDecode_Len21(b *testing.B) {
	benchmarkRecordDecode(
		"d008843b549400130000008f002f008948006a007c",
		_uap.Cat032StrV70,
		b)
}
func BenchmarkRecordDecode_Len55(b *testing.B) {
	benchmarkRecordDecode(
		"fff702 0836 429b52 a0 94c70181 0913 02d0 6002b7 490d01 38a178cf4220 02e79a5d27a00c0060a3280030a4000040 063a 0743ce5b 40 20f5",
		_uap.Cat048V127,
		b)
}
*/

/*func BenchmarkRecordDecode_Len68(b *testing.B) {
	benchmarkRecordDecode(
		"fc ffff fffffe 03ffff 02ffffffff ab80 ff fffe 02ffffffff 04ffffff ffff 0101ffff",
		_uap.CatForTest,
		b)
}*/
/*func BenchmarkRecordDecode_Len73(b *testing.B) {
	benchmarkRecordDecode(
		"afbbf317f1300883040070a8bcf3ff07070723f0a8800713feb7022b0389038b140704012c080811580000001e7004f04aa004b0012400544e49413531313206c84c45424c48454c58",
		_uap.Cat030ArtasV70,
		b)
}*/

func benchmarkDataBlockDecode(input string, b *testing.B) {
	data, _ := util.HexStringToByte(input)
	for n := 0; n < b.N; n++ {
		dataB := NewDataBlock()
		unRead, err := dataB.Decode(data)

		if err != nil {
			b.Errorf("MsgFailInValue: error = %v; Expected: %v", err, nil)
		}
		if unRead != 0 {
			b.Errorf("MsgFailInValue: unRead = %v; Expected: %v", unRead, 0)
		}
	}
}

// benchmark one cat048 datablock
func BenchmarkDataBlock_Len280(b *testing.B) {
	benchmarkDataBlockDecode(
		"300118fff7020836429b52a094c70181091302d06002b7490d0138a178cf422002e79a5d27a00c0060a3280030a4000040063a0743ce5b4020f5fff7020836429b54e000bc020901a2005c7802e800263946e50464b1cb6ca0029ea9491062a4546093880032d4000040059602f639590220f5fff7020836429b58a0909703ff026405a26002bb4066740815f6e795e002e56a0530ffdff860b0d80032fc00004003cf0810c9ef4020fdfff7020836429b56a0775d03700ec205786002be4060910815f9c363a002a49a0f30bfffff60c4600030a4000040057207674a004020fdfff7020836429b55a0468c029804b105786002c57101124d6070d3282002adfa3333a0140060c4600030a4000040026e07d75fc04020f5",
		b)
}

func benchmarkWrapperDataBlockDecode(input string, b *testing.B) {
	data, _ := util.HexStringToByte(input)
	for n := 0; n < b.N; n++ {
		dataB, _ := NewWrapperDataBlock()
		_, _ = dataB.Decode(data)
	}
}

func BenchmarkWapperDataBlock_Len768(b *testing.B) {
	benchmarkWrapperDataBlockDecode(
		"300180fff70208364eadc8a2a44411850fff07a86002c5382fdb4cd4f240e8200100000000000000e10004000cd3bd4022a0fff70208364eadc8a2544411940fff07946001cb382fbb4cd4f140e8200100000000000000e10005001d32884022a0fff70208364eadd0a03d09158f045605c86002c94853d4512071d3706002c919ff3160140060c8480030a800004002ea07e392944022f5ffd70208364eadcfa0accc153d058304386002b744f1a20811b2e3282006810856feb7402aa0fff70208364eadc7a07420113c045a06016002c24853d2512073cca82002c839ef3161542960d0180030a800004005a007da911b4022f5fff70208364eadcca07fff1371056305ef6002bf43ec3ec931d31e082002ea99f331201c0160ca3c0130a800004003e30804d2f74022f5ff1608364eadd26007ba15b80e000038f84c07d43d4600cb0173530e00fff70208364eadc5a03e95104105e606406002c84ca97c4994b710582002eff9d13020240060ce267130a800004002ae07c3dfc64022fd",
		b)
}
