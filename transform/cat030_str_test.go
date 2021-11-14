package transform

import (
	"encoding/json"
	"github.com/mokhtarimokhtar/goasterix"
	"github.com/mokhtarimokhtar/goasterix/uap"
	"reflect"
	"testing"
)

func TestCat030STRModel_ToJsonRecord(t *testing.T) {
	// Arrange
	input := "bfff0160 0885 5801b8 6092fc 010e 0200 0925f483 0c 04e6 04ea fb5ff9c4 f8 fd9a 0d0174 48455b 2cc371cf1de0"
	output := []byte(`{"sourceIdentifier":{"sac":8,"sic":133},"num":{"version":2,"nap":3,"st":"operational","ns":"principal","numero":220},"hptu":49445.96875,"pist":{"liv":"trafic_reel","cnf":"piste_confirmee","man":"defaut","tva":"defaut","type":"piste_association_multiple_primaire_secondaire","mort":"defaut","cre":"defaut","slr":"coordonnees_projetees_niveau_calcule","cor":"piste_non_correlee_plan_vol"},"alis":{"v":"code_valide","g":"defaut","c":"code_pas_changement","code":1000},"pos":{"x":36.578125,"y":-45.953125},"qual":6,"flpc":{"vc":"code_validated","gc":"default","niveauVol":313.5},"flpm":{"vc":"code_validated","gc":"default","niveauVol":314.5},"vit":{"x":-0.072326475,"y":-0.09741186},"mov":{"trans":"tendance_indeterminee","longi":"tendance_indeterminee","verti":"vol_descente"},"taux":-3597.65625,"spe":{"sy":1,"m":1,"s":0,"o1":0,"o2":0,"o3":0,"o4":0,"o5":0,"o6":0,"o7":1,"o8":0,"o9":1,"o10":1,"o11":1,"o12":0,"o13":0,"o14":0,"o15":0,"o16":0,"o17":0,"o18":0,"o19":0,"r":0,"c":0},"adrs":"48455B","ids":"KLM1317 "}`)

	uap030 := uap.Cat030StrV51
	data, _ := goasterix.HexStringToByte(input)
	rec := new(goasterix.Record)
	_, err := rec.Decode(data, uap030)

	cat030Model := new(Cat030STRModel)
	cat030Model.write(rec.Items)

	// Act
	recJson, _ := json.Marshal(cat030Model)

	// Assert
	if err != nil {
		t.Errorf("FAIL: error = %v; Expected: %v", err, nil)
	} else {
		t.Logf("SUCCESS: error: %v; Expected: %v", err, nil)
	}

	if reflect.DeepEqual(recJson, output) == false {
		t.Errorf("FAIL: %s; \nExpected: %s", recJson, output)
	} else {
		t.Logf("SUCCESS: %s; Expected: %s", recJson, output)
	}
}

func TestCat030STRModel_VitCal(t *testing.T) {
	// Arrange
	input := [4]byte{0x27, 0x10, 0x27, 0x10}
	output := Vit{
		X: 0.6103500000000001,
		Y: 0.6103500000000001,
	}

	// Act
	res := vitCal(input)

	// Assert
	if reflect.DeepEqual(res, output) == false {
		t.Errorf("FAIL: %v; Expected: %v", res, output)
	} else {
		t.Logf("SUCCESS: %v; Expected: %v", res, output)
	}

}

func TestCat030STRModel_Flp(t *testing.T) {
	// setup
	type dataTest struct {
		TestCaseName string
		input        [2]byte
		output       Flstr
	}
	dataset := []dataTest{
		{
			TestCaseName: "testcase 1",
			input:        [2]byte{0x00, 0xFF},
			output: Flstr{
				Vc:        "code_validated",
				Gc:        "default",
				NiveauVol: 63.75,
			},
		},
		{
			TestCaseName: "testcase 2",
			input:        [2]byte{0xc0, 0xFF}, //1100 0000 1111 1111
			output: Flstr{
				Vc:        "code_not_validated",
				Gc:        "garbled_code",
				NiveauVol: 63.75,
			},
		},
	}

	for _, row := range dataset {
		// Act
		res := flp(row.input)

		// Assert
		if reflect.DeepEqual(res, row.output) == false {
			t.Errorf("FAIL: %v; Expected: %v", res, row.output)
		} else {
			t.Logf("SUCCESS: %v; Expected: %v", res, row.output)
		}
	}
}

