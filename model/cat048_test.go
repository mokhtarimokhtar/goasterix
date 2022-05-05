package model

import (
	"bytes"
	"encoding/json"
	"github.com/mokhtarimokhtar/goasterix"
	"github.com/mokhtarimokhtar/goasterix/util"
	"reflect"
	"testing"
)

func TestCat048ModelWrite(t *testing.T) {
	// Arrange
	type testCase struct {
		Name   string
		input  string
		output *Cat048Model
	}
	dataSet := []testCase{
		{
			Name:  "testcase 1",
			input: "80 0836",
			output: &Cat048Model{
				SacSic: &SourceIdentifier{Sac: 8, Sic: 54},
			},
		},
		{
			Name:  "testcase 2",
			input: "40 429b52",
			output: &Cat048Model{
				TimeOfDay: 34102.640625,
			},
		},
		{
			Name:  "testcase 14",
			input: "0102 4100", // 0000-0001 0000-0010
			output: &Cat048Model{
				TrackStatus: &Status{
					CNF: "confirmed_track",
					RAD: "ssr_modes_track",
					DOU: "normal_confidence",
					MAH: "no_horizontal_man_sensed",
					CDM: "maintaining",

					TRE: "track_still_alive",
					GHO: "true_target_track",
					SUP: "no",
					TCC: "radar_plane",
				},
			},
		},

		//{
		//	Name:  "testcase 2: I004/171",
		//	input: "fdf16008a2070108826b2100000608001fc1c05354434130333100194d40c1c33c6000002bd700bc000000001a491a4900000000001ec1c05354434130333000184d40c1c33c2000003039",
		//	output: &Cat004Model{
		//		SacSic:                  &SourceIdentifier{Sac: 8, Sic: 162},
		//		MessageType:             &MsgType{Code: "STCA", Desc: "short_term_conflict_alert"},
		//		SDPSIdentifier:          []SourceIdentifier{{Sac: 8, Sic: 130}},
		//		TimeOfMessage:           54850,
		//		AlertIdentifier:         6,
		//		AlertStatus:             4,
		//		TrackNumberOne:          31,
		//		ConflictCharacteristics: &ConflictCharacteristics{},
		//		ConflictTimingSeparation: &ConflictTimingSeparation{
		//			CurrentHorizontalSeparation: 3364.5,
		//			MinimumHorizontalSeparation: 3364.5,
		//		},
		//		AircraftOne: &AircraftIdentification{
		//			AircraftIdentifier: "STCA031",
		//			Mode3ACodeAircraft: "31",
		//			ModeSIdentifier:    "STCA031 ",
		//			FlightPlanNumber:   11223,
		//		},
		//		AircraftTwo: &AircraftIdentification{
		//			AircraftIdentifier: "STCA030",
		//			Mode3ACodeAircraft: "30",
		//			ModeSIdentifier:    "STCA030 ",
		//			FlightPlanNumber:   12345,
		//		},
		//		TrackNumberTwo: 30,
		//	},
		//},
		//{
		//	Name:  "testcase 3: I004/100",
		//	input: "fdf18008a2050108826b2a000005080015c1c04150573030323100110505f0c32c4000003039400080000000800505e050d060",
		//	output: &Cat004Model{
		//		SacSic:          &SourceIdentifier{Sac: 8, Sic: 162},
		//		MessageType:     &MsgType{Code: "APW", Desc: "area_proximity_warning"},
		//		SDPSIdentifier:  []SourceIdentifier{{Sac: 8, Sic: 130}},
		//		TimeOfMessage:   54868,
		//		AlertIdentifier: 5,
		//		AlertStatus:     4,
		//		TrackNumberOne:  21,
		//		ConflictCharacteristics: &ConflictCharacteristics{
		//			ConflictClassification: &ConflictClassification{
		//				TableId:            0,
		//				ConflictProperties: 0,
		//				CS:                 "low",
		//			},
		//		},
		//		ConflictTimingSeparation: &ConflictTimingSeparation{},
		//		AircraftOne: &AircraftIdentification{
		//			AircraftIdentifier: "APW0021",
		//			Mode3ACodeAircraft: "21",
		//			ModeSIdentifier:    "APW0021",
		//			FlightPlanNumber:   12345,
		//		},
		//		AreaDefinition: &AreaDefinition{AreaName: "APW TMA "},
		//	},
		//},
		//{
		//	Name:  "testcase 4: I004/070",
		//	input: "fdf18008a2050108826b29000005080015c1c04150573030323100110505f0c32c4000003039400080000000800505e050d060",
		//	output: &Cat004Model{
		//		SacSic:          &SourceIdentifier{Sac: 8, Sic: 162},
		//		MessageType:     &MsgType{Code: "APW", Desc: "area_proximity_warning"},
		//		SDPSIdentifier:  []SourceIdentifier{{Sac: 8, Sic: 130}},
		//		TimeOfMessage:   54866,
		//		AlertIdentifier: 5,
		//		AlertStatus:     4,
		//		TrackNumberOne:  21,
		//		AreaDefinition:  &AreaDefinition{AreaName: "APW TMA "},
		//		ConflictCharacteristics: &ConflictCharacteristics{
		//			ConflictClassification: &ConflictClassification{
		//				TableId:            0,
		//				ConflictProperties: 0,
		//				CS:                 "low",
		//			},
		//		},
		//		ConflictTimingSeparation: &ConflictTimingSeparation{},
		//		AircraftOne: &AircraftIdentification{
		//			AircraftIdentifier: "APW0021",
		//			Mode3ACodeAircraft: "21",
		//			ModeSIdentifier:    "APW0021",
		//			FlightPlanNumber:   12345,
		//		},
		//	},
		//},
	}

	for _, tc := range dataSet {
		// Arrange
		uap := goasterix.Cat048V127
		data, _ := util.HexStringToByte(tc.input)
		rb := bytes.NewReader(data)
		rec := goasterix.NewRecord()
		_, _ = rec.Decode(rb, uap)
		model := new(Cat048Model)

		// Act
		model.write(*rec)

		// Assert
		if reflect.DeepEqual(model, tc.output) == false {
			t.Errorf(util.MsgFailInValue, tc.Name, model, tc.output)
		} else {
			t.Logf(util.MsgSuccessInValue, tc.Name, model, tc.output)
		}
	}
}

