package transform

import (
	"encoding/json"
	"github.com/mokhtarimokhtar/goasterix"
	"github.com/mokhtarimokhtar/goasterix/uap"
	"reflect"
	"testing"
)

func TestCat255STRModel_ToJsonRecord(t *testing.T) {
	// setup
	type dataTest struct {
		TestCaseName string
		input        string
		output       []byte
	}
	dataSet := []dataTest{
		{
			TestCaseName: "Testcase 1",
			input:        "e0 08 83 7dfd9c 58",
			output:       []byte(`{"SourceIdentifier":{"sac":8,"sic":131},"hem":64507.21875,"spe":{"version":2,"nap":3,"ns":"principal"}}`),
		},
	}
	for _, row := range dataSet {
		// Arrange
		uap255 := uap.Cat255StrV51
		data, _ := goasterix.HexStringToByte(row.input)
		rec := new(goasterix.Record)
		_, err := rec.Decode(data, uap255)

		cat255Model := new(Cat255STRModel)
		cat255Model.write(*rec)

		// Act
		recJson, _ := json.Marshal(cat255Model)

		// Assert
		if err != nil {
			t.Errorf("FAIL: error = %v; Expected: %v", err, nil)
		} else {
			t.Logf("SUCCESS: error: %v; Expected: %v", err, nil)
		}

		if reflect.DeepEqual(recJson, row.output) == false {
			t.Errorf("FAIL: %s; \nExpected: %s", recJson, row.output)
		} else {
			t.Logf("SUCCESS: %s; Expected: %s", recJson, row.output)
		}
	}
}

func TestCat255STRModel_SpeStpv(t *testing.T) {
	// setup
	type dataTest struct {
		TestCaseName string
		input        goasterix.Extended
		output       PresenceSTPV
	}
	dataSet := []dataTest{
		{
			TestCaseName: "testcase 1",
			input: goasterix.Extended{
				Primary:   []byte{0x29},
				Secondary: []byte{0x00},
			},
			output: PresenceSTPV{
				Version: 1,
				Nap:     1,
				NS:      "principal",
				ST:      "operational",
				PS:      "stpv_connecte_str",
			},
		},
		{
			TestCaseName: "testcase 2",
			input: goasterix.Extended{
				Primary:   []byte{0x2b},
				Secondary: []byte{0xc0},
			},
			output: PresenceSTPV{
				Version: 1,
				Nap:     1,
				NS:      "secours",
				ST:      "evaluation",
				PS:      "stpv_deconnecte_str",
			},
		},
		{
			TestCaseName: "testcase 3",
			input: goasterix.Extended{
				Primary:   []byte{0x2d},
				Secondary: []byte{0xc0},
			},
			output: PresenceSTPV{
				Version: 1,
				Nap:     1,
				NS:      "test",
				ST:      "evaluation",
				PS:      "stpv_deconnecte_str",
			},
		},
	}

	for _, row := range dataSet {
		// Act
		res := speStpv(row.input)

		// Assert
		if reflect.DeepEqual(res, row.output) == false {
			t.Errorf("FAIL: %s - %v; Expected: %v", row.TestCaseName, res, row.output)
		} else {
			t.Logf("SUCCESS: %v; Expected: %v", res, row.output)
		}
	}
}

func TestCat255STRModel_NivCarte(t *testing.T) {
	// setup
	type dataTest struct {
		TestCaseName string
		input        [4]byte
		output       NivC
	}
	dataSet := []dataTest{
		{
			TestCaseName: "testcase 1",
			input:        [4]byte{0x00, 0x00, 0xff, 0xff},
			output: NivC{
				NivInf: 0,
				NivSup: -1,
			},
		},
	}

	for _, row := range dataSet {
		// Act
		res := nivCarte(row.input)

		// Assert
		if reflect.DeepEqual(res, row.output) == false {
			t.Errorf("FAIL: %s - %v; Expected: %v", row.TestCaseName, res, row.output)
		} else {
			t.Logf("SUCCESS: %v; Expected: %v", res, row.output)
		}
	}
}

func TestCat255STRModel_Carte(t *testing.T) {
	// setup
	type dataTest struct {
		TestCaseName string
		input        [9]byte
		output       CarteActive
		err          error
	}
	dataSet := []dataTest{
		{
			TestCaseName: "testcase 1",
			input:        [9]byte{0x74, 0x65, 0x73, 0x74, 0x63, 0x61, 0x72, 0x74, 0x00},
			output: CarteActive{
				Nom: "testcart",
				Ord: "activation_carte",
			},
			err: nil,
		},
		{
			TestCaseName: "testcase 2",
			input:        [9]byte{0x74, 0x65, 0x73, 0x74, 0x63, 0x61, 0x72, 0x74, 0x20},
			output: CarteActive{
				Nom: "testcart",
				Ord: "annulation_carte",
			},
			err: nil,
		},
		{
			TestCaseName: "testcase 3",
			input:        [9]byte{0x74, 0x65, 0x73, 0x74, 0x63, 0x61, 0x72, 0x74, 0x40},
			output: CarteActive{
				Nom: "testcart",
				Ord: "unknowm",
			},
			err: ErrCartOrdUnknown,
		},
	}

	for _, row := range dataSet {
		// Act
		res, err := carte(row.input)

		// Assert
		if err != row.err {
			t.Errorf("FAIL: %s - error = %v; Expected: %v", row.TestCaseName, err, row.err)
		} else {
			t.Logf("SUCCESS: error: %v; Expected: %v", err, row.err)
		}

		if reflect.DeepEqual(res, row.output) == false {
			t.Errorf("FAIL: %s - %v; Expected: %v", row.TestCaseName, res, row.output)
		} else {
			t.Logf("SUCCESS: %v; Expected: %v", res, row.output)
		}
	}
}

func TestCat255STRModel_BiaisExtract(t *testing.T) {
	// setup
	type dataTest struct {
		TestCaseName string
		input        goasterix.Repetitive
		output       []BiaisRadar
	}
	dataSet := []dataTest{
		{
			TestCaseName: "testcase 1",
			input: goasterix.Repetitive{
				Rep: 0x01,
				Data: []byte{0x08, 0x81,
					0x18, 0xf0,
					0x00, 0xff,
					0x03, 0xe8,
					0x04, 0x00},
			},
			output: []BiaisRadar{
				{
					SacSic: SourceIdentifier{
						Sac: 8,
						Sic: 129,
					},
					GainDistance:  1,
					BiaisDistance: 255,
					BiaisAzimut:   5.5,
					BiaisDatation: 1,
				},
			},
		},
	}

	for _, row := range dataSet {
		// Act
		res := biaisExtract(row.input)

		// Assert
		if reflect.DeepEqual(res, row.output) == false {
			t.Errorf("FAIL: %s - %v; Expected: %v", row.TestCaseName, res, row.output)
		} else {
			t.Logf("SUCCESS: %v; Expected: %v", res, row.output)
		}
	}
}
