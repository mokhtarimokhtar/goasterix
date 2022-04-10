package goasterix

import (
	"bytes"
	"encoding/hex"
	"io"
	"testing"

	"github.com/mokhtarimokhtar/goasterix/uap"
)

/**
Wrapper DataBlock Testing
a Wrapper DataBlock correspond to one or more category and contains one or more Records.
DataBlock = CAT + LEN + [FSPEC + items...] + [...] + ...
WrapperDataBlock = [CAT + LEN + RECORD + ...] + [ DATABLOCK ] + [...]
*/
type WrapperDataBlockTest struct {
	input          string // data block test input
	nbOfDataBlocks int
	err            error // error expected
	unRead         int   // the number of bytes not read.
}

func TestWrapperDataBlockDecode(t *testing.T) {
	// setup
	dataSet := []WrapperDataBlockTest{
		{
			// CAT048 + CAT034
			input:          "300180fff70208364eadc8a2a44411850fff07a86002c5382fdb4cd4f240e8200100000000000000e10004000cd3bd4022a0fff70208364eadc8a2544411940fff07946001cb382fbb4cd4f140e8200100000000000000e10005001d32884022a0fff70208364eadd0a03d09158f045605c86002c94853d4512071d3706002c919ff3160140060c8480030a800004002ea07e392944022f5ffd70208364eadcfa0accc153d058304386002b744f1a20811b2e3282006810856feb7402aa0fff70208364eadc7a07420113c045a06016002c24853d2512073cca82002c839ef3161542960d0180030a800004005a007da911b4022f5fff70208364eadcca07fff1371056305ef6002bf43ec3ec931d31e082002ea99f331201c0160ca3c0130a800004003e30804d2f74022f5ff1608364eadd26007ba15b80e000038f84c07d43d4600cb0173530e00fff70208364eadc5a03e95104105e606406002c84ca97c4994b710582002eff9d13020240060ce267130a800004002ae07c3dfc64022fd 220014f60836024eadd618940028200094008000",
			err:            nil,
			nbOfDataBlocks: 2,
			unRead:         0,
		},
		{
			// CAT062 + CAT065
			input:          "3e02bcbf5ffd0304090c03424cf3008582050001c5fa0172e1ffca36012702c404c5e301010c345691242173cb0da00d2005a0010600640cb0190301487fa0ff0808ffffffffffff0893110101c008081212121205b3107d05b2ffa7dc08022b1b867f05b304c5a01ba004083605054508110500a608090506a608020501d4012102c5bf5ffd0304090c03424cf30084dd100017504b131b670003b1028a02490ae5c10101003c308f1025880608200027190301487fa0ff0808ffffffffffff0891010101000808053d1589053bff14dc08048f192ab2053d0ae5a00ca00108040504da024a028abf5ffd0304090c03424cf300856f110008281306aa7cffd158009701600c01c10101003b9bab5891f2d208200a7b190301487fa0ff0c0cffffffffffff0891010101000c0c00f401fc00f4ff8098081100f40c0111a00208020500480811050121008a0166bf4ffd0304090c03424cf30084f719000f623c0c9b5fff924cfc0dff6d01000289190301587fa0ff0bffffffffffffff0a91010101000b0b00c8187b00bff98a98080480c8010011a00208120203f808040204b0fc19ff28bf5ffd0304090c03424cf300848b4c0014877110dd24ff68f1feeb0029096ac10101004b2a3720241a4a08200d6b990301487fa0ff1414ffffffffffff0091010101001414012c00a0012c0000980812012c096a0620fee8000fbf4ffd0304090c03424cf30084f2d1001493ea10dbfeffe50efecfffd10c4005ac990301587fa0ff14ffffffffffffff149101010100141400ec00ac00ef007f98081200ec0c400ca0010812020659fed4ffb4bf5ffd0304090c03424cf30085a192000029f200223dffee62ffae00f30e18c1010100392a001872d1060820090e190301487fa0ff0606ffffffffffff0891010101001306005400a00054fff2dc080231509a8200540e188011a00208110504700802040450ffad00f241000cf8090c0203424cf30a",
			err:            nil,
			nbOfDataBlocks: 2,
			unRead:         0,
		},
		{
			// CAT030 STR + CAT255 STR
			input:          "1e015bbfff8160088358009c7dfb27090e0e00450cfcd30e009b009a0175003df4003d27110428214b1a972022c25a08203fff81605806ac7dfb27090e0e0042fef8260e008e0090feb5ffddf8ff9a2711042821384ffc18a18142082037ff7f605806707dfb2702004e9ecf4f0e00510052fe210439f8ff52270904584b313036424901dc415437354d4c464b4a4c464d4e3945f760bc70d8226037ff7f605804a07dfb270dfd49b4ecf40e062d062c022b0568f400cc2707c44a464131394d2004ec433235424c4c464d444544464d4d02bc286071e4d8203fff81605804927dfb27090e0e00482bff260e00d000cdfe21fecbf4008527011c28214b1b992022cc5208203ffb816058047e7dfb27090e0e004c5107a20e00cc00ccfe0a0048f027110428214b1ba52022cd1a08203fff81605802b27dfb27090e0e004aa104c00e002e003000ac00aef8ffaf27110428214b1bab2022cd320820 ff000ae008837e019d58",
			err:            nil,
			nbOfDataBlocks: 2,
			unRead:         0,
		},
		{
			// CAT002 + CAT001
			input:          "02000cf4083902105fb35b02010076f502083990002018aa134c06db08f000750290003b1595114104f15f470075029001b407b5115401ca25dc0075029000f212e514b2067fdc230075029000f508271713020875c90075029001050b78178e02743f7c007502900114090417b80190e4ca00750690010412e915bf079cba112000",
			err:            nil,
			nbOfDataBlocks: 2,
			unRead:         0,
		},
		{
			// empty
			input:          "",
			err:            io.EOF,
			nbOfDataBlocks: 0,
			unRead:         0,
		},
	}

	for _, row := range dataSet {
		// Arrange
		data, _ := HexStringToByte(row.input)
		w, _ := NewWrapperDataBlock()

		// Act
		unRead, err := w.Decode(data)

		// Assert
		if err != row.err {
			t.Errorf("FAIL: error: %s; Expected: %v", err, row.err)
		} else {
			t.Logf("SUCCESS: error: %v; Expected: %v", err, row.err)
		}
		if unRead != row.unRead {
			t.Errorf("FAIL: unRead = %v; Expected: %v", unRead, row.unRead)
		} else {
			t.Logf("SUCCESS: unRead = %v; Expected: %v", unRead, row.unRead)
		}
		if len(w.DataBlocks) != row.nbOfDataBlocks {
			t.Errorf("FAIL: nbOfDataBlocks = %v; Expected: %v", len(w.DataBlocks), row.nbOfDataBlocks)
		} else {
			t.Logf("SUCCESS: nbOfDataBlocks = %v; Expected: %v", len(w.DataBlocks), row.nbOfDataBlocks)
		}
	}
}