func TestCat048Model_ToJsonRecord(t *testing.T) {
	// Arrange
	// bds 02 e79a5d27a00c00 60 a3280030a40000 40
	input := "ffff02 0836 429b52 a0 94c70181 0913 02d0 6002b7 490d01 38a178cf4220 02e79a5d27a00c0060a3280030a4000040 063a 00800080 0743ce5b 40 20f5"
	output := []byte(`{"sourceIdentifier":{"sac":8,"sic":54},"aircraftAddress":"490D01","aircraftIdentification":"NJE834H ","timeOfDay":34102.640625,"rhoTheta":{"rho":148.77734375,"theta":2.1174999999999997},"cartesianXY":{"x":1,"y":1},"flightLevel":{"v":"code_validated","g":"default","level":180},"radarPlotCharacteristics":{"srr":2,"sam":-73},"mode3ACode":{"squawk":"4423","v":"code_validated","g":"default","l":"code_derived_from_transponder"},"trackNumber":1594,"trackVelocity":{"groundSpeed":0.113464065,"heading":290.5485},"trackStatus":{"cnf":"confirmed_track","rad":"ssr_modes_track","dou":"normal_confidence","mah":"no_horizontal_man_sensed","cdm":"maintaining"},"bdsRegisterData":[{"transponderRegisterNumber":"60","code60":{"magneticHeading":-68,"indicatedAirspeed":302,"mach":0.632,"barometricAltitudeRate":32}},{"transponderRegisterNumber":"40","code40":{"mcpSelectAltitude":18000,"barometricPressureSetting":1013}}],"comAcasCapabilityFlightStatus":{"com":"comm_a_and_comm_b_capability","stat":"no_alert_no_spi_aircraft_airborne","si":"si_code_capable","mssc":"yes","arc":"25_ft_resolution","aic":"yes","b1a":"1","b1b":"5"}}`)
	/*
		outputStr := `{
					"sourceIdentifier":{"sac":8,"sic":54},
					"aircraftAddress":"490D01",
					"aircraftIdentification":"NJE834H ",
					"timeOfDay":34102.640625,
					"rhoTheta":{"rho":148.77734375,"theta":2.1174999999999997},
					"cartesianXY":{"x":1,"y":1},
					"flightLevel":{"v":"code_validated","g":"default","level":180},
					"radarPlotCharacteristics":{"srr":2,"sam":-73},
					"mode3ACode":{"squawk":"4423","v":"code_validated","g":"default","l":"code_derived_from_transponder"},
					"trackNumber":1594,
					"trackVelocity":{"groundSpeed":0.113464065,"heading":290.5485},
					"trackStatus":{"cnf":"confirmed_track","rad":"ssr_modes_track","dou":"normal_confidence","mah":"no_horizontal_man_sensed","cdm":"maintaining"},
					"bdsRegisterData":[
						{"transponderRegisterNumber":"60","code60":{"magneticHeading":-68,"indicatedAirspeed":302,"mach":0.632,"barometricAltitudeRate":32}},
						{"transponderRegisterNumber":"40","code40":{"mcpSelectAltitude":18000,"barometricPressureSetting":1013}}
						],
					"comAcasCapabilityFlightStatus":{"com":"comm_a_and_comm_b_capability","stat":"no_alert_no_spi_aircraft_airborne","si":"si_code_capable","mssc":"yes","arc":"25_ft_resolution","aic":"yes","b1a":"1","b1b":"5"}
					}
					`

		output := []byte(util.CleanStringMultiline(outputStr))
	*/

	uap := goasterix.Cat048V127
	data, _ := util.HexStringToByte(input)
	rb := bytes.NewReader(data)
	rec := goasterix.NewRecord()
	_, err := rec.Decode(rb, uap)
	model := new(Cat048Model)
	model.write(*rec)

	// Act
	recJson, _ := json.Marshal(model)

	// Assert
	if err != nil {
		t.Errorf(util.MsgFailInValue, "", err, nil)
	} else {
		t.Logf(util.MsgSuccessInValue, "", err, nil)
	}

	if reflect.DeepEqual(recJson, output) == false {
		t.Errorf(util.MsgFailInString, "", recJson, output)
	} else {
		t.Logf(util.MsgSuccessInString, "", recJson, output)
	}
}
