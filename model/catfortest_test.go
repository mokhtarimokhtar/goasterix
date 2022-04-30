package model

import (
	"github.com/mokhtarimokhtar/goasterix"
	"github.com/mokhtarimokhtar/goasterix/util"
	"reflect"
	"testing"
)

func TestCat026ModelWrite(t *testing.T) {
	// Arrange
	type testCase struct {
		Name   string
		input  string
		output *CatForTestModel
	}
	dataSet := []testCase{
		{
			Name:  "testcase 1",
			input: "80 08a2",
			output: &CatForTestModel{
				SacSic: &SourceIdentifier{Sac: 8, Sic: 162},
			},
		},
		/*
			{
				Name:  "testcase 2: I004/171",
				input: "fdf16008a2070108826b2100000608001fc1c05354434130333100194d40c1c33c6000002bd700bc000000001a491a4900000000001ec1c05354434130333000184d40c1c33c2000003039",
				output: &Cat004Model{
					SacSic:                  &SourceIdentifier{Sac: 8, Sic: 162},
					MessageType:             &MsgType{Code: "STCA", Desc: "short_term_conflict_alert"},
					SDPSIdentifier:          []SourceIdentifier{{Sac: 8, Sic: 130}},
					TimeOfMessage:           54850,
					AlertIdentifier:         6,
					AlertStatus:             4,
					TrackNumberOne:          31,
					ConflictCharacteristics: &ConflictCharacteristics{},
					ConflictTimingSeparation: &ConflictTimingSeparation{
						CurrentHorizontalSeparation: 3364.5,
						MinimumHorizontalSeparation: 3364.5,
					},
					AircraftOne: &AircraftIdentification{
						AircraftIdentifier: "STCA031",
						Mode3ACodeAircraft: "31",
						ModeSIdentifier:    "STCA031 ",
						FlightPlanNumber:   11223,
					},
					AircraftTwo: &AircraftIdentification{
						AircraftIdentifier: "STCA030",
						Mode3ACodeAircraft: "30",
						ModeSIdentifier:    "STCA030 ",
						FlightPlanNumber:   12345,
					},
					TrackNumberTwo: 30,
				},
			},
			{
				Name:  "testcase 3: I004/100",
				input: "fdf18008a2050108826b2a000005080015c1c04150573030323100110505f0c32c4000003039400080000000800505e050d060",
				output: &Cat004Model{
					SacSic:          &SourceIdentifier{Sac: 8, Sic: 162},
					MessageType:     &MsgType{Code: "APW", Desc: "area_proximity_warning"},
					SDPSIdentifier:  []SourceIdentifier{{Sac: 8, Sic: 130}},
					TimeOfMessage:   54868,
					AlertIdentifier: 5,
					AlertStatus:     4,
					TrackNumberOne:  21,
					ConflictCharacteristics: &ConflictCharacteristics{
						ConflictClassification: &ConflictClassification{
							TableId:            0,
							ConflictProperties: 0,
							CS:                 "low",
						},
					},
					ConflictTimingSeparation: &ConflictTimingSeparation{},
					AircraftOne: &AircraftIdentification{
						AircraftIdentifier: "APW0021",
						Mode3ACodeAircraft: "21",
						ModeSIdentifier:    "APW0021",
						FlightPlanNumber:   12345,
					},
					AreaDefinition: &AreaDefinition{AreaName: "APW TMA "},
				},
			},
			{
				Name:  "testcase 4: I004/070",
				input: "fdf18008a2050108826b29000005080015c1c04150573030323100110505f0c32c4000003039400080000000800505e050d060",
				output: &Cat004Model{
					SacSic:          &SourceIdentifier{Sac: 8, Sic: 162},
					MessageType:     &MsgType{Code: "APW", Desc: "area_proximity_warning"},
					SDPSIdentifier:  []SourceIdentifier{{Sac: 8, Sic: 130}},
					TimeOfMessage:   54866,
					AlertIdentifier: 5,
					AlertStatus:     4,
					TrackNumberOne:  21,
					AreaDefinition:  &AreaDefinition{AreaName: "APW TMA "},
					ConflictCharacteristics: &ConflictCharacteristics{
						ConflictClassification: &ConflictClassification{
							TableId:            0,
							ConflictProperties: 0,
							CS:                 "low",
						},
					},
					ConflictTimingSeparation: &ConflictTimingSeparation{},
					AircraftOne: &AircraftIdentification{
						AircraftIdentifier: "APW0021",
						Mode3ACodeAircraft: "21",
						ModeSIdentifier:    "APW0021",
						FlightPlanNumber:   12345,
					},
				},
			},
		*/
	}

	for _, tc := range dataSet {
		// Arrange
		uap := goasterix.CatForTest
		data, _ := util.HexStringToByte(tc.input)
		rec := goasterix.NewRecord()
		_, _ = rec.Decode(data, uap)
		t.Log(rec.String())
		model := new(CatForTestModel)

		// Act
		model.write(*rec)

		//recJson, _ := json.Marshal(model)
		//t.Log(string(recJson))
		//t.Log(rec.String())

		// Assert
		if reflect.DeepEqual(model, tc.output) == false {
			t.Errorf(util.MsgFailInValue, tc.Name, model, tc.output)
		} else {
			t.Logf(util.MsgSuccessInValue, tc.Name, model, tc.output)
		}
	}
}