// DataBlock Testing
// a DataBlock correspond to one (only) category and contains one or more Records.
// DataBlock = CAT + LEN + [FSPEC + items...] + [...] + ...
type DataBlockTest struct {
	TestCaseName string
	input        string // data block test input
	nbOfRecords  int
	err          error // error expected
	unRead       int   // the number of bytes not read.
}

func TestDataBlockDecode(t *testing.T) {
	// setup
	dataSet := []DataBlockTest{
		{
			TestCaseName: "empty data block",
			input:        "",
			err:          io.EOF,
			nbOfRecords:  0,
			unRead:       0,
		},
		{
			TestCaseName: "CAT048 only category field",
			input:        "30",
			err:          io.EOF,
			nbOfRecords:  0,
			unRead:       0,
		},
		{
			TestCaseName: "CAT062",
			input:        "3e0547bf5ffd0304090001532100008e6f3e0017d0961247f10b7086fed3019a0fc8e301010c87304a04e072c34820e300820800eb003104b2190301487fa0ff0614ffffffffffff0493110101c006061414141400e0045b00e00182dc622931a410a800e00fc84010e001622b05010d01622902fea60177bf5ffd0304090001532100008f45be000478e9036aa20b78f8fdbc023c0f55e301010c40123f0815f5cf1820dee002d0010f005f002c190301087fa02a0707ffffffffffff0893110101c0070707070707051f13c5051bfdd4dc085066b0f616051f0f55a02aa0070814050221060b0500b108360502ea0813050761060a05056808090504340850050236fdb10230bf5ffd0304090001532100008ea9d100149a720fcb720b75af033a014d0baae301010c780de50c54f7c39e202bc003c0012b00560336190301087fa00e0606ffffffffffff0493110101c0060606060606039715f00399013e98060a03970baa1fe00408120501f6080a05065f06030701f4060a0503b10162290203180199bf5ffd0304090001532100008f5d7f0000d79400a4960b8ca4026cfd4b0ec7e301010c4009c70815f4e0356064c00578011d00570321190301487fa0ff0808ffffffffffff0493110101c0080808080808040b16b9040d00bddc08097dd1e0a6040b0ec7a025a006060b0506e5081305075c08140506c3083605038408500507e6080905051c026ffd4fbf5ffd0304090001532100008e74b40018ecd31320cf0b8f89fe8bff780684e301010c4bcdee4d84f7cc3820af0000c800bd002601e2190301487fa0ff1313ffffffffffff0893110101c01313131313130067023f0065ff4d98622b006726840ca001622b05007efe9cff4dbf5ffd0304090001532100008e9f9d00172dfb11c53c0b9c93fd09014d064be301010c4bc846407532dd7820d3600550010b006001e5190301487fa0ff060bffffffffffff0093110101c006060b0b0b0b0550149f05500000dc6229330c09100550064b4015e002622b050901080a05018701622902fce900fcbf4ffd0304090001532100008e66d600187dc812cdb00b7555ff07ff960991006f990301587fa0ff19ffffffffffffff28910101010019190030009a002eff71986229003009910620ff13ff7abf5ffd0304090001532100008ef1c50010032c0c41460b7f70037efec10200e301010c4ca4f84994b3e774a04a400528011b00590359190301087fa0050404ffffffffffff4493110101c0040404040404044d1726044f013b980814044d020033e00808120505a5080a0500bb08360504bc060b0503190603070760060a07071d0809050527081405044a016229020393ff01bf5ffd0304090001532100008ef000001110340d0f160b8d0703dbffaa0200e301010c4ca8af4994b8e4d320404005280118005f0277190301087fa00c0404ffffffffffff0493110101c004040404040404d117fe04d200e698081404d102002ee00708120506ce080a0500ce08360504d106030706dc060a0504f908090502bb08140500fe0162290203dffff7bf5ffd0304090001532100008f19fe000bb36808f2860b7977ff4d00190400c101010044d9c93cf58e2608200146990301087fa0313131ffffffffffff0091010101003131004800a000480000980603004804000620ff4b000fbf5ffd0304090001532100008ef943001611ec10e10c0bf096041afec70ab1e301010c8013c2594270c78820466085c801130069002e190301087fa00f0606ffffffffffff0093110101c006060606060605c7190005c70000dc62293988fcf005c70ab14024e005622b05050d08120506ad080a0507160603070756060a0504e5016229020435ff30",
			err:          nil,
			nbOfRecords:  11,
			unRead:       0,
		},
		{
			TestCaseName: "CAT065",
			input:        "41000cf8090c0203424d730f",
			err:          nil,
			nbOfRecords:  1,
			unRead:       0,
		},
		{
			TestCaseName: "CAT048",
			input:        "300118fff7020836429b52a094c70181091302d06002b7490d0138a178cf422002e79a5d27a00c0060a3280030a4000040063a0743ce5b4020f5fff7020836429b54e000bc020901a2005c7802e800263946e50464b1cb6ca0029ea9491062a4546093880032d4000040059602f639590220f5fff7020836429b58a0909703ff026405a26002bb4066740815f6e795e002e56a0530ffdff860b0d80032fc00004003cf0810c9ef4020fdfff7020836429b56a0775d03700ec205786002be4060910815f9c363a002a49a0f30bfffff60c4600030a4000040057207674a004020fdfff7020836429b55a0468c029804b105786002c57101124d6070d3282002adfa3333a0140060c4600030a4000040026e07d75fc04020f5",
			err:          nil,
			nbOfRecords:  5,
			unRead:       0,
		},
		{
			TestCaseName: "CAT048",
			input:        "300142ffd7020836429dc6e0235f3b2201b003087802c726373986e60464b1cf63a000ec06963aea0220f4fff7020836429dc4a057e839ad01ae047c6003c43991ea0464b1c34d2002997a5d30a0ec1960c2680030a4000040044d0701334a4020f5fff7020836429dcfa090163f5d020005ee6002bd4ca75905a071cb08200280f9e9302037ff60ca380032fc0000400438078afa6d4020f5fff7020836429dc3a04d9a397e020004476002c43991e80464b3e541e00297da7130e16c1e60c2680030a4000040076707202b554020f5fff7020836429dc3e007473976020000eb7802dd3d433985660464b4c50560029fa9f31a62845060a5252930a8018740054004c43c5d0220f5fff7020836429dc6e034883b28020003147802c32632394c0c0464b4c93620029e0a7d2ae17c2b60b2c80032d4000040028f071a3bb20020f5220014f6083602429dd140940028200094008000",
			err:          nil,
			nbOfRecords:  6,
			unRead:       20, // bytes left of other category
		},
		{
			TestCaseName: "CAT048 in end a byte 0xf5 has been removed",
			input:        "300142ffd7020836429dc6e0235f3b2201b003087802c726373986e60464b1cf63a000ec06963aea0220f4fff7020836429dc4a057e839ad01ae047c6003c43991ea0464b1c34d2002997a5d30a0ec1960c2680030a4000040044d0701334a4020f5fff7020836429dcfa090163f5d020005ee6002bd4ca75905a071cb08200280f9e9302037ff60ca380032fc0000400438078afa6d4020f5fff7020836429dc3a04d9a397e020004476002c43991e80464b3e541e00297da7130e16c1e60c2680030a4000040076707202b554020f5fff7020836429dc3e007473976020000eb7802dd3d433985660464b4c50560029fa9f31a62845060a5252930a8018740054004c43c5d0220f5fff7020836429dc6e034883b28020003147802c32632394c0c0464b4c93620029e0a7d2ae17c2b60b2c80032d4000040028f071a3bb20020",
			err:          ErrUndersized,
			nbOfRecords:  0,
			unRead:       318, // bytes left of other category
		},
		{
			TestCaseName: "CAT034",
			input:        "220014f6083602429b7110940028200094008000",
			err:          nil,
			nbOfRecords:  1,
			unRead:       0,
		},
		{
			TestCaseName: "CAT034",
			input:        "22",
			err:          io.EOF,
			nbOfRecords:  0,
			unRead:       0,
		},
		{
			TestCaseName: "CAT034",
			input:        "220014f6083602429b71109400282000940080", // in end a byte 0x00 has been removed
			err:          ErrUndersized,
			nbOfRecords:  0,
			unRead:       16,
		},
		{
			TestCaseName: "CAT004",
			input:        "04 003d fdcb8008a2080108826ae1800000080001d1c041504d3030303100010bc51ef7a55900f5050370c30c4000003039ff50ffd8a880404cb3820820",
			err:          nil,
			nbOfRecords:  1,
			unRead:       0,
		},
		{
			TestCaseName: "CAT030 STR",
			input:        "1e009fbffb0160088358052c7dfc04010e0fe86601c4720e008c008c01beff8bf027190439cc821885050e08203fff01605800847dfc04010e0a6968a7d6160e029d02a2fc660498f8feb917010c4caa2358f171dc15603ffb01605801d27dfc04010e0b1a6d60cf860e02d002d0fd460370f017010c4d02a6286076d518203ffb805805387dfc040f0e0e007593ccb20e00500050feb9ff5df017010c2205",
			err:          nil,
			nbOfRecords:  4,
			unRead:       0,
		},
		{
			TestCaseName: "CAT255 STR",
			input:        "ff000ae008837dfb9c58",
			err:          nil,
			nbOfRecords:  1,
			unRead:       0,
		},
		{
			TestCaseName: "Cat4Test full record",
			//
			//input:       "1a 0026 FD80 FFFF FFFE AAFFFFFE 02FFFF FFFF 03FFFF 02FFFFFFFF 04FFFFFF 0101FFFF 04FFFFFF",
			input:       "1a 0029 fd 40 ffff fffffe 03ffff 02ffffffff ab80 ff fffe 02ffffffff 04ffffff ffff 0101ffff 03ffff",
			err:         nil,
			nbOfRecords: 1,
			unRead:      0,
		},
		{
			TestCaseName: "Cat4Test EOF record",
			input:        "1a 0026 FD80 FFFF FFFE AAFFFFFE 02FFFF FFFF 03FFFF 02FFFFFFFF 04FFFFFF 0101FFFF 05FFFFFF",
			err:          io.ErrUnexpectedEOF,
			nbOfRecords:  1,
			unRead:       0,
		},
		{
			TestCaseName: "ErrCategoryUnknown",
			input:        "00 0005FFFF",
			err:          ErrCategoryUnknown,
			nbOfRecords:  0,
			unRead:       0,
		},
		{
			TestCaseName: "CAT063",
			input:        "3f04c9bff0090c7cd2cc08294000000000000000000000000000000000bff0090c7cd2cc08234000000000000000000000000000000000bff0090c7cd2cc14844000000000000000000000000000000000bff0090c7cd2cc22014000000000000000000000000000000000bff0090c7cd2cc220540000000000000ffb70000000000000000bff0090c7cd2cc08214000000000000000000000000000000000bff0090c7cd2cc08084000000000000000000000000000000000bff0090c7cd2cc08044000000000000000000000000000000000bff0090c7cd2cc28214000000000000000000000000000000000bff0090c7cd2cc08264000000000000000000000000000000000bff0090c7cd2cc08054000000000000000000000000000000000bff0090c7cd2cc08114000000000000000000000000000000000bff0090c7cd2cc08174000000000000000000000000000000000bff0090c7cd2cc220c40ff6a0004000400000004000400000000bff0090c7cd2cc08404000000000000000000000000000000000bff0090c7cd2cc08034000000000000000000000000000000000bff0090c7cd2cc14834000000000000000000000000000000000bff0090c7cd2cc08554000000000000000000000000000000000bff0090c7cd2cc14854000000000000000000000000000000000bff0090c7cd2cc08154000000000000000000000000000000000bff0090c7cd2cc08064000000000000000000000000000000000bff0090c7cd2cc08184000000000000000000000000000000000bff0090c7cd2cc08194000000000000000000000000000000000bff0090c7cd2cc14074000000000000000000000000000000000bff0090c7cd2cc08504000000000000000000000000000000000bff0090c7cd2cc08024000000000000000000000000000000000bff0090c7cd2cc140a4000000000000000000000000000000000bff0090c7cd2cc08134000000000000000000000000000000000bff0090c7cd2cc08104000000000000000000000000000000000bff0090c7cd2cc08364000000000000000000000000000000000bff0090c7cd2cc14064000000000000000000000000000000000bff0090c7cd2cc720c4000000000000000000000000000000000bff0090c7cd2cc720e4000000000000000000000000000000000bff0090c7cd2cc08094000000000000000000000000000000000bff0090c7cd2cc08204000000000000000000000000000000000bff0090c7cd2cc060a4000000000000000000000000000000000bff0090c7cd2cc08144000000000000000000000000000000000bff0090c7cd2cc06034000000000000000000000000000000000bff0090c7cd2cc08014000000000000000000000000000000000bff0090c7cd2cc08124000000000000000000000000000000000bff0090c7cd2cc060b4000000000000000000000000000000000bff0090c7cd2cc622b4000000000000000000000000000000000bff0090c7cd2cc281e4000000000000000000000000000000000bff0090c7cd2cc62294000000000000000000000000000000000bff0090c7cd2cc080a4000000000000000000000000000000000bff0090c7cd2cc350840000000000000001b0000000000000000bff0090c7cd2cc353c4000000000000000000000000000000000",
			err:          nil,
			nbOfRecords:  47,
			unRead:       0,
		},
		{
			TestCaseName: "CAT034: over sized data block",
			input:        "220014f6081002412998d89400002000940000811a",
			err:          nil,
			nbOfRecords:  1,
			unRead:       1,
		},
	}

	for _, row := range dataSet {
		// Arrange
		data, _ := HexStringToByte(row.input)
		dataB := NewDataBlock()

		// Act
		unRead, err := dataB.Decode(data)

		// Assert
		if err != row.err {
			t.Errorf("FAIL: %s error: %s; Expected: %v", row.TestCaseName, err, row.err)
		} else {
			t.Logf("SUCCESS: error: %v; Expected: %v", err, row.err)
		}
		if unRead != row.unRead {
			t.Errorf("FAIL: %s unRead = %v; Expected: %v", row.TestCaseName, unRead, row.unRead)
		} else {
			t.Logf("SUCCESS: unRead = %v; Expected: %v", unRead, row.unRead)
		}
		if row.nbOfRecords != len(dataB.Records) {
			t.Errorf("FAIL: %s nbOfRecords = %v; Expected: %v", row.TestCaseName, len(dataB.Records), row.nbOfRecords)
		} else {
			t.Logf("SUCCESS: nbOfRecords = %v; Expected: %v", len(dataB.Records), row.nbOfRecords)
		}
	}
}

