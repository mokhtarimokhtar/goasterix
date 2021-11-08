package goasterix

import (
	"bytes"
	"fmt"
	"github.com/mokhtarimokhtar/goasterix/uap"
	"io"
	"testing"
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
	dataSetWrapperDataBlockTests := []WrapperDataBlockTest{
		/*{
			// CAT048 + CAT034
			input:          "300180fff70208364eadc8a2a44411850fff07a86002c5382fdb4cd4f240e8200100000000000000e10004000cd3bd4022a0fff70208364eadc8a2544411940fff07946001cb382fbb4cd4f140e8200100000000000000e10005001d32884022a0fff70208364eadd0a03d09158f045605c86002c94853d4512071d3706002c919ff3160140060c8480030a800004002ea07e392944022f5ffd70208364eadcfa0accc153d058304386002b744f1a20811b2e3282006810856feb7402aa0fff70208364eadc7a07420113c045a06016002c24853d2512073cca82002c839ef3161542960d0180030a800004005a007da911b4022f5fff70208364eadcca07fff1371056305ef6002bf43ec3ec931d31e082002ea99f331201c0160ca3c0130a800004003e30804d2f74022f5ff1608364eadd26007ba15b80e000038f84c07d43d4600cb0173530e00fff70208364eadc5a03e95104105e606406002c84ca97c4994b710582002eff9d13020240060ce267130a800004002ae07c3dfc64022fd 220014f60836024eadd618940028200094008000",
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
		},*/
		{
			// CAT002 + CAT001
			input:          "02000cf4083902105fb35b02010076f502083990002018aa134c06db08f000750290003b1595114104f15f470075029001b407b5115401ca25dc0075029000f212e514b2067fdc230075029000f508271713020875c90075029001050b78178e02743f7c007502900114090417b80190e4ca00750690010412e915bf079cba112000",
			err:            nil,
			nbOfDataBlocks: 2,
			unRead:         0,
		},
	}

	for _, row := range dataSetWrapperDataBlockTests {
		// Arrange
		data := HexStringToByte(row.input)
		w := new(WrapperDataBlock)

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

/**
DataBlock Testing
a DataBlock correspond to one (only) category and contains one or more Records.
DataBlock = CAT + LEN + [FSPEC + items...] + [...] + ...
*/
type DataBlockTest struct {
	input       string // data block test input
	nbOfRecords int
	err         error // error expected
	unRead      int   // the number of bytes not read.
}

func TestDataBlockDecode(t *testing.T) {
	// setup
	dataSet := []DataBlockTest{
		{
			// empty
			input:       "",
			err:         io.EOF,
			nbOfRecords: 0,
			unRead:      0,
		},
		{
			// CAT048
			input:       "30",
			err:         io.EOF,
			nbOfRecords: 0,
			unRead:      0,
		},
		{
			// CAT048
			input:       "300118fff7020836429b52a094c70181091302d06002b7490d0138a178cf422002e79a5d27a00c0060a3280030a4000040063a0743ce5b4020f5fff7020836429b54e000bc020901a2005c7802e800263946e50464b1cb6ca0029ea9491062a4546093880032d4000040059602f639590220f5fff7020836429b58a0909703ff026405a26002bb4066740815f6e795e002e56a0530ffdff860b0d80032fc00004003cf0810c9ef4020fdfff7020836429b56a0775d03700ec205786002be4060910815f9c363a002a49a0f30bfffff60c4600030a4000040057207674a004020fdfff7020836429b55a0468c029804b105786002c57101124d6070d3282002adfa3333a0140060c4600030a4000040026e07d75fc04020f5",
			err:         nil,
			nbOfRecords: 5,
			unRead:      0,
		},
		{
			// CAT048
			input:       "300142ffd7020836429dc6e0235f3b2201b003087802c726373986e60464b1cf63a000ec06963aea0220f4fff7020836429dc4a057e839ad01ae047c6003c43991ea0464b1c34d2002997a5d30a0ec1960c2680030a4000040044d0701334a4020f5fff7020836429dcfa090163f5d020005ee6002bd4ca75905a071cb08200280f9e9302037ff60ca380032fc0000400438078afa6d4020f5fff7020836429dc3a04d9a397e020004476002c43991e80464b3e541e00297da7130e16c1e60c2680030a4000040076707202b554020f5fff7020836429dc3e007473976020000eb7802dd3d433985660464b4c50560029fa9f31a62845060a5252930a8018740054004c43c5d0220f5fff7020836429dc6e034883b28020003147802c32632394c0c0464b4c93620029e0a7d2ae17c2b60b2c80032d4000040028f071a3bb20020f5220014f6083602429dd140940028200094008000",
			err:         nil,
			nbOfRecords: 6,
			unRead:      20, // bytes left of other category
		},
		{
			// CAT048 // in end a byte 0xf5 has been removed
			input:       "300142ffd7020836429dc6e0235f3b2201b003087802c726373986e60464b1cf63a000ec06963aea0220f4fff7020836429dc4a057e839ad01ae047c6003c43991ea0464b1c34d2002997a5d30a0ec1960c2680030a4000040044d0701334a4020f5fff7020836429dcfa090163f5d020005ee6002bd4ca75905a071cb08200280f9e9302037ff60ca380032fc0000400438078afa6d4020f5fff7020836429dc3a04d9a397e020004476002c43991e80464b3e541e00297da7130e16c1e60c2680030a4000040076707202b554020f5fff7020836429dc3e007473976020000eb7802dd3d433985660464b4c50560029fa9f31a62845060a5252930a8018740054004c43c5d0220f5fff7020836429dc6e034883b28020003147802c32632394c0c0464b4c93620029e0a7d2ae17c2b60b2c80032d4000040028f071a3bb20020",
			err:         ErrUndersized,
			nbOfRecords: 0,
			unRead:      318, // bytes left of other category
		},
		{
			// CAT034
			input:       "220014f6083602429b7110940028200094008000",
			err:         nil,
			nbOfRecords: 1,
			unRead:      0,
		},
		{
			// CAT034
			input:       "22",
			err:         io.EOF,
			nbOfRecords: 0,
			unRead:      0,
		},
		{
			// CAT034
			input:       "220014f6083602429b71109400282000940080", // in end a byte 0x00 has been removed
			err:         ErrUndersized,
			nbOfRecords: 0,
			unRead:      16,
		},
		{
			// CAT030 STR
			input:       "1e009fbffb0160088358052c7dfc04010e0fe86601c4720e008c008c01beff8bf027190439cc821885050e08203fff01605800847dfc04010e0a6968a7d6160e029d02a2fc660498f8feb917010c4caa2358f171dc15603ffb01605801d27dfc04010e0b1a6d60cf860e02d002d0fd460370f017010c4d02a6286076d518203ffb805805387dfc040f0e0e007593ccb20e00500050feb9ff5df017010c2205",
			err:         nil,
			nbOfRecords: 4,
			unRead:      0,
		},
		{
			// CAT255 STR
			input:       "ff000ae008837dfb9c58",
			err:         nil,
			nbOfRecords: 1,
			unRead:      0,
		},
	}

	for _, row := range dataSet {
		// Arrange
		data := HexStringToByte(row.input)
		dataB := new(DataBlock)

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

func TestDataBlockPayload(t *testing.T) {
	// Arrange
	data := HexStringToByte("300118fff7020836429b52a094c70181091302d06002b7490d0138a178cf422002e79a5d27a00c0060a3280030a4000040063a0743ce5b4020f5fff7020836429b54e000bc020901a2005c7802e800263946e50464b1cb6ca0029ea9491062a4546093880032d4000040059602f639590220f5fff7020836429b58a0909703ff026405a26002bb4066740815f6e795e002e56a0530ffdff860b0d80032fc00004003cf0810c9ef4020fdfff7020836429b56a0775d03700ec205786002be4060910815f9c363a002a49a0f30bfffff60c4600030a4000040057207674a004020fdfff7020836429b55a0468c029804b105786002c57101124d6070d3282002adfa3333a0140060c4600030a4000040026e07d75fc04020f5")
	nbOfRecords := 5
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

func TestDataBlockString(t *testing.T) {
	// Arrange
	data := HexStringToByte("300118fff7020836429b52a094c70181091302d06002b7490d0138a178cf422002e79a5d27a00c0060a3280030a4000040063a0743ce5b4020f5fff7020836429b54e000bc020901a2005c7802e800263946e50464b1cb6ca0029ea9491062a4546093880032d4000040059602f639590220f5fff7020836429b58a0909703ff026405a26002bb4066740815f6e795e002e56a0530ffdff860b0d80032fc00004003cf0810c9ef4020fdfff7020836429b56a0775d03700ec205786002be4060910815f9c363a002a49a0f30bfffff60c4600030a4000040057207674a004020fdfff7020836429b55a0468c029804b105786002c57101124d6070d3282002adfa3333a0140060c4600030a4000040026e07d75fc04020f5")
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
	uap.Profiles[30] = uap.Cat030ArtasV62

	for _, row := range dataSet {
		// Arrange
		data := HexStringToByte(row.input)
		dataB, _ := NewDataBlock()

		fmt.Println(len(data))

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

func TestHexStringToByte_Valid(t *testing.T) {
	// Arrange
	input := "01 0203 04"
	output := []byte{0x01, 0x02, 0x03, 0x04}

	// Act
	data := HexStringToByte(input)

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
	data := HexStringToByte(input)

	// Assert
	if bytes.Equal(data, output) == false {
		t.Errorf("FAIL: data = % X; Expected: % X", data, output)
	} else {
		t.Logf("SUCCESS: data = % X; Expected: % X", data, output)
	}
}



