package transform

import (
	"encoding/json"
	"github.com/mokhtarimokhtar/goasterix"
	"github.com/mokhtarimokhtar/goasterix/uap"
	"github.com/mokhtarimokhtar/goasterix/util"
	"reflect"
	"testing"
)

func TestCat030STRModel_ToJsonRecord(t *testing.T) {
	// setup
	type dataTest struct {
		TestCaseName string
		input        string
		output       []byte
	}
	dataSet := []dataTest{
		{
			TestCaseName: "Testcase 1",
			input:        "bfff0160 0885 5801b8 6092fc 010e 0200 0925f483 0c 04e6 04ea fb5ff9c4 f8 fd9a 0d0174 48455b 2cc371cf1de0",
			output:       []byte(`{"sourceIdentifier":{"sac":8,"sic":133},"num":{"version":2,"nap":3,"st":"operational","ns":"principal","numero":220},"hptu":49445.96875,"pist":{"liv":"trafic_reel","cnf":"piste_confirmee","man":"defaut","tva":"defaut","type":"piste_association_multiple_primaire_secondaire","mort":"defaut","cre":"defaut","slr":"coordonnees_projetees_niveau_calcule","cor":"piste_non_correlee_plan_vol"},"alis":{"v":"code_valide","g":"defaut","c":"code_pas_changement","code":1000},"pos":{"x":36.578125,"y":-45.953125},"qual":6,"flpc":{"vc":"code_validated","gc":"default","niveauVol":313.5},"flpm":{"vc":"code_validated","gc":"default","niveauVol":314.5},"vit":{"x":-0.072326475,"y":-0.09741186},"mov":{"trans":"tendance_indeterminee","longi":"tendance_indeterminee","verti":"vol_descente"},"taux":-3597.65625,"spe":{"sy":1,"m":1,"s":0,"o1":0,"o2":0,"o3":0,"o4":0,"o5":0,"o6":0,"o7":1,"o8":0,"o9":1,"o10":1,"o11":1,"o12":0,"o13":0,"o14":0,"o15":0,"o16":0,"o17":0,"o18":0,"o19":0,"r":0,"c":0},"adrs":"48455B","ids":"KLM1317 "}`),
		},
		{
			TestCaseName: "Testcase 2",
			input:        "37fb7f604806f466ee0a094be45bc08c0e05f005f00540060cf0370b0252595234303537019e423733384d4c454d47454444484ca2aa4994b4c35de0",
			output:       []byte(`{"num":{"version":2,"nap":1,"st":"operational","ns":"principal","numero":890},"hptu":52700.078125,"alis":{"v":"code_valide","g":"defaut","c":"code_pas_changement","code":4513},"pos":{"x":-110.578125,"y":-253.8125},"qual":7,"flpc":{"vc":"code_validated","gc":"default","niveauVol":380},"flpm":{"vc":"code_validated","gc":"default","niveauVol":380},"vit":{"x":0.08203104,"y":0.09448218},"mov":{"trans":"tendance_indeterminee","longi":"tendance_indeterminee","verti":"vol_palier"},"spe":{"sy":6,"m":1,"s":1,"o1":0,"o2":0,"o3":0,"o4":0,"o5":0,"o6":1,"o7":0,"o8":0,"o9":0,"o10":0,"o11":0,"o12":0,"o13":1,"o14":0,"o15":1,"o16":0,"o17":0,"o18":0,"o19":0,"r":0,"c":0},"ivol":"RYR4057","pln":414,"av":"B738","turb":"M","terd":"LEMG","tera":"EDDH","adrs":"4CA2AA","ids":"RYR4057 "}`),
		},
		{
			TestCaseName: "Testcase 3",
			input:        "3ffb81604806e466ee0a090e0ecee134bee00e0154015401a20156f037090c08183465101826721724e0",
			output:       []byte(`{"num":{"version":2,"nap":1,"st":"operational","ns":"principal","numero":882},"hptu":52700.078125,"pist":{"liv":"trafic_reel","cnf":"piste_confirmee","man":"defaut","tva":"defaut","type":"piste_monoradar_secondaire_pure","mort":"defaut","cre":"defaut","slr":"coordonnees_projetees_niveau_calcule","cor":"piste_non_correlee_plan_vol"},"alis":{"v":"code_valide","g":"defaut","c":"code_pas_changement","code":7316},"pos":{"x":-123.1875,"y":-260.5},"qual":7,"flpc":{"vc":"code_validated","gc":"default","niveauVol":85},"flpm":{"vc":"code_validated","gc":"default","niveauVol":85},"vit":{"x":0.02551263,"y":0.02087397},"mov":{"trans":"tendance_indeterminee","longi":"tendance_indeterminee","verti":"vol_palier"},"spe":{"sy":6,"m":1,"s":1,"o1":0,"o2":0,"o3":0,"o4":0,"o5":0,"o6":0,"o7":1,"o8":1,"o9":0,"o10":0,"o11":0,"o12":0,"o13":0,"o14":0,"o15":1,"o16":0,"o17":0,"o18":0,"o19":0,"r":0,"c":0},"radSacSic":{"sac":8,"sic":24},"adrs":"346510","ids":"FBY2E2S "}`),
		},
		{
			TestCaseName: "Testcase 4",
			input:        "3ffb7f605800a23abfd8014002004e84d1460e3ffe3ffe013700f3f017010c444c48313945202057453139304d4c464d4e454444463c546510c231e45820",
			output:       []byte(`{"num":{"version":2,"nap":3,"st":"operational","ns":"principal","numero":81},"hptu":30079.6875,"pist":{"liv":"trafic_reel","cnf":"piste_confirmee","man":"defaut","tva":"defaut","type":"piste_association_multiple_primaire_secondaire","mort":"defaut","cre":"creation_de_piste","slr":"coordonnees_projetees_niveau_calcule","cor":"piste_correlation_plan_vol_confirmee"},"alis":{"v":"code_valide","g":"defaut","c":"code_pas_changement","code":1000},"pos":{"x":314.0625,"y":-186.90625},"qual":7,"flpc":{"vc":"code_validated","gc":"default","niveauVol":-0.5},"flpm":{"vc":"code_validated","gc":"default","niveauVol":-0.5},"vit":{"x":0.018981885,"y":0.014831505},"mov":{"trans":"tendance_indeterminee","longi":"tendance_indeterminee","verti":"vol_palier"},"spe":{"sy":2,"m":1,"s":1,"o1":0,"o2":0,"o3":0,"o4":0,"o5":0,"o6":0,"o7":1,"o8":1,"o9":0,"o10":0,"o11":0,"o12":0,"o13":0,"o14":0,"o15":0,"o16":0,"o17":0,"o18":0,"o19":0,"r":0,"c":0},"ivol":"DLH19E ","pln":8279,"av":"E190","turb":"M","terd":"LFMN","tera":"EDDF","adrs":"3C5465","ids":"DLH19E  "}`),
		},
	}

	for _, row := range dataSet {
		// Arrange
		uap030 := uap.Cat030StrV51
		data, _ := util.HexStringToByte(row.input)
		rec := new(goasterix.Record)
		_, err := rec.Decode(data, uap030)

		cat030Model := new(Cat030STRModel)
		cat030Model.write(*rec)

		// Act
		recJson, _ := json.Marshal(cat030Model)

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
	dataSet := []dataTest{
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

	for _, row := range dataSet {
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

func TestCat030STRModel_Pist(t *testing.T) {
	// setup
	type dataTest struct {
		TestCaseName string
		input        goasterix.Extended
		output       Pist
	}
	dataset := []dataTest{
		{
			TestCaseName: "testcase 1",
			input: goasterix.Extended{
				Primary:   []byte{0x01},
				Secondary: []byte{0x01, 0x01},
			},
			output: Pist{
				Liv:    "trafic_reel",
				Cnf:    "piste_confirmee",
				Man:    "defaut",
				Tva:    "defaut",
				Type:   "piste_association_multiple_primaire_secondaire",
				Mort:   "defaut",
				Cre:    "defaut",
				Slr:    "coordonnees_projetees_niveau_calcule",
				Cor:    "piste_correlation_plan_vol_confirmee",
				Ds1ds2: "defaut",
				For:    "defaut",
				Ama:    "piste_amalgamee",
				Spi:    "defaut",
				Me:     "defaut",
			},
		},
		{
			TestCaseName: "testcase 2",
			input: goasterix.Extended{
				Primary:   []byte{0xf3},
				Secondary: []byte{0xd3, 0x7c},
			},
			output: Pist{
				Liv:    "simule_ou_plot_test",
				Cnf:    "piste_initialisation",
				Man:    "piste_virage",
				Tva:    "piste_pas_niveau_vol_valide",
				Type:   "piste_association_primaire_pure",
				Mort:   "mort_de_piste",
				Cre:    "creation_de_piste",
				Slr:    "coordonnees_projetees_niveau_mesure",
				Cor:    "piste_correlation_plan_vol_associee",
				Ds1ds2: "detournement_code_75000",
				For:    "vol_en_formation",
				Ama:    "piste_non_amalgamee",
				Spi:    "special_pulse_ident",
				Me:     "detresse_militaire",
			},
		},
	}

	for _, row := range dataset {
		// Act
		res := pist(row.input)

		// Assert
		if reflect.DeepEqual(res, row.output) == false {
			t.Errorf("FAIL: %s - %v; Expected: %v", row.TestCaseName, res, row.output)
		} else {
			t.Logf("SUCCESS: %v; Expected: %v", res, row.output)
		}
	}
}

func TestCat030STRModel_Alis(t *testing.T) {
	// setup
	type dataTest struct {
		TestCaseName string
		input        [2]byte
		output       ModeA
	}
	dataset := []dataTest{
		{
			TestCaseName: "testcase 1",
			input:        [2]byte{0xef, 0xff},
			output: ModeA{
				V:    "code_invalide",
				G:    "code_garbling",
				C:    "code_changement",
				Code: 7777,
			},
		},
		{
			TestCaseName: "testcase 2",
			input:        [2]byte{0x0f, 0xff},
			output: ModeA{
				V:    "code_valide",
				G:    "defaut",
				C:    "code_pas_changement",
				Code: 7777,
			},
		},
	}

	for _, row := range dataset {
		// Act
		res := alis(row.input)

		// Assert
		if reflect.DeepEqual(res, row.output) == false {
			t.Errorf("FAIL: %s %v; Expected: %v", row.TestCaseName, res, row.output)
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

func TestCat030STRModel_Spe(t *testing.T) {
	// setup
	type dataTest struct {
		TestCaseName string
		input        goasterix.Extended
		output       Spe
	}
	dataset := []dataTest{
		{
			TestCaseName: "testcase 1",
			input: goasterix.Extended{
				Primary:   []byte{0x00},
				Secondary: []byte{0x00, 0x00, 0x00},
			},
			output: Spe{
				SY:  0,
				M:   0,
				S:   0,
				O1:  0,
				O2:  0,
				O3:  0,
				O4:  0,
				O5:  0,
				O6:  0,
				O7:  0,
				O8:  0,
				O9:  0,
				O10: 0,
				O11: 0,
				O12: 0,
				O13: 0,
				O14: 0,
				O15: 0,
				O16: 0,
				O17: 0,
				O18: 0,
				O19: 0,
				R:   0,
				C:   0,
			},
		},
		{
			TestCaseName: "testcase 2",
			input: goasterix.Extended{
				Primary:   []byte{0x7f},
				Secondary: []byte{0xff, 0xff, 0xff},
			},
			output: Spe{
				SY:  15,
				M:   1,
				S:   1,
				O1:  1,
				O2:  1,
				O3:  1,
				O4:  1,
				O5:  1,
				O6:  1,
				O7:  1,
				O8:  1,
				O9:  1,
				O10: 1,
				O11: 1,
				O12: 1,
				O13: 1,
				O14: 1,
				O15: 1,
				O16: 1,
				O17: 1,
				O18: 1,
				O19: 1,
				R:   1,
				C:   1,
			},
		},
	}

	for _, row := range dataset {
		// Act
		res := spe(row.input)

		// Assert
		if reflect.DeepEqual(res, row.output) == false {
			t.Errorf("FAIL: %s - %v; Expected: %v", row.TestCaseName, res, row.output)
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