func TestDataBlockPayload(t *testing.T) {
	// Arrange
	data, _ := HexStringToByte("30 0118 fff7020836429b52a094c70181091302d06002b7490d0138a178cf422002e79a5d27a00c0060a3280030a4000040063a0743ce5b4020f5fff7020836429b54e000bc020901a2005c7802e800263946e50464b1cb6ca0029ea9491062a4546093880032d4000040059602f639590220f5fff7020836429b58a0909703ff026405a26002bb4066740815f6e795e002e56a0530ffdff860b0d80032fc00004003cf0810c9ef4020fdfff7020836429b56a0775d03700ec205786002be4060910815f9c363a002a49a0f30bfffff60c4600030a4000040057207674a004020fdfff7020836429b55a0468c029804b105786002c57101124d6070d3282002adfa3333a0140060c4600030a4000040026e07d75fc04020f5")
	nbOfRecords := 7 // cat + len + 5 records
	dataB := new(DataBlock)
	_, _ = dataB.Decode(data)

	// Act
	records := dataB.Payload()

	// Assert
	if len(records) != nbOfRecords {
		t.Errorf("FAIL: len(items) = %v; Expected: %v", len(records), nbOfRecords)
	} else {
		t.Logf("SUCCESS: len(items) = %v; Expected: %v", len(records), nbOfRecords)
	}
}