func TestCat030STRModel_Pos(t *testing.T) {
	// Arrange
	input := [4]byte{0x27, 0x10, 0x27, 0x10}
	output := CartesianXYPosition{
		X: 156.25,
		Y: 156.25,
	}

	// Act
	res := pos(input)

	// Assert
	if reflect.DeepEqual(res, output) == false {
		t.Errorf("FAIL: %v; Expected: %v", res, output)
	} else {
		t.Logf("SUCCESS: %v; Expected: %v", res, output)
	}

}

func TestCat030STRModel_Num(t *testing.T) {
	// setup
	type dataTest struct {
		TestCaseName string
		input        [3]byte
		output       NumPiste
	}
	dataset := []dataTest{
		{
			TestCaseName: "testcase 1",
			input:        [3]byte{0x30, 0x00, 0xc8},
			output: NumPiste{
				Version: 1,
				Nap:     2,
				ST:      "operational",
				NS:      "principal",
				Numero:  100,
			},
		},
		{
			TestCaseName: "testcase 2",
			input:        [3]byte{0x35, 0x00, 0xc8}, //0011 0101
			output: NumPiste{
				Version: 1,
				Nap:     2,
				ST:      "evaluation",
				NS:      "secours",
				Numero:  100,
			},
		},
		{
			TestCaseName: "testcase 3",
			input:        [3]byte{0x36, 0x00, 0xc8},
			output: NumPiste{
				Version: 1,
				Nap:     2,
				ST:      "evaluation",
				NS:      "test",
				Numero:  100,
			},
		},
	}

	for _, row := range dataset {
		// Act
		res := num(row.input)

		// Assert
		if reflect.DeepEqual(res, row.output) == false {
			t.Errorf("FAIL: %v; Expected: %v", res, row.output)
		} else {
			t.Logf("SUCCESS: %v; Expected: %v", res, row.output)
		}
	}
}

func TestCat030STRModel_Mov(t *testing.T) {
	// setup
	type dataTest struct {
		TestCaseName string
		input        [1]byte
		output       Mov
	}
	dataset := []dataTest{
		{
			TestCaseName: "testcase 1",
			input:        [1]byte{0x00},
			output: Mov{
				Trans: "ligne_droite",
				Longi: "vitesse_sol_constante",
				Verti: "vol_palier",
			},
		},
		{
			TestCaseName: "testcase 2",
			input:        [1]byte{0x54}, //0101 0100
			output: Mov{
				Trans: "virage_droite",
				Longi: "vitesse_sol_augmentation",
				Verti: "vol_montee",
			},
		},
		{
			TestCaseName: "testcase 3",
			input:        [1]byte{0xa8}, //1010 1000
			output: Mov{
				Trans: "virage_gauche",
				Longi: "vitesse_sol_diminution",
				Verti: "vol_descente",
			},
		},
		{
			TestCaseName: "testcase 4",
			input:        [1]byte{0xfc}, //1111 1100
			output: Mov{
				Trans: "tendance_indeterminee",
				Longi: "tendance_indeterminee",
				Verti: "tendance_indeterminee",
			},
		},
	}

	for _, row := range dataset {
		// Act
		res := mov(row.input)

		// Assert
		if reflect.DeepEqual(res, row.output) == false {
			t.Errorf("FAIL: %v; Expected: %v", res, row.output)
		} else {
			t.Logf("SUCCESS: %v; Expected: %v", res, row.output)
		}
	}
}

func TestCat030STRModel_Altic(t *testing.T) {
	// Arrange
	input := [2]byte{0x80, 0xff}
	output := Altic{
		QNC: 1,
		Alt: 255,
	}

	// Act
	res := altic(input)

	// Assert
	if reflect.DeepEqual(res, output) == false {
		t.Errorf("FAIL: %v; Expected: %v", res, output)
	} else {
		t.Logf("SUCCESS: %v; Expected: %v", res, output)
	}

}