func TestDataBlock_String(t *testing.T) {
	// Arrange
	data, _ := HexStringToByte("300118fff7020836429b52a094c70181091302d06002b7490d0138a178cf422002e79a5d27a00c0060a3280030a4000040063a0743ce5b4020f5fff7020836429b54e000bc020901a2005c7802e800263946e50464b1cb6ca0029ea9491062a4546093880032d4000040059602f639590220f5fff7020836429b58a0909703ff026405a26002bb4066740815f6e795e002e56a0530ffdff860b0d80032fc00004003cf0810c9ef4020fdfff7020836429b56a0775d03700ec205786002be4060910815f9c363a002a49a0f30bfffff60c4600030a4000040057207674a004020fdfff7020836429b55a0468c029804b105786002c57101124d6070d3282002adfa3333a0140060c4600030a4000040026e07d75fc04020f5")
	nbOfRecords := 5
	dataB := new(DataBlock)
	_, _ = dataB.Decode(data)

	// Act
	records := dataB.String()

	// Assert
	if len(records) != nbOfRecords {
		t.Errorf("FAIL: len(items) = %v; Expected: %v", len(records), nbOfRecords)
	} else {
		t.Logf("SUCCESS: len(items) = %v; Expected: %v", len(records), nbOfRecords)
	}
}

func TestDataBlockDecode_ARTAS(t *testing.T) {
	// setup
	dataSet := []DataBlockTest{
		{
			// CAT030 ARTAS
			input:       "1e00f3afbbf317f1300883040070a8bcf3ff07070723f0a8800713feb7022b0389038b140704012c080811580000001e7004f04aa004b0012400544e49413531313206c84c45424c48454c584d413332300101a5389075c71ca0afbbf317f130088304002aa8bcf3ff04040447fda703f7d2008f0df705280528140700000008171158000000087002f0c3c00528012d006955414c3931202007314c4c42474b4557524842373757a290f3541339c60820afbbf31101300883040335a8bcf3ff0b0b0b2be9a9b5fffefffa0fff08c008c01d0e070000001484115800000200700400ffffffffffffffff344045df7df76021d3",
			err:         nil,
			nbOfRecords: 3,
			unRead:      0,
		},
	}
	uap.DefaultProfiles[30] = uap.Cat030ArtasV62

	for _, row := range dataSet {
		// Arrange
		data, _ := HexStringToByte(row.input)
		dataB := NewDataBlock()

		// Act
		unRead, err := dataB.Decode(data)

		// Assert
		if err != row.err {
			t.Errorf("FAIL: error: %s; Expected: %v", err, row.err)
		} else {
			t.Logf("SUCCESS: error: %v; Expected: %v", err, row.err)
		}
		if unRead != row.unRead {
			t.Errorf("FAIL: unRead = %v; Expected: %v", unRead, row.unRead)
		} else {
			t.Logf("SUCCESS: unRead = %v; Expected: %v", unRead, row.unRead)
		}
		if row.nbOfRecords != len(dataB.Records) {
			t.Errorf("FAIL: nbOfRecords = %v; Expected: %v", len(dataB.Records), row.nbOfRecords)
		} else {
			t.Logf("SUCCESS: nbOfRecords = %v; Expected: %v", len(dataB.Records), row.nbOfRecords)
		}
	}
}

func TestTwoComplement16_PositiveNumber(t *testing.T) {
	// Arrange
	input := uint16(0x010F) // 0000 0001 0000 1111
	size := uint8(10)       // ---- --01 0000 1111  -> tenth bit
	output := int16(271)    // 01 0000 1111 = 271

	// Act
	result := TwoComplement16(size, input)

	// Assert
	if result != output {
		t.Errorf("FAIL: result = %v; Expected: %v", result, output)
	} else {
		t.Logf("SUCCESS: result = %v; Expected: %v", result, output)
	}
}

func TestTwoComplement16_NegativeNumber(t *testing.T) {
	// Arrange
	input := uint16(0x040F) // 0000 0100 0000 1111
	size := uint8(11)       // ---- -100 0000 1111  -> tenth bit
	output := int16(-1009)  // ---- -011 1111 0001 = -1009

	// Act
	result := TwoComplement16(size, input)

	// Assert
	if result != output {
		t.Errorf("FAIL: result = %v; Expected: %v", result, output)
	} else {
		t.Logf("SUCCESS: result = %v; Expected: %v", result, output)
	}
}

func TestTwoComplement32_PositiveNumber(t *testing.T) {
	// Arrange
	input := uint32(0x0007EE0F) // 0000 0000 0000 0111 1110 1110 0000 1111
	size := uint8(20)           // 	   ---- ---- ---- 0111 1110 1110 0000 1111  -> twentieth bit
	output := int32(519695)

	// Act
	result := TwoComplement32(size, input)

	// Assert
	if result != output {
		t.Errorf("FAIL: result = %v; Expected: %v", result, output)
	} else {
		t.Logf("SUCCESS: result = %v; Expected: %v", result, output)
	}
}

func TestTwoComplement32_NegativeNumber(t *testing.T) {
	// Arrange
	input := uint32(0x000FEE0F) // 0000 0000 0000 0111 1110 1110 0000 1111
	size := uint8(20)           // 	   ---- ---- ---- 0111 1110 1110 0000 1111  -> twentieth bit
	output := int32(-4593)

	// Act
	result := TwoComplement32(size, input)

	// Assert
	if result != output {
		t.Errorf("FAIL: result = %v; Expected: %v", result, output)
	} else {
		t.Logf("SUCCESS: result = %v; Expected: %v", result, output)
	}
}

func TestHexStringToByte_Valid(t *testing.T) {
	// Arrange
	input := "01 0203 04"
	output := []byte{0x01, 0x02, 0x03, 0x04}

	// Act
	data, _ := HexStringToByte(input)

	// Assert
	if bytes.Equal(data, output) == false {
		t.Errorf("FAIL: data = % X; Expected: % X", data, output)
	} else {
		t.Logf("SUCCESS: data = % X; Expected: % X", data, output)
	}
}

func TestHexStringToByte_Empty(t *testing.T) {
	// Arrange
	input := ""
	var output []byte

	// Act
	data, err := HexStringToByte(input)

	// Assert
	if err != nil {
		t.Errorf("FAIL: error: %s; Expected: %v", err, nil)
	} else {
		t.Logf("SUCCESS: error: %v; Expected: %v", err, nil)
	}
	if bytes.Equal(data, output) == false {
		t.Errorf("FAIL: data = % X; Expected: % X", data, output)
	} else {
		t.Logf("SUCCESS: data = % X; Expected: % X", data, output)
	}
}

func TestHexStringToByte_Error(t *testing.T) {
	// Arrange
	input := "01 0203 0"
	var output []byte

	// Act
	data, err := HexStringToByte(input)

	// Assert
	if err != hex.ErrLength {
		t.Errorf("FAIL: error: %s; Expected: %v", err, hex.ErrLength)
	} else {
		t.Logf("SUCCESS: error: %v; Expected: %v", err, hex.ErrLength)
	}
	if bytes.Equal(data, output) == false {
		t.Errorf("FAIL: % X; Expected: % X", data, output)
	} else {
		t.Logf("SUCCESS: % X; Expected: % X", data, output)
	}
}

